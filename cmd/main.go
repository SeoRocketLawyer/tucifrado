package main

import (
	"fmt"
	"tucifrado/internal/ui"
	"tucifrado/internal/version"
)

func main() {
	fmt.Printf("Tu Cifrado Versión: %s\n", version.Version)
	ui.StartApp()
}
