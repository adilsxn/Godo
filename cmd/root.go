/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Godo.git",
	Short: "A small cli task management app",
	Long: `A small cli task management app that lets you create,
    list, complete and delete your task . For example:
    godo add <your task description>
    godo list -- lists all the task to be done
    godo list -a -> lists all the tasks
    godo complete <task id> completes the task
    godo delete <task id> deletes the given task`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


