package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"tdocli/task"
	"strconv"
)

var undoCmd = &cobra.Command{
	Use:   "undo <id>",
	Short: "Undo a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		err = task.Undo(id)
		if err != nil {
			return err
		}
		fmt.Println("Task undone:", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}