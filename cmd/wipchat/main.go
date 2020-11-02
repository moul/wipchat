package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	ff "github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/motd"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

var (
	apiKey string
	debug  bool
)

func run(_ []string) error {
	flags := flag.NewFlagSet("root", flag.ExitOnError)
	flags.StringVar(&apiKey, "key", "", "Your private API key from https://wip.co/api")
	flags.BoolVar(&debug, "debug", false, "More verbose output")

	root := &ffcli.Command{
		Name:       "wipchat",
		ShortUsage: "wipchat --key=<key> <subcommand>",
		FlagSet:    flags,
		Options:    []ff.Option{ff.WithEnvVarPrefix("WIPCHAT")},
		Subcommands: []*ffcli.Command{
			meCommand(),
			todoCommand(),
			doneCommand(),
			todosCommand(),
			productsCommand(),
			userCommand(),
		},
		Exec: func(_ context.Context, _ []string) error {
			fmt.Fprintln(os.Stderr, motd.Default())
			return flag.ErrHelp
		},
	}

	return root.ParseAndRun(context.Background(), os.Args[1:])
}
