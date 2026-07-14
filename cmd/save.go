package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/Prominence673/tdocli/task"
)

var expCmd = &cobra.Command{
	Use:   "export <src>",
	Short: "Export tasks to disk",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := task.Export(args[0])
		if err != nil {
			return err
		}
		fmt.Println("Tasks saved successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(expCmd)
}