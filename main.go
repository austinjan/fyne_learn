package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// FileEntry custom form entry for file selection
type FileEntry struct {
	widget.BaseWidget
	// Entry widget
	Entry *widget.Entry
	// Button widget
	Button *widget.Button
	// File path
	Path string
}

// NewFileEntry creates a new FileEntry
func NewFileEntry(w fyne.Window) *fyne.Container {

	entry := widget.NewEntry()
	button := widget.NewButton("...", func() {
		fileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err == nil && reader != nil {
					entry.SetText(reader.URI().Path())
				}
			}, w)

		fileDialog.Show()
	})
	content := container.NewBorder(nil, nil, nil, button, entry)

	return content
}

// DownloadWidget create a download widget
func DownloadWidget(nimbls []string, w fyne.Window) *fyne.Container {
	//nimbl files listbox
	selected := ""
	label := widget.NewLabel("Select a Nimbl file to download:")
	nimblList := widget.NewSelect(nimbls, func(s string) {
		fmt.Println("Selected:", s)
		selected = s
	})
	// download button
	downloadButton := widget.NewButton("Download", func() {
		if selected != "" {
			dialog.ShowInformation("Downloading", "Downloading file "+selected, w)
			// download file
			err := downloadFile(selected, "./nimbl.zip")
			if err != nil {
				dialog.ShowError(err, w)
			} else {
				dialog.ShowInformation("Downloaded", "Downloaded file to ./nimbl.zip", w)
			}
		}
	})
	return container.NewHBox(label, nimblList, downloadButton)
}

//	func NimblLocationWidget(w fyne.Window) *fyne.Container {
//		button := widget.NewButton("...", func() {
//			folderDialog := dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
//				if err == nil && uri != nil {
//					fmt.Println("Selected folder:", uri.Path())
//				}
//			}
//		}
//	}

// createStartForm StartForm, have licenseFile, privatekey, remote syslog, inputs
func createStartForm(w fyne.Window) *widget.Form {
	licenseFile := NewFileEntry(w)
	privatekey := NewFileEntry(w)
	remoteSyslog := widget.NewEntry()
	inputs := []*widget.FormItem{
		{Text: "License File", Widget: licenseFile},
		{Text: "Private Key", Widget: privatekey},
		{Text: "Remote Syslog", Widget: remoteSyslog},
	}
	return &widget.Form{Items: inputs}
}

func main() {
	a := app.New()
	w := a.NewWindow("Nimbl")

	// load files from blackbeartechhive.com
	nimbls, err := GetNimblFileList()
	if err != nil {
		fyne.LogError("Failed to get nimbl files", err)
	}
	fmt.Println("Nimbl files:", nimbls)

	// logo
	logoFile, err := fyne.LoadResourceFromPath("bblogo.svg")
	if err != nil {
		fyne.LogError("Failed to load logo", err)
		return
	}

	svgImage := canvas.NewImageFromResource(logoFile)
	svgImage.FillMode = canvas.ImageFillOriginal
	a.SetIcon(logoFile)

	downloadWidget := DownloadWidget(nimbls, w)
	startForm := createStartForm(w)
	hello := widget.NewLabel("Hello Nimbl!")
	w.SetContent(container.NewVBox(
		svgImage,
		hello,
		downloadWidget,
		startForm,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.Resize(fyne.NewSize(600, 400))

	w.ShowAndRun()
	fmt.Println("Nimbl exited.")
}
