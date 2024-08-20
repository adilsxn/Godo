/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all the tasks to be done",
	Long: `Displays all the task to be done, if given
    -a flag display all the tasks independent of status 
    For example:
    godo list -- shows tasks to be done
    godo list -a -- shows all the tasks created`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
