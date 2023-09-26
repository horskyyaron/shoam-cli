/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
    "shoam-cli/utils"

	"github.com/spf13/cobra"
)

const (
	ROUTINES = 100
)

func fetchData(lids []string) {
	for _, lid := range lids {
		url := "https://shoham.biu.ac.il/BiuCoursesViewer/ENCourseDetails.aspx?lid=" + lid
		cmd := exec.Command(utils.SCRIPTS_DIR + "/course_info", url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("there was a problem with 'course_info'", url, err)
		}

	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// create represents the generateDb command
var create = &cobra.Command{
	Use:   "create",
	Short: "fetches the courses data and creates a db locally.",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		lids := getLids()
		addresses := len(lids) / ROUTINES
		base := 0
		fmt.Printf("# Routines: %d\n# Addresses/routine %d\n", ROUTINES, addresses)
		fmt.Println("fetching data...")
		go spinner(300 * time.Millisecond)
		wg.Add(ROUTINES)
		for i := 0; i < ROUTINES-1; i++ {
			go func(start, end int) {
				defer wg.Done()
				fetchData(lids[start:end])
			}(base, base+addresses)
			base = base + addresses
		}

		go func(start, end int) {
			defer wg.Done()
			fetchData(lids[start:end])
		}(base, len(lids))
		wg.Wait()
		fmt.Println("\rdb created!")

		genCache := exec.Command(utils.SCRIPTS_DIR + "/gen_cache")
		err := genCache.Run()
		if err != nil {
			fmt.Println("Couldn't create cache", err)
		}
		fmt.Println("\rcreated cache file")
	},
}

func getLids() []string {
	file, err := os.Open(utils.LINKS_DIR + "/pages_ids")
	if err != nil {
		fmt.Println("error in reading the file")
		fmt.Println(err)
		return nil
	}

	lids := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lids = append(lids, scanner.Text())
	}
	return lids
}

func init() {
	dbCmd.AddCommand(create)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateDbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateDbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
