package cmd

import (
	"github.com/spf13/cobra"
	"github.com/charmbracelet/lipgloss"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "A CLI tool for managing tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if version, _ := cmd.Flags().GetBool("version"); version {
			var estilo = lipgloss.NewStyle().
			    Foreground(lipgloss.Color("202")).
			    Bold(true)
			const logo =
` _______  _____   ____   _____ _      _____
|__   __||  __ \ / __ \ / ____| |    |_   _|
   | |   | |  | | |  | | |    | |      | |
   | |   | |  | | |  | | |    | |      | |
   | |   | |__| | |__| | |____| |____ _| |_
   |_|   |_____/ \____/ \_____|______|_____|`
   fmt.Println(estilo.Render(logo + "\n\nVersion: 1.0.0\nA CLI tool for managing tasks"))
			return
		}
		cmd.Help()
	},
}
func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}



