package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"some_tests/CLI/db"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete. Why not take a break?")
			return
		}
		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
			//fmt.Printf("%d. %s \n", i+1, task.Value)
			if(task.Id==1){

			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}