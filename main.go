package main

import (
	"log"

	icmd "go_terminal/internal/cmd"
)

func main() {
	if err := icmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
