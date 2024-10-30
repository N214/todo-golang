package cmd

import (
	"flag"
	"fmt"
	"log"
	"todo/todo"
)

func DeleteTasks(todos *todo.Todos, args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// If no --id=1 flag defined, todoID will default to 0
	// but if --id is present but didn't set any value, an error will be thrown
	deleteId := deleteCmd.Int("id", 0, "The id of todo to be deleted")

	deleteCmd.Parse(args)

	// the approach here deletes the records from the slice and restore in the json
	// an alternative approach could be to read the json data and just delete the target line
	err := todos.Delete(*deleteId)
	if err != nil {
		log.Fatal(err)
	}

	err = todos.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	todos.Print(2, "")
	fmt.Println("Todo item deleted successfully.")
}