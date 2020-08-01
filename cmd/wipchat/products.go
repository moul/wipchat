package main

import (
	"context"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func productsCommand() *ffcli.Command {
	return &ffcli.Command{
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
	}
}
