/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/package main

import (
	"fmt"
	"os"

	"shoam-cli/cmd"
)

func main() {
	if os.Getenv("SHOAM_DIR") == "" {
		fmt.Println(
			"please export the env variable SHOAM_DIR with the value of the path to 'shoam' directory.",
		)
		os.Exit(1)
	}
	cmd.Execute()
}
