package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strconv"
	"github.com/Prominence673/tdocli/task"
)

var editCmd = &cobra.Command{
	Use:   "edit <task id>",
	Short: "Edit a task",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		t := args[1]
		for _, task := range args[2:] {
			t += " " + task
		}
		err = task.Edit(t, id)
		if err != nil {
			return err
		}
		fmt.Println("Task edited:", t)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}