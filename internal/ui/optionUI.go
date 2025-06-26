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
	RootPanel        *tview.Flex
	FormPanel        *tview.Flex
	Form             *tview.Form
	DescriptionPanel *tview.TextView
	ButtonView       *tview.Form
	Fields           map[*option.Option]tview.Primitive
}

func NewOptionsUI(options []*option.Option) *OptionUI {
	optionUI := &OptionUI{
		RootPanel:        tview.NewFlex(),
		FormPanel:        tview.NewFlex(),
		Form:             tview.NewForm(),
		DescriptionPanel: tview.NewTextView(),
		ButtonView:       tview.NewForm(),
		Fields:           make(map[*option.Option]tview.Primitive),
	}

	for _, opt := range options {
		var field tview.FormItem

		if len(opt.GetPossibleOptions()) > 0 {
			field = tview.NewDropDown().
				SetLabel(opt.GetOptionName()).
				SetOptions(opt.GetPossibleOptions(), func(text string, index int) {})

			field.(*tview.DropDown).SetFocusFunc(func() {
				optionUI.DescriptionPanel.SetText(opt.GetOptionDesc())
			})
		} else if opt.HasNoParameter() {
			field = tview.NewCheckbox().
				SetLabel(opt.GetOptionName())

			field.(*tview.Checkbox).SetFocusFunc(func() {
				optionUI.DescriptionPanel.SetText(opt.GetOptionDesc())
			})
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

			field.(*tview.DropDown).SetFocusFunc(func() {
				optionUI.DescriptionPanel.SetText(opt.GetOptionDesc())
			})
		} else {
			field = tview.NewInputField().
				SetLabel(opt.GetOptionName()).
				SetText("")
			field.(*tview.InputField).SetFocusFunc(func() {
				optionUI.DescriptionPanel.SetText(opt.GetOptionDesc())
			})
		}

		if opt.IsOptional() && !opt.HasNoParameter() {
			checkbox := tview.NewCheckbox().
				SetLabel(fmt.Sprintf(resourceBundle.GetString("Include"), opt.GetOptionName()))

			checkbox.SetFocusFunc(func() {
				optionUI.DescriptionPanel.SetText(opt.GetOptionDesc())
			})

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

	// Root settings
	optionUI.RootPanel.AddItem(optionUI.FormPanel, 0, 1, true)
	optionUI.RootPanel.AddItem(optionUI.ButtonView, 1, 0, false)
	optionUI.RootPanel.
		SetDirection(tview.FlexRow).
		SetBackgroundColor(Colors.BackgroundColor).
		SetBorder(true).
		SetBorderColor(Colors.BorderColor).
		SetTitle("jpackageTUI")

	optionUI.RootPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			focusedFormItem, _ := optionUI.Form.GetFocusedItemIndex()
			_, focusedButton := optionUI.ButtonView.GetFocusedItemIndex()

			if optionUI.Form.HasFocus() && optionUI.Form.GetFormItemCount()-1 == focusedFormItem {
				optionUI.ButtonView.SetFocus(0)
				UI.app.SetFocus(optionUI.ButtonView)
				optionUI.DescriptionPanel.SetText("")

				return nil
			}

			if optionUI.ButtonView.HasFocus() && optionUI.ButtonView.GetButtonCount()-1 == focusedButton {
				optionUI.Form.SetFocus(0)
				UI.app.SetFocus(optionUI.Form)

				return nil
			}

			if focusedFormItem != -1 {
				if checkbox, ok := optionUI.Form.GetFormItem(focusedFormItem).(*tview.Checkbox); ok {
					if !checkbox.IsChecked() && optionUI.Form.GetFormItemCount()-2 == focusedFormItem {
						optionUI.ButtonView.SetFocus(0)
						UI.app.SetFocus(optionUI.ButtonView)
						optionUI.DescriptionPanel.SetText("")

						return nil
					}
				}
			}
		}

		return event
	})

	// FormPanel settings
	optionUI.FormPanel.AddItem(optionUI.Form, 0, 1, true)
	optionUI.FormPanel.AddItem(optionUI.DescriptionPanel, 0, 1, false)

	// Main Form settings
	optionUI.Form.
		SetFieldBackgroundColor(Colors.FieldBackgroundColor).
		SetBackgroundColor(Colors.BackgroundColor)

	optionUI.Form.
		SetLabelColor(Colors.LabelColor)

	// DescriptionPanel settings
	optionUI.DescriptionPanel.
		SetTextColor(Colors.LabelColor).
		SetBackgroundColor(Colors.BackgroundColor).
		SetBorderPadding(1, 1, 1, 1)

	// ButtonView Settings
	optionUI.ButtonView.
		SetButtonBackgroundColor(Colors.ButtonBackgroundColor).
		SetBackgroundColor(Colors.BackgroundColor).
		SetBorderPadding(0, 0, 1, 1)

	optionUI.ButtonView.SetButtonsAlign(tview.AlignRight)

	return optionUI
}

func (optionUI *OptionUI) GetPrimitive() tview.Primitive {
	return optionUI.RootPanel
}

func (optionUI *OptionUI) getData() map[*option.Option]string {
	var data = make(map[*option.Option]string)

	for opt, field := range optionUI.Fields {
		switch field.(type) {
		case *tview.InputField:
			inputField := field.(*tview.InputField)
			if opt.IsOptional() {
				if optionUI.Form.GetFormItem(optionUI.Form.GetFormItemIndex(inputField.GetLabel()) - 1).(*tview.Checkbox).IsChecked() {
					data[opt] = inputField.GetText()
				}
			} else {
				data[opt] = inputField.GetText()
			}
		case *tview.DropDown:
			dropdown := field.(*tview.DropDown)
			if opt.IsOptional() {
				if optionUI.Form.GetFormItem(optionUI.Form.GetFormItemIndex(dropdown.GetLabel()) - 1).(*tview.Checkbox).IsChecked() {
					_, data[opt] = dropdown.GetCurrentOption()
				}
			} else {
				_, data[opt] = dropdown.GetCurrentOption()
			}
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
	optionUI.ButtonView.AddButton(resourceBundle.GetString("Back"), func() {
		previousPage()
		optionUI.resetFocus()
	})
}

func (optionUI *OptionUI) addNextButton() {
	optionUI.ButtonView.AddButton(resourceBundle.GetString("Next"), func() {
		nextPage()
		optionUI.resetFocus()
	})
}

func (optionUI *OptionUI) addFinishButton() {
	optionUI.ButtonView.AddButton(resourceBundle.GetString("Finish"), func() {
		finish()
	})
}

func (optionUI *OptionUI) resetFocus() {
	optionUI.ButtonView.SetFocus(0)
	optionUI.Form.SetFocus(0)
}
