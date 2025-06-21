package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/rivo/tview"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type OptionUI struct {
	Form    *tview.Form
	Fields  map[*option.Option]tview.Primitive
	Summary map[string]string
}

func NewOptionsUI(options []*option.Option) *OptionUI {
	optionUI := &OptionUI{
		Form:    tview.NewForm(),
		Fields:  make(map[*option.Option]tview.Primitive),
		Summary: make(map[string]string),
	}

	for _, opt := range options {
		var field tview.FormItem

		if len(opt.GetPossibleOptions()) > 0 {
			field = tview.NewDropDown().
				SetLabel(opt.GetOptionName()).
				SetLabelColor(tcell.NewHexColor(0xd4c57b)).
				SetOptions(opt.GetPossibleOptions(), func(text string, index int) {})
		} else if opt.HasNoParameter() {
			field = tview.NewCheckbox().
				SetLabel(opt.GetOptionName()).
				SetLabelColor(tcell.NewHexColor(0xd4c57b))
		} else if opt.GetInputType() == option.File {
			field = tview.NewInputField().
				SetLabel(opt.GetOptionName()).
				SetLabelColor(tcell.NewHexColor(0xd4c57b)).
				SetText("").
				SetAutocompleteFunc(filePathAutocomplete()).SetPlaceholder("(Enter an filepath)")
		} else {
			field = tview.NewInputField().
				SetLabel(opt.GetOptionName()).
				SetLabelColor(tcell.NewHexColor(0xd4c57b)).
				SetText("")
		}

		if opt.IsOptional() && !opt.HasNoParameter() {
			checkbox := tview.NewCheckbox()
			checkbox.SetLabel(fmt.Sprintf("Include <%s>?", opt.GetOptionName()))

			field.SetDisabled(true)

			checkbox.SetChangedFunc(func(checked bool) {
				if checked {
					field.SetDisabled(false)
				} else {
					field.SetDisabled(true)
				}
			})

			optionUI.Form.AddFormItem(checkbox)
		}

		optionUI.Form.AddFormItem(field)
		optionUI.Fields[opt] = field
	}

	optionUI.Form.
		SetFieldBackgroundColor(tcell.NewHexColor(0x343d46)).
		SetButtonBackgroundColor(tcell.NewHexColor(0x343d46)).
		SetBackgroundColor(tcell.NewHexColor(0x2b303b)).
		SetBorder(true).
		SetBorderColor(tcell.NewHexColor(0x8fa1b3)).
		SetTitle("jpackageTUI")

	return optionUI
}

func (optionUI *OptionUI) GetPrimitive() tview.Primitive {
	return optionUI.Form
}

func (optionUI *OptionUI) getData() map[*option.Option]string {
	var data = make(map[*option.Option]string)

	for opt, field := range optionUI.Fields {
		switch field.(type) {
		case *tview.InputField:
			inputField := field.(*tview.InputField)
			data[opt] = inputField.GetText()
		case *tview.DropDown:
			dropdown := field.(*tview.DropDown)
			_, data[opt] = dropdown.GetCurrentOption()
		case *tview.Checkbox:
			checkbox := field.(*tview.Checkbox)
			if checkbox.IsChecked() {
				data[opt] = ""
			}
		}
	}

	return data
}

func (optionUI *OptionUI) addBackButton() {
	optionUI.Form.AddButton("Back", func() {
		lastPage()
	})
}

func (optionUI *OptionUI) addNextButton() {
	optionUI.Form.AddButton("Next", func() {
		nextPage()
	})
}

func (optionUI *OptionUI) addFinishButton() {
	optionUI.Form.AddButton("Finish", func() {
		finish()
	})
}

func filePathAutocomplete() func(currentText string) []string {
	return func(currentText string) []string {
		if currentText == "" {
			return nil
		} else if currentText == "." {
			currDir, err := filepath.Abs(currentText)
			if err == nil {
				currentText = filepath.Join(currDir)
			}
		} else if strings.HasPrefix(currentText, "~/") {
			homeDir, err := os.UserHomeDir()
			if err == nil {
				currentText = filepath.Join(homeDir, currentText[2:])
			}
		}

		currentText = strings.ReplaceAll(currentText, "/", "\\")

		dir := filepath.Dir(currentText)
		prefix := filepath.Base(currentText)

		if _, err := os.Stat(dir); err != nil {
			return nil
		}

		entries := getMatchingEntries(dir, prefix)

		for i, entry := range entries {
			if strings.HasSuffix(currentText, string(os.PathSeparator)) {
				entries[i] = filepath.Join(currentText, entry)
			} else if dir == "." && !strings.Contains(currentText, string(os.PathSeparator)) {
				entries[i] = entry
			} else {
				base := strings.TrimSuffix(currentText, prefix)
				if !strings.HasSuffix(base, string(os.PathSeparator)) {
					base += string(os.PathSeparator)
				}
				entries[i] = base + entry
			}
		}

		return entries
	}
}

func getMatchingEntries(dir, prefix string) []string {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var matches []string

	for _, entry := range dirEntries {
		name := entry.Name()

		if strings.HasPrefix(name, ".") && prefix != "." && !strings.HasPrefix(prefix, ".") {
			continue
		}

		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			if entry.IsDir() {
				name += string(os.PathSeparator)
			}
			matches = append(matches, name)
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		isDir1 := strings.HasSuffix(matches[i], string(os.PathSeparator))
		isDir2 := strings.HasSuffix(matches[j], string(os.PathSeparator))

		if isDir1 && !isDir2 {
			return true
		}
		if !isDir1 && isDir2 {
			return false
		}

		return strings.ToLower(matches[i]) < strings.ToLower(matches[j])
	})

	return matches
}
