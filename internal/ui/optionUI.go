package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/nils/jpackageTUI/internal/ui/tvchooser"
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
		} else if opt.GetInputType() == option.Folder || opt.GetInputType() == option.File {
			field = tview.NewDropDown().
				SetLabel(opt.GetOptionName()).
				SetOptions([]string{" "}, nil)

			field.(*tview.DropDown).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() != tcell.KeyTab && event.Key() != tcell.KeyEsc {
					var path string

					if opt.GetInputType() == option.Folder {
						// Original code adapted from https://github.com/AEROGU/tvchooser
						path = tvchooser.DirectoryChooser(UI.app, false)
					} else {
						// Original code adapted from https://github.com/AEROGU/tvchooser
						path = tvchooser.FileChooser(UI.app, false)
					}

					if path == "" {
						path = " "
					}
					field.(*tview.DropDown).SetOptions([]string{path}, nil)
					field.(*tview.DropDown).SetCurrentOption(0)
					field.Blur()
				}

				return event
			})
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
