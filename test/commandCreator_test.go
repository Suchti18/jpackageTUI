package test

import (
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorrectConversion(t *testing.T) {
	option.ClearRepo()
	option1 := option.NewOption("", "", "-t", option.Mac, false, false, []string{}, false)
	option.AddToRepo(option1, "app-image")
	option2 := option.NewOption("", "", "-m", option.Mac, false, false, []string{}, false)
	option.AddToRepo(option2, "This is a test")

	actual := option.CreateParameters()
	expected := []string{"-t app-image", " -m \"This is a test\""}

	assert.Subsetf(t, actual, expected, "The expected parameters and the actual parameters are not equal")
}

func TestCorrectConversionWithAOptionWithNoParameters(t *testing.T) {
	option.ClearRepo()
	option1 := option.NewOption("", "", "-t", option.Mac, false, false, []string{}, false)
	option.AddToRepo(option1, "app-image")
	option2 := option.NewOption("", "", "-m", option.Mac, true, true, []string{}, true)
	option.AddToRepo(option2, "This is a test")

	actual := option.CreateParameters()
	expected := []string{"-t app-image", " -m"}

	assert.Subsetf(t, actual, expected, "The expected parameters and the actual parameters are not equal")
}
