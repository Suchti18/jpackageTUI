package main

import (
	"fmt"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/nils/jpackageTUI/internal/ui"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	if err := ui.New().Start(); err != nil {
		log.Fatal(err)
	}

	runjpackageCommand(option.CreateParameters())

	os.Exit(0)
}

func runjpackageCommand(parameters []string) {
	var splittedParameters []string
	re := regexp.MustCompile(`"([^"]*)"|(\S+)`)

	for _, item := range parameters {
		matches := re.FindAllStringSubmatch(item, -1)
		for _, match := range matches {
			if match[1] != "" {
				splittedParameters = append(splittedParameters, match[1])
			} else if match[2] != "" {
				splittedParameters = append(splittedParameters, match[2])
			}
		}
	}

	cmd := exec.Command("jpackage", splittedParameters...)

	fmt.Print("Args: ")
	fmt.Println(cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Command finished successfully")
	}
}
