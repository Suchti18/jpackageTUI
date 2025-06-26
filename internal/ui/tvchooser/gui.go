package tvchooser

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nils/jpackageTUI/internal/Const/Colors"
	"github.com/nils/jpackageTUI/internal/Const/resourceBundle"
	"github.com/rivo/tview"
	"os"
)

// FileChooser let the user selects a file and returns the selected file path.
//
// Takes in a parent application to be paused or nil, and a boolean showHidden to determine if hidden files should be shown.
// Returns a string representing the selected file path.
func FileChooser(parentApp *tview.Application, showHidden bool, fastAccessPaths ...string) string {
	selectedPath := ""

	app := tview.NewApplication()
	runApp := func() {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}

	selectedPathView := tview.NewTextView()
	selectedPathView.SetBorder(true)

	dirView := newDirectoryView(showHidden, selectedPathView, nil, fastAccessPaths)
	fileView := newFileView("", dirView.showHidden, selectedPathView, dirView)
	dirView.onSelectedFunc = func(node *tview.TreeNode) {
		fileView.updatePath(node.GetReference().(nodeInfo).Path)
	}
	selectionPanel := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(dirView.dirView, 0, 1, true).
		AddItem(fileView.fileList, 0, 1, false)

	buttonsView := tview.NewForm()
	buttonsView.SetButtonsAlign(tview.AlignRight)
	// Cancel button
	buttonsView.AddButton(resourceBundle.GetString("Cancel"), func() {
		selectedPath = ""
		app.Stop()
	})
	// Accept button
	buttonsView.AddButton(resourceBundle.GetString("Accept"), func() {
		selectedPath = selectedPathView.GetText(false)
		//if selected path ends with PathSeparator, is a directory, so set selected path to ""
		if len(selectedPath) > 0 && selectedPath[len(selectedPath)-1] == os.PathSeparator {
			selectedPath = ""
		}
		app.Stop()
	})

	buttonsView.SetButtonBackgroundColor(Colors.ButtonBackgroundColor)

	rootPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(selectedPathView, 3, 0, false).
		AddItem(selectionPanel, 0, 1, true).
		AddItem(buttonsView, 3, 0, false)

	selectedPathView.SetBackgroundColor(Colors.BackgroundColor)
	dirView.dirView.SetBackgroundColor(Colors.BackgroundColor)
	fileView.fileList.SetBackgroundColor(Colors.BackgroundColor)
	buttonsView.SetBackgroundColor(Colors.BackgroundColor)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			selectedPath = ""
			app.Stop()
		} else if event.Key() == tcell.KeyRight {
			if dirView.dirView.HasFocus() {
				if fileView.fileList.GetItemCount() != 0 {
					app.SetFocus(fileView.fileList)
				} else {
					app.SetFocus(buttonsView)
				}
			} else if fileView.fileList.HasFocus() {
				app.SetFocus(buttonsView)
			} else {
				app.SetFocus(dirView.dirView)
			}

			return nil
		} else if event.Key() == tcell.KeyLeft {
			if dirView.dirView.HasFocus() {
				app.SetFocus(buttonsView)
			} else if fileView.fileList.HasFocus() {
				app.SetFocus(dirView.dirView)
			} else {
				app.SetFocus(fileView.fileList)
			}

			return nil
		}

		return event
	})

	app.SetRoot(rootPanel, true).EnableMouse(true).EnablePaste(true)
	if parentApp != nil {
		parentApp.Suspend(func() {
			runApp()
		})
	} else {
		runApp()
	}

	return selectedPath
}

// DirectoryChooser selects a directory using a GUI and returns the selected directory path.
//
// It takes in a parent application to be paused or nil, and a boolean showHidden to determine if hidden directories should be shown.
// It returns a string representing the selected directory path.
func DirectoryChooser(parentApp *tview.Application, showHidden bool, fastAccessPaths ...string) string {
	selectedPath := ""

	app := tview.NewApplication()
	runApp := func() {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}

	selectedPathView := tview.NewTextView()
	selectedPathView.SetBorder(true)

	dirView := newDirectoryView(showHidden, selectedPathView, nil, fastAccessPaths)
	selectionPanel := tview.NewFlex().SetDirection(tview.FlexColumn).AddItem(dirView.dirView, 0, 2, true)

	buttonsView := tview.NewForm()
	buttonsView.SetButtonsAlign(tview.AlignRight)

	// Cancel button
	buttonsView.AddButton(resourceBundle.GetString("Cancel"), func() {
		selectedPath = ""
		app.Stop()
	})

	// Accept button
	buttonsView.AddButton(resourceBundle.GetString("Accept"), func() {
		selectedPath = dirView.selectedPath
		app.Stop()
	})

	buttonsView.SetButtonBackgroundColor(Colors.ButtonBackgroundColor)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			selectedPath = ""
			app.Stop()
		} else if event.Key() == tcell.KeyRight || event.Key() == tcell.KeyLeft {
			if dirView.dirView.HasFocus() {
				app.SetFocus(buttonsView)
			} else {
				app.SetFocus(dirView.dirView)
			}

			return nil
		}

		return event
	})

	rootPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(selectedPathView, 3, 0, false).
		AddItem(selectionPanel, 0, 1, true).
		AddItem(buttonsView, 3, 0, false)

	selectedPathView.SetBackgroundColor(Colors.BackgroundColor)
	dirView.dirView.SetBackgroundColor(Colors.BackgroundColor)
	buttonsView.SetBackgroundColor(Colors.BackgroundColor)

	app.SetRoot(rootPanel, true).EnableMouse(true).EnablePaste(true)
	if parentApp != nil {
		parentApp.Suspend(func() {
			runApp()
		})
	} else {
		runApp()
	}

	return selectedPath
}
