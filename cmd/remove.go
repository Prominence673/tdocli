package cmd

import (
	"fmt"
	"strconv"
	"github.com/Prominence673/tdocli/task"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <id|title>",
	Short: "Remove a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title, _ := cmd.Flags().GetBool("title")
		var err error
		if title {
			err = task.RemoveByTask(args[0])
		} else {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			err = task.RemoveById(id)
			if err != nil {
				return err
			}
		}
		fmt.Printf("Task %s removed\n", args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolP("title", "t", false, "Remove tasks by title")
}
