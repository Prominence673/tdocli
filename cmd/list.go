package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"tdocli/task"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks := []task.Task{}
		if done, _ := cmd.Flags().GetBool("done"); done {
			tasks = task.GetCompleted()
		}
		if pending, _ := cmd.Flags().GetBool("pending"); pending {
			tasks = task.GetPending()
		}
		if len(tasks) == 0 {
			tasks = task.GetAll()
		}
		for _, t := range tasks {
			if t.Completed {
				fmt.Printf("%d.%s ● %s %s\n", t.Id, Green, t.Title, Reset)
			} else {
				switch t.Priority {
				case 0:
					fmt.Printf("%d. ○ %s\n", t.Id, t.Title)
				case 1:
					fmt.Printf("%d.%s ○ %s %s\n", t.Id, Yellow, t.Title, Reset)
				case 2:
					fmt.Printf("%d.%s ○ %s %s\n", t.Id, Red, t.Title, Reset)
				}
			}
			
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("done", "d", false, "List completed tasks")
	listCmd.Flags().BoolP("pending", "p", false, "List pending tasks")
}
