/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Sets the given task status to done",
	Long: `Sets the given task status to done if not already. 
    For example:
    godo complete <task id>
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
	},
}

// var t = Task{
// }
func init() {
	rootCmd.AddCommand(completeCmd)
}
