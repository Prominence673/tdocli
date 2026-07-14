package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/Prominence673/tdocli/task"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := task.Clear()
		if err != nil {
			return err
		}
		fmt.Println("All tasks cleared")
		return nil
	},
}

func init(){
	rootCmd.AddCommand(clearCmd)
}