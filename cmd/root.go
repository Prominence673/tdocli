package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Prominence673/tdocli/cmd/assets"
	"fmt"
	"os"
	"github.com/Prominence673/tdocli/task"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "A CLI tool for managing tasks",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		if file != "" {
			if err := task.SetFilename(file); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if version, _ := cmd.Flags().GetBool("version"); version {
			assets.RenderLogo()
			return
		}
		cmd.Help()
	},
}
func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
	rootCmd.PersistentFlags().StringP("file", "f", "", "Task file")
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}



