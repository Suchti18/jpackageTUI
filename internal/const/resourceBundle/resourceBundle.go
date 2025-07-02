package resourceBundle

import (
	"github.com/Xuanwo/go-locale"
	"github.com/nils/jpackageTUI/internal/const/resourceBundle/english"
	"github.com/nils/jpackageTUI/internal/const/resourceBundle/german"
	"golang.org/x/text/language"
	"log"
)

var (
	currLang      language.Tag
	languagePacks map[language.Tag]func(string) string
)

func init() {
	lang, err := locale.Detect()
	if err != nil {
		log.Println(err)
		currLang = language.English
	} else {
		currLang = lang.Parent()
	}

	languagePacks = map[language.Tag]func(string) string{
		language.English: english.GetString,
		language.German:  german.GetString,
	}
}

func SetLang(newLang language.Tag) {
	currLang = newLang
}

func GetLang() language.Tag {
	return currLang
}

func GetString(id string) string {
	if languagePack, exists := languagePacks[currLang]; exists {
		return languagePack(id)
	}

	// Fallback to english
	return english.GetString(id)
}
