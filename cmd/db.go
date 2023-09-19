/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "creates and updates the courses db based on the Shoam system of BIU",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("db called")
	// },
}

func init() {
	rootCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
