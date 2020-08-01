package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func userCommand() *ffcli.Command {
	var opts wipchat.QueryUserOptions
	opts.ApplyDefaults()
	flags := flag.NewFlagSet("user", flag.ExitOnError)
	flags.BoolVar(&opts.TodosCompleted, "todos-completed", opts.TodosCompleted, "todos completed")
	flags.IntVar(&opts.TodosLimit, "todos-limit", opts.TodosLimit, "todos limit")
	flags.IntVar(&opts.TodosOffset, "todos-offset", opts.TodosOffset, "todos offset")
	flags.StringVar(&opts.TodosFilter, "todos-filter", opts.TodosFilter, "todos filter")
	flags.StringVar(&opts.TodosOrder, "todos-order", opts.TodosOrder, "todos order")
	flags.IntVar(&opts.AvatarSize, "avatar-size", opts.AvatarSize, "avatar size")
	flags.StringVar(&opts.UserID, "id", opts.UserID, "user ID")
	flags.StringVar(&opts.Username, "username", opts.Username, "username")

	return &ffcli.Command{
		Name:       "user",
		ShortHelp:  "retrieve user info",
		ShortUsage: "wipchat user --id=1780\n  wipchat user --username=manfred",
		FlagSet:    flags,
		Exec: func(ctx context.Context, _ []string) error {
			if opts.UserID == "" && opts.Username == "" {
				return flag.ErrHelp
			}
			if opts.UserID != "" && opts.Username != "" {
				return flag.ErrHelp
			}

			client := wipchat.New(apiKey)
			user, err := client.QueryUser(ctx, &opts)
			if err != nil {
				return err
			}
			fmt.Println(godev.PrettyJSON(user))
			return nil
		},
	}
}
