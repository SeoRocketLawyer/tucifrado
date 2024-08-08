package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"tucifrado/internal/cifrado"
)

func StartApp() {
	a := app.New()
	w := a.NewWindow("Tu cifrado Encripta y Desencripta")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Introducir Password")

	fileLabel := widget.NewLabel("No hay archivo seleccionado")
	var selectedFile string

	selectFileButton := widget.NewButton("Elegir archivo", func() {
		dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil || reader == nil {
					return
				}
				selectedFile = reader.URI().Path()
				fileLabel.SetText(selectedFile)
			}, w).Show()
	})

	encryptButton := widget.NewButton("Encriptar", func() {
		if selectedFile == "" {
			dialog.ShowInformation("Error", "Es necesario elegir un archivo", w)
			return
		}
		if passwordEntry.Text == "" {
			dialog.ShowInformation("Error", "Es necesario introducir una contraseña", w)
			return
		}
		err := cifrado.EncryptFile(passwordEntry.Text, selectedFile)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Éxito", "Archivo encriptado con éxito", w)
	})

	decryptButton := widget.NewButton("Desencriptar", func() {
		if selectedFile == "" {
			dialog.ShowInformation("Error", "No ha seleccionado archivo", w)
			return
		}
		if passwordEntry.Text == "" {
			dialog.ShowInformation("Error", "No ha introducido contraseña", w)
			return
		}
		err := cifrado.DecryptFile(passwordEntry.Text, selectedFile)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Éxito", "Archivo desencriptado con éxito", w)
	})

	content := container.NewVBox(
		passwordEntry,
		selectFileButton,
		fileLabel,
		encryptButton,
		decryptButton,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
