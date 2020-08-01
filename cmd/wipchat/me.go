package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func meCommand() *ffcli.Command {
	var opts wipchat.QueryViewerOptions
	opts.ApplyDefaults()
	flags := flag.NewFlagSet("me", flag.ExitOnError)
	flags.BoolVar(&opts.TodosCompleted, "todos-completed", opts.TodosCompleted, "todos completed")
	flags.IntVar(&opts.TodosLimit, "todos-limit", opts.TodosLimit, "todos limit")
	flags.IntVar(&opts.TodosOffset, "todos-offset", opts.TodosOffset, "todos offset")
	flags.StringVar(&opts.TodosFilter, "todos-filter", opts.TodosFilter, "todos")
	flags.StringVar(&opts.TodosOrder, "todos-order", opts.TodosOrder, "todos order")
	flags.IntVar(&opts.AvatarSize, "avatar-size", opts.AvatarSize, "avatar size")

	return &ffcli.Command{
		Name:      "me",
		ShortHelp: "retrieve user info about current token",
		FlagSet:   flags,
		Exec: func(ctx context.Context, _ []string) error {
			client := wipchat.New(apiKey)
			viewer, err := client.QueryViewer(ctx, &opts)
			if err != nil {
				return err
			}
			fmt.Println(godev.PrettyJSON(viewer))
			return nil
		},
	}
}
