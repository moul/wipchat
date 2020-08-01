package main

import (
	"context"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func todosCommand() *ffcli.Command {
	return &ffcli.Command{
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
	}
}
