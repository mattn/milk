package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"log"
	"os"
)

func main() {
	cfg := config{}

	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "remember-the-milk inteface for cli",
	}
	command.Subcommands = []*commander.Command{
		make_cmd_list(cfg),
		make_cmd_task(cfg),
		make_cmd_add(cfg),
		make_cmd_del(cfg),
	}

	err := cfg.load()
	if err != nil {
		log.Fatal(err)
	}

	if cfg["auth_token"] == "" {
		token, err := auth()
		if err != nil {
			log.Fatal(err)
		}
		cfg["auth_token"] = token
		cfg.save()
	}

	err = command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
