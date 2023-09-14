/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

const (
	THREADS   = 2000
	ADDRESSES = 25
)

func fetchData(url string, base int) {
	fmt.Println(base)
	for i := base; i < base+ADDRESSES; i++ {
		url = "https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=" + strconv.Itoa(
			i,
		)
		fmt.Printf("trying %s\n", url)
		cmd := exec.Command("/home/yaron/projects/shoam_bash/course_info", url)
		err := cmd.Run()
		if err == nil {
			fmt.Println("course added")
		}
	}
}

// generateDbCmd represents the generateDb command
var generateDbCmd = &cobra.Command{
	Use:   "generateDb",
	Short: "Generate the db based on the shoam system",
	Long: ` this script will scan the Shoam system for CS courses.
            it will save the data locally in a folder called "db" (see 'course_info' program in the parent folder)
            each course file is of the following form:
            ------- [course #]-[group #]-[lid #] (lid is explained below)
            since the urls do not corresponeds to the course number or group,
            this script will try 50,000 different urls, and will extract the data for every CS course it finds.
            the url is of the form:
            ------ https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=[...]
            this script can scan more addresses by changing the number of ADDRESSES in the global variable above.
            can also change the # of threads running for efficiency.
            results for 8RAM memory and Intel CPU i5-8250U - about 25 minutes. `,

	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		base := 760000

		wg.Add(THREADS)
		url := "https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=" + strconv.Itoa(
			base,
		)

		for i := 0; i < THREADS; i++ {
			go fetchData(url, base)
			base = base + ADDRESSES
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(generateDbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
