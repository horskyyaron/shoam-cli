/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc <courses codes...>",
	Short: "calculates total credit points for a list of courses",
	Long: `Can be used in multiple ways:
    1. shoam calc 89230 -> will return the credit for one course
    2. shoam calc 89230 89220 -> will return the credit for two courses (add more as you wish)
    3. shoam calc -f file -> will calculate total points for a file of a list of courses.
        3.1 each line in the file should be a course code. do NOT include group number.
        e.g.
        file:
        89230
        89220
    4. cat file | shoam calc -> should be in the same format as in the file usage.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
		}
		if fileFlag != "" {
			handleFileFlag(fileFlag)
			return
		}
		if len(args) == 0 {
			handleStdin()
			return
		} else {
			handleCoursesArgs(args)
		}
	},
}

func handleCoursesArgs(args []string) {
	fmt.Printf("number of args %d\n", len(args))
}

func handleStdin() {
	fmt.Printf("stdin\n")
}

func handleFileFlag(file string) {
	fmt.Printf("file flag on, file name: %s\n", file)
}

func init() {
	rootCmd.AddCommand(calcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	calcCmd.Flags().StringP("file", "f", "", "-f <course file>")
}
