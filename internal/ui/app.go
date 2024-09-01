package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"tucifrado/internal/cifrado"
	"tucifrado/internal/version"
)

func StartApp() {
	a := app.NewWithID("com.seorocketlawyer.tucifrado")

	windowTitle := fmt.Sprintf("Tu Cifrando Encripta y Desencripta - Version: %s", version.Version)
	w := a.NewWindow(windowTitle)

	w.Resize(fyne.NewSize(800, 600))

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Introducir contraseña")

	confirmPasswordEntry := widget.NewPasswordEntry()
	confirmPasswordEntry.SetPlaceHolder("Confirma la contraseña")

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
		if passwordEntry.Text == "" || confirmPasswordEntry.Text == "" {
			dialog.ShowInformation("Error", "Es necesario introducir la contraseña", w)
			return
		}

		if passwordEntry.Text != confirmPasswordEntry.Text {
			dialog.ShowInformation("Error", "Las contraseñas no coinciden", w)
			return
		}
		err := cifrado.EncryptFile(passwordEntry.Text, selectedFile)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Éxito", "Archivo encriptado con éxito", w)

		passwordEntry.SetText("")
		confirmPasswordEntry.SetText("")
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

		passwordEntry.SetText("")
		confirmPasswordEntry.SetText("")
	})

	content := container.NewVBox(
		passwordEntry,
		confirmPasswordEntry,
		selectFileButton,
		fileLabel,
		encryptButton,
		decryptButton,
	)

	showLicensePopup(w, content)

	w.SetContent(content)
	w.ShowAndRun()
}

func showLicensePopup(w fyne.Window, content fyne.CanvasObject) {

	licenseText := `Al usar esta aplicación, aceptas que se entrega "tal cual" y su uso es bajo tu propia responsabilidad.`
	licenseLabel := widget.NewLabel(licenseText)
	acceptButton := widget.NewButton("Aceptar", func() {
		w.SetContent(content)
		w.Show()
	})
	licenseDialog := dialog.NewCustom("Términos de la Licencia", "", container.NewVBox(licenseLabel, acceptButton), w)

	licenseDialog.SetDismissText("")
	licenseDialog.Show()
}
