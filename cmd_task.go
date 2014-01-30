package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"github.com/mattn/go-scan"
	"log"
)

func make_cmd_task(cfg config) *commander.Command {
	cmd_task := func(cmd *commander.Command, args []string) error {
		res, err := api(api_param{
			"method":     "rtm.tasks.getList",
			"auth_token": cfg["auth_token"],
		})
		if err != nil {
			log.Fatal(err)
		}
		var list []interface{}
		err = scan.ScanTree(res, "/rsp/tasks/list", &list)
		if err != nil {
			log.Fatal(err)
		}
		var tasks []map[string]interface{}
		for _, iter := range list {
			taskseries, _ := iter.(map[string]interface{})["taskseries"]
			if taskseries != nil {
				if task, ok := taskseries.(map[string]interface{}); ok {
					tasks = append(tasks, task)
				} else if iter, ok := taskseries.([]interface{}); ok {
					for _, elem := range iter {
						task := elem.(map[string]interface{})
						tasks = append(tasks, task)
					}
				}
			}
		}
		for _, task := range tasks {
			completed := task["task"].(map[string]interface{})["completed"]
			if completed != nil && completed.(string) != "" {
				continue
			}
			fmt.Printf("%10v: %s\n", task["id"], task["name"])
		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_task,
		UsageLine: "task [options]",
		Short:     "show task index",
	}
}
