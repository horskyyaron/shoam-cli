/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"shoam-cli/utils"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   `search "term"`,
	Short: "Searching for a course using a string pattern",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please provide a search term")
			return
		}

		term := args[0]
		// groupTerm, _ := cmd.Flags().GetString("group")
		getInfoCommand := exec.Command(
			utils.SCRIPTS_DIR+"/search",
			term,
		)
		output, err := getInfoCommand.CombinedOutput()
		if err != nil {
			return
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
