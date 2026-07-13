package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"tdocli/task"
)

var addCmd = &cobra.Command{
	Use:   "add <task title>",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t := args[0]
		p := 0
		if v, _ := cmd.Flags().GetBool("high"); v {
			p = 2
		} else if v, _ := cmd.Flags().GetBool("medium"); v {
			p = 1
		} else if v, _ := cmd.Flags().GetBool("low"); v {
			p = 0
		}
		for _, task := range args[1:] {
			t += " " + task
		}
		err := task.Add(t, p)
		if err != nil {
			return err
		}
		fmt.Println("Task added:", t)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("high", "i", false, "Mark task as high priority")
	addCmd.Flags().BoolP("medium", "m", false, "Mark task as medium priority")
	addCmd.Flags().BoolP("low", "l", false, "Mark task as low priority")
}