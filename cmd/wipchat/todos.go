package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func todosCommand() *ffcli.Command {
	var opts wipchat.QueryTodosOptions
	opts.ApplyDefaults()
	flags := flag.NewFlagSet("todos", flag.ExitOnError)
	flags.BoolVar(&opts.TodosCompleted, "completed", opts.TodosCompleted, "todos completed")
	flags.IntVar(&opts.TodosLimit, "limit", opts.TodosLimit, "todos limit")
	flags.StringVar(&opts.TodosFilter, "filter", opts.TodosFilter, "todos filter")

	return &ffcli.Command{
		Name:      "todos",
		ShortHelp: "todos posted by makers",
		FlagSet:   flags,
		Exec: func(ctx context.Context, _ []string) error {
			client := wipchat.New(apiKey)
			todos, err := client.QueryTodos(ctx, &opts)
			if err != nil {
				return err
			}
			fmt.Println(godev.PrettyJSON(todos))
			return nil
		},
	}
}
