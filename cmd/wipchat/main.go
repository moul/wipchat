package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

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
	fs := flag.NewFlagSet("repeat", flag.ExitOnError)
	apiKey := fs.String("key", "", "Your private API key from https://wip.chat/api")

	root := &ffcli.Command{
		FlagSet: fs,
		Options: []ff.Option{ff.WithEnvVarPrefix("WIPCHAT")},
		Exec: func(ctx context.Context, _ []string) error {
			client := wipchat.New(*apiKey)
			viewer, err := client.QueryViewer(ctx)
			fmt.Println(godev.PrettyJSON(viewer), err)
			return nil
		},
	}

	return root.ParseAndRun(context.Background(), os.Args[1:])
}
