package main

import (
	"fmt"
	"github.com/nils/jpackageTUI/internal/Const/exitCodes"
	"github.com/nils/jpackageTUI/internal/Const/resourceBundle"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/nils/jpackageTUI/internal/ui"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	if !isjpackageInstalled() && (len(os.Args) > 1 && os.Args[1] != "--force") {
		fmt.Println(resourceBundle.GetString("JpackageNotInstalled"))
		os.Exit(exitCodes.JpackageNotInstalled)
	}

	if err := ui.New().Start(); err != nil {
		log.Fatal(err)
	}

	parameters := option.CreateParameters()

	if len(parameters) > 0 {
		runjpackageCommand(parameters)
	}

	os.Exit(exitCodes.Success)
}

func isjpackageInstalled() bool {
	cmd := exec.Command("jpackage")

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
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

	fmt.Print("Command: ")
	fmt.Println(cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(resourceBundle.GetString("SuccessfulExit"))
	}
}
