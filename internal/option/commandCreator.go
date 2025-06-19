package option

import "strings"

func CreateParameters() []string {
	first := true
	parameters := []string{}

	for option, value := range GetRepo() {
		var newOption string

		if len(value) > 0 || option.HasNoParameter() {
			newOption = option.GetOptionCommand()

			if !first {
				newOption = " " + newOption

			} else {
				first = false
			}

			if strings.Contains(value, " ") {
				newOption = newOption + " \"" + value + "\""
			} else {
				newOption = newOption + " " + value
			}

			parameters = append(parameters, newOption)
		}
	}

	return parameters
}
