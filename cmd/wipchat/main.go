package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	ff "github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/shurcooL/graphql"
	"moul.io/godev"
	"moul.io/motd"
	"moul.io/wipchat"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

func run(_ []string) error {
	var (
		attachPaths          stringSlice
		apiKey               string
		debug                bool
		todosCompletedFilter bool
		todosLimitFilter     int
		todosOffsetFilter    int
		todosFilterFilter    string
		todosOrderFilter     string
	)
	rootFlags := flag.NewFlagSet("root", flag.ExitOnError)
	rootFlags.StringVar(&apiKey, "key", "", "Your private API key from https://wip.chat/api")
	rootFlags.BoolVar(&debug, "debug", false, "More verbose output")
	todoFlags := flag.NewFlagSet("todo", flag.ExitOnError)
	todoFlags.Var(&attachPaths, "attach", "attachment paths or URLs")
	meFlags := flag.NewFlagSet("me", flag.ExitOnError)
	meFlags.BoolVar(&todosCompletedFilter, "todos-completed", true, "todos completed filter")
	meFlags.IntVar(&todosLimitFilter, "todos-limit", 20, "todos limit filter")
	meFlags.IntVar(&todosOffsetFilter, "todos-offset", 0, "todos offset filter")
	meFlags.StringVar(&todosFilterFilter, "todos-filter", "", "todos filter filter")
	meFlags.StringVar(&todosOrderFilter, "todos-order", "", "todos order filter")

	root := &ffcli.Command{
		Name:       "wipchat",
		ShortUsage: "wipchat --key=<key> <subcommand>",
		FlagSet:    rootFlags,
		Options:    []ff.Option{ff.WithEnvVarPrefix("WIPCHAT")},
		Subcommands: []*ffcli.Command{
			{
				Name:      "me",
				ShortHelp: "retrieve user info about current token",
				FlagSet:   meFlags,
				Exec: func(ctx context.Context, _ []string) error {
					client := wipchat.New(apiKey)
					viewer, err := client.QueryViewer(ctx, &wipchat.QueryViewerOptions{
						TodosCompleted: graphql.Boolean(todosCompletedFilter),
						TodosLimit:     graphql.Int(todosLimitFilter),
						TodosOffset:    graphql.Int(todosOffsetFilter),
						TodosFilter:    graphql.String(todosFilterFilter),
						TodosOrder:     graphql.String(todosOrderFilter),
					})
					if err != nil {
						return err
					}
					fmt.Println(godev.PrettyJSON(viewer))
					return nil
				},
			}, {
				Name:       "todo",
				ShortUsage: "wipchat todo <lorem ipsum>",
				ShortHelp:  "create a new todo task",
				FlagSet:    todoFlags,
				Exec: func(ctx context.Context, args []string) error {
					body := strings.TrimSpace(strings.Join(args, " "))
					if len(body) < 1 {
						return flag.ErrHelp
					}
					attachments, err := loadAttachPaths(attachPaths)
					if err != nil {
						return err
					}
					client := wipchat.New(apiKey)
					task, err := client.MutateCreateTodo(ctx, body, nil, attachments)
					if err != nil {
						return err
					}

					if debug {
						fmt.Fprintln(os.Stderr, godev.PrettyJSON(task))
					}
					fmt.Println(task.CanonicalURL())
					return nil
				},
			}, {
				Name:       "done",
				ShortUsage: "wipchat done <lorem ipsum>",
				ShortHelp:  "create a new completed task",
				FlagSet:    todoFlags,
				Exec: func(ctx context.Context, args []string) error {
					body := strings.TrimSpace(strings.Join(args, " "))
					if len(body) < 1 {
						return flag.ErrHelp
					}
					attachments, err := loadAttachPaths(attachPaths)
					if err != nil {
						return err
					}
					client := wipchat.New(apiKey)
					now := time.Now()
					task, err := client.MutateCreateTodo(ctx, body, &now, attachments)
					if err != nil {
						return err
					}

					if debug {
						fmt.Fprintln(os.Stderr, godev.PrettyJSON(task))
					}
					fmt.Println(task.CanonicalURL())
					return nil
				},
			}, {
				Name:      "todos",
				ShortHelp: "todos posted by makers",
				Exec: func(ctx context.Context, _ []string) error {
					client := wipchat.New(apiKey)
					todos, err := client.QueryTodos(ctx)
					if err != nil {
						return err
					}
					fmt.Println(godev.PrettyJSON(todos))
					return nil
				},
			}, {
				Name:      "products",
				ShortHelp: "products people are making",
				Exec: func(ctx context.Context, _ []string) error {
					client := wipchat.New(apiKey)
					products, err := client.QueryProducts(ctx)
					if err != nil {
						return err
					}
					fmt.Println(godev.PrettyJSON(products))
					return nil
				},
			},
		},
		Exec: func(_ context.Context, _ []string) error {
			fmt.Fprintln(os.Stderr, motd.Default())
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), os.Args[1:])
}

func loadAttachPaths(paths []string) ([]wipchat.FilePayload, error) {
	// FIXME: support URL
	ret := make([]wipchat.FilePayload, len(paths))
	for i, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return nil, fmt.Errorf("open %q: %w", p, err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("read %q: %w", p, err)
		}
		if len(b) < 1 {
			return nil, fmt.Errorf("empty file %q: %w", p, err)
		}
		contentType := http.DetectContentType(b)
		if !strings.HasPrefix(contentType, "image/") {
			return nil, fmt.Errorf("invalid content-type %q: %q", p, contentType)
		}
		ret[i] = wipchat.FilePayload{
			Filename: path.Base(p),
			Bytes:    b,
		}
	}
	return ret, nil
}
