package main

import (
	"github.com/gonuts/commander"
	"log"
	"os"
)

func make_cmd_add(cfg config) *commander.Command {
	cmd_add := func(cmd *commander.Command, args []string) error {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}
		tl, err := timeline(cfg)
		if err != nil {
			log.Fatal(err)
		}
		_, err = api(api_param{
			"method":     "rtm.tasks.add",
			"auth_token": cfg["auth_token"],
			"timeline":   tl,
			"name":       args[0],
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_add,
		UsageLine: "add [options] [name]",
		Short:     "add task",
	}
}
