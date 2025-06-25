package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/Const/Colors"
	"github.com/nils/jpackageTUI/internal/Const/resourceBundle"
	"github.com/nils/jpackageTUI/internal/option"
	"github.com/nils/jpackageTUI/internal/ui/tvchooser"
	"github.com/rivo/tview"
	"os"
)

type OptionUI struct {
	Form   *tview.Form
	Fields map[*option.Option]tview.Primitive
}

func NewOptionsUI(options []*option.Option) *OptionUI {
	optionUI := &OptionUI{
		Form:   tview.NewForm(),
		Fields: make(map[*option.Option]tview.Primitive),
	}

	for _, opt := range options {
		var field tview.FormItem

		if len(opt.GetPossibleOptions()) > 0 {
			field = tview.NewDropDown().
				SetLabel(opt.GetOptionName()).
				SetOptions(opt.GetPossibleOptions(), func(text string, index int) {})
		} else if opt.HasNoParameter() {
			field = tview.NewCheckbox().
				SetLabel(opt.GetOptionName())
		} else if opt.GetInputType() == option.Folder || opt.GetInputType() == option.File {
			field = tview.NewDropDown().
				SetLabel(opt.GetOptionName()).
				SetOptions([]string{" "}, nil)

			field.(*tview.DropDown).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyEnter {
					var path string

					currDir, _ := os.Getwd()
					if opt.GetInputType() == option.Folder {
						// Original code adapted from https://github.com/AEROGU/tvchooser
						path = tvchooser.DirectoryChooser(UI.app, false,
							currDir+"|"+resourceBundle.GetString("CurrentDir"))
					} else {
						// Original code adapted from https://github.com/AEROGU/tvchooser
						path = tvchooser.FileChooser(UI.app, false,
							currDir+"|"+resourceBundle.GetString("CurrentDir"))
					}

					if path == "" {
						path = " "
					}
					field.(*tview.DropDown).SetOptions([]string{path}, nil)
					field.(*tview.DropDown).SetCurrentOption(0)
				} else if event.Key() == tcell.KeyTab {
					return event
				}

				return nil
			})
		} else {
			field = tview.NewInputField().
				SetLabel(opt.GetOptionName()).
				SetText("")
		}

		if opt.IsOptional() && !opt.HasNoParameter() {
			checkbox := tview.NewCheckbox().
				SetLabel(fmt.Sprintf(resourceBundle.GetString("Include"), opt.GetOptionName()))

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
		SetFieldBackgroundColor(Colors.FieldBackgroundColor).
		SetButtonBackgroundColor(Colors.ButtonBackgroundColor).
		SetBackgroundColor(Colors.BackgroundColor).
		SetBorder(true).
		SetBorderColor(Colors.BorderColor).
		SetTitle("jpackageTUI")

	optionUI.Form.SetLabelColor(Colors.LabelColor)

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
	optionUI.Form.AddButton(resourceBundle.GetString("Back"), func() {
		previousPage()
	})
}

func (optionUI *OptionUI) addNextButton() {
	optionUI.Form.AddButton(resourceBundle.GetString("Next"), func() {
		nextPage()
	})
}

func (optionUI *OptionUI) addFinishButton() {
	optionUI.Form.AddButton(resourceBundle.GetString("Finish"), func() {
		finish()
	})
}
