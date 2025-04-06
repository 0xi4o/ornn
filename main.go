package main

import (
	"log"
	"os"

	"github.com/0xi4o/ornn/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
		os.Exit(1)
	}
}
