package test

import (
	"github.com/nils/jpackageTUI/internal/option"
	"testing"
)

func TestCorrectConversion(t *testing.T) {
	actual := option.ConvertToCommand()
	expected := "jpackage"

	if actual != expected {
		t.Errorf("Command got wrong converted. Got <%s>, expected <%s>", actual, expected)
	}
}
