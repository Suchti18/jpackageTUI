package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/rivo/tview"
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
		if len(opt.GetPossibleOptions()) > 0 {
			dropdown := tview.NewDropDown()
			dropdown.SetLabel(opt.GetOptionName())
			dropdown.SetOptions(opt.GetPossibleOptions(), func(text string, index int) {

			})

			optionUI.Fields[opt] = dropdown
			optionUI.Form.AddFormItem(dropdown)
		} else {
			inputField := tview.NewInputField()
			inputField.SetLabel(opt.GetOptionName())
			inputField.SetText("")
			optionUI.Fields[opt] = inputField
			optionUI.Form.AddFormItem(inputField)

		}
	}

	optionUI.Form.
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetButtonBackgroundColor(tcell.ColorDarkBlue)

	return optionUI
}

func (optionUI *OptionUI) GetPrimitive() tview.Primitive {
	return optionUI.Form
}
