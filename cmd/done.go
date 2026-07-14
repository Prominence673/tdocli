package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/Prominence673/tdocli/task"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		err = task.MarkCompleted(id)
		if err != nil {
			return err
		}
		fmt.Printf("Task %d marked as done\n", id)
		return nil
	},
}

func init(){
	rootCmd.AddCommand(doneCmd)
}