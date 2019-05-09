package main

import (
	"fmt"
	"log"

	"github.com/figglewatts/daybookr/internal/daybookr"
)

func run() error {
	_, err := daybookr.CreatePost("template.txt", "first-entry.md")
	if err != nil {
		return fmt.Errorf("Could not open post: %v", err)
	}
	//fmt.Printf(output)

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}
