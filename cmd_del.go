package main

import (
	"errors"
	"github.com/gonuts/commander"
	//"log"
	//"os"
)

func make_cmd_del(cfg config) *commander.Command {
	cmd_del := func(cmd *commander.Command, args []string) error {
		/*
		if len(args) != 2 {
			cmd.Usage()
			os.Exit(1)
		}
		tl, err := timeline(cfg)
		if err != nil {
			log.Fatal(err)
		}
		_, err = api(api_param{
			"method":     "rtm.tasks.delete",
			"auth_token": cfg["auth_token"],
			"timeline":   tl,
			"list_id":    args[0],
			"list_id":    args[1],
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
		*/
		return errors.New("Not implemented yet")
	}

	return &commander.Command{
		Run:       cmd_del,
		UsageLine: "del [options] [name]",
		Short:     "delete task",
	}
}
