package main

import (
	"github.com/nils/jpackageTUI/internal/ui"
	"log"
)

func main() {
	if err := ui.New().Start(); err != nil {
		log.Fatal(err)
	}
}
