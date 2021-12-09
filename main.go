package main

import (
	"log"

	"github.com/diy0663/go-cmd/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
