package args

import (
	"os"
	"strings"
)

const (
	ForceArg = "--force"
	AllArg   = "--all"
)

func HasArg(arg string) bool {
	if len(os.Args) == 1 {
		return false
	}

	for _, givenArg := range os.Args {
		if strings.EqualFold(arg, givenArg) {
			return true
		}
	}
	return false
}
