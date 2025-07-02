package resourceBundle

import (
	"github.com/nils/jpackageTUI/internal/const/resourceBundle"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestCorrectLang(t *testing.T) {
	resourceBundle.SetLang(language.German)

	actualLang := resourceBundle.GetLang()
	expectedLang := language.German

	assert.Equal(t, actualLang, expectedLang)

	actual := resourceBundle.GetString("Finish")
	expected := "Fertig"

	assert.Equal(t, expected, actual)
}
