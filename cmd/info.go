/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [course code]",
	Short: "Get information for the input course",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please provide course code")
			return
		}

		course_code := args[0]
		groupTerm, _ := cmd.Flags().GetString("group")
		getInfoCommand := exec.Command(
			"/home/yaron/projects/shoam/get_course_info",
			course_code,
			groupTerm,
		)
		output, err := getInfoCommand.CombinedOutput()
		if err != nil {
			return
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().String("group", "01", "spesific course group")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
