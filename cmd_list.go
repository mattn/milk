package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"github.com/mattn/go-scan"
	"log"
)

func make_cmd_list(cfg config) *commander.Command {
	cmd_list := func(cmd *commander.Command, args []string) error {
		res, err := api(api_param{
			"method":     "rtm.lists.getList",
			"auth_token": cfg["auth_token"],
		})
		if err != nil {
			log.Fatal(err)
		}
		var list []interface{}
		err = scan.ScanTree(res, "/rsp/lists/list", &list)
		if err != nil {
			log.Fatal(err)
		}
		for _, iter := range list {
			item, _ := iter.(map[string]interface{})
			if item != nil {
				fmt.Printf("%10v: %s\n", item["id"], item["name"])
			}
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_list,
		UsageLine: "list [options]",
		Short:     "show list index",
	}
}
