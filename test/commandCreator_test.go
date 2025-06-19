package test

import (
	"github.com/nils/jpackageTUI/internal/option"
	"testing"
)

func TestCorrectConversion(t *testing.T) {
	option1 := option.NewOption("", "", "-t", option.Mac, false, false, []string{}, false)
	option.AddToRepo(option1, "app-image")
	option2 := option.NewOption("", "", "-m", option.Mac, false, false, []string{}, false)
	option.AddToRepo(option2, "This is a test")

	actual := option.CreateParameters()
	expected := []string{"-t app-image", "-m \"This is a test\""}

	if slicesEqual(actual, expected) {
		t.Errorf("Command got wrong converted. Got <%s>, expected <%s>", actual, expected)
	}
}

func TestCorrectConversionWithAOptionWithNoParameters(t *testing.T) {
	option1 := option.NewOption("", "", "-t", option.Mac, false, false, []string{}, false)
	option.AddToRepo(option1, "app-image")
	option2 := option.NewOption("", "", "-m", option.Mac, true, true, []string{}, true)
	option.AddToRepo(option2, "This is a test")

	actual := option.CreateParameters()
	expected := []string{"-t app-image", "-m"}

	if slicesEqual(actual, expected) {
		t.Errorf("Command got wrong converted. Got <%s>, expected <%s>", actual, expected)
	}
}

func slicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
