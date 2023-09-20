/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"shoam/utils"
)

var calcCmd = &cobra.Command{
	Use:   "calc <courses codes...> ",
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
		verboseFlag, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println(err)
		}
		fileFlag, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
		}
		if fileFlag != "" {
			handleFileFlag(fileFlag, verboseFlag)
			return
		}
		if len(args) == 0 {
			handleStdin(verboseFlag)
			return
		} else {
			handleCoursesArgs(args, verboseFlag)
		}
	},
}

func handleCoursesArgs(courses []string, verboseFlag bool) {
	total := 0.0
	var v string
	if verboseFlag == true {
		v = "true"
	} else {
		v = "false"
	}
	for _, c := range courses {
		pointsCmd := exec.Command(utils.SCRIPTS_DIR+"/get_points", c, v)
		points, err := pointsCmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		lines := strings.Split(string(points), "\n")
		p_str := strings.Split(lines[len(lines)-2], ",")[0]
		p, err := strconv.ParseFloat(p_str, 8)
		total += p
		// fmt.Println(p)
		if verboseFlag == true {
			fmt.Println(string(points))
		}
	}
	fmt.Printf("total points (%d courses): %.2f\n", len(courses), total)
}

func handleStdin(verboseFlag bool) {
	fmt.Printf("stdin\n")
}

func handleFileFlag(file string, verboseFlag bool) {
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
	calcCmd.Flags().BoolP("verbose", "v", false, "verbose output")
}
