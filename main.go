package main

import (
	"github.com/OldBigBuddha/socialnet/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to execute account service; %w", err)
	}
}
