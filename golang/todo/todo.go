package main

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	data := readFile()

	myApp := app.New()
	myWindow := myApp.NewWindow("Todo List")

	list := widget.NewLabel("List")
	value := widget.NewLabel(string(data[:]))
	grid := container.New(layout.NewFormLayout(), list, value)

	myWindow.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Sync...", func() {}))))

	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(500, 320))
	myWindow.ShowAndRun()

	tidyUp()
}

func readFile() []byte {
	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := make([]byte, 1024)
	file.Read(data)
	return data
}

func appendToList(s string) {
	file, err := os.OpenFile("list.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err = file.WriteString("\n" + s); err != nil {
		panic(err)
	}
}

func tidyUp() {
	fmt.Println("Exited")
}
