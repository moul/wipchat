package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/godev"
	"moul.io/wipchat"
)

func productsCommand() *ffcli.Command {
	var limit int
	flags := flag.NewFlagSet("products", flag.ExitOnError)
	flags.IntVar(&limit, "limit", 20, "limit")
	return &ffcli.Command{
		Name:      "products",
		ShortHelp: "products people are making",
		FlagSet:   flags,
		Exec: func(ctx context.Context, _ []string) error {
			client := wipchat.New(apiKey)
			products, err := client.QueryProducts(ctx, limit)
			if err != nil {
				return err
			}
			fmt.Println(godev.PrettyJSON(products))
			return nil
		},
	}
}
