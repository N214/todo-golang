package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"todo/todo"
)

func UpdateTask(todos *todo.Todos, args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateTask := updateCmd.String("task", "", "The content of the updated todo item")
	updateId := updateCmd.Int("id", 0, "The id of todo to be updated")
	updateDone := updateCmd.Int("done", 2, "Is the updated task done?")
	updateCat := updateCmd.String("cat", "Uncategorized", "The category of the updated todo item")

	updateCmd.Parse(args)

	if *updateId == 0 {
		fmt.Println("Error: the --id flag is required for the 'update' subcommand.")
		os.Exit(1)
	}

	// In memory update
	err := todos.Update(*updateId, *updateTask, *updateCat, *updateDone)
	if err != nil {
		log.Fatal(err)
	}

	// Save to disk
	err = todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}
	todos.Print(2, "")
	fmt.Println("Todo item updated successfully.")

}