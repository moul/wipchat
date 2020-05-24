package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	ff "github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

func run(_ []string) error {
	rootFlags := flag.NewFlagSet("root", flag.ExitOnError)
	apiKey := rootFlags.String("key", "", "Your private API key from https://wip.chat/api")

	root := &ffcli.Command{
		Name:       "wipchat",
		ShortUsage: "wipchat --key=<key> <subcommand>",
		FlagSet:    rootFlags,
		Options:    []ff.Option{ff.WithEnvVarPrefix("WIPCHAT")},
		Subcommands: []*ffcli.Command{
			{
				Name:      "me",
				ShortHelp: "retrieve user info about current token",
				Exec: func(ctx context.Context, _ []string) error {
					client := wipchat.New(*apiKey)
					viewer, err := client.QueryViewer(ctx)
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
				Exec: func(ctx context.Context, args []string) error {
					body := strings.TrimSpace(strings.Join(args, " "))
					if len(body) < 1 {
						return flag.ErrHelp
					}
					client := wipchat.New(*apiKey)
					todo, err := client.MutateCreateTodo(ctx, body, nil, nil)
					if err != nil {
						return err
					}
					fmt.Println(godev.PrettyJSON(todo))
					return nil
				},
			}, {
				Name:       "done",
				ShortUsage: "wipchat done <lorem ipsum>",
				ShortHelp:  "create a new completed task",
				Exec: func(ctx context.Context, args []string) error {
					body := strings.TrimSpace(strings.Join(args, " "))
					if len(body) < 1 {
						return flag.ErrHelp
					}
					client := wipchat.New(*apiKey)
					now := time.Now()
					todo, err := client.MutateCreateTodo(ctx, body, &now, nil)
					if err != nil {
						return err
					}
					fmt.Println(godev.PrettyJSON(todo))
					return nil
				},
			},
		},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), os.Args[1:])
}
