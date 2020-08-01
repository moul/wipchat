package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func todoCommand() *ffcli.Command {
	var attachPaths stringSlice
	flags := flag.NewFlagSet("todo", flag.ExitOnError)
	flags.Var(&attachPaths, "attach", "attachment paths or URLs")

	return &ffcli.Command{
		Name:       "todo",
		ShortUsage: "wipchat todo <lorem ipsum>",
		ShortHelp:  "create a new todo task",
		FlagSet:    flags,
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
	}
}

func doneCommand() *ffcli.Command {
	var attachPaths stringSlice
	flags := flag.NewFlagSet("todo", flag.ExitOnError)
	flags.Var(&attachPaths, "attach", "attachment paths or URLs")

	return &ffcli.Command{
		Name:       "done",
		ShortUsage: "wipchat done <lorem ipsum>",
		ShortHelp:  "create a new completed task",
		FlagSet:    flags,
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
	}
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
