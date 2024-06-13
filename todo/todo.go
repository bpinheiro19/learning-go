package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	data := readFile()
	myApp := app.New()
	myWindow := myApp.NewWindow("Todo List")

	clock := widget.NewLabel("")
	updateTime(clock)

	myWindow.SetContent(clock)
	go func() {
		for range time.Tick(time.Minute) {
			updateTime(clock)
		}
	}()

	entry := widget.NewEntry()

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	add := widget.NewButton("Delete", func() {
		removeFromList("123", data)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "", Widget: entry},
		},
		OnSubmit: func() { // optional, handle form submission
			if entry.Text != "" {
				log.Println("Form submitted:", entry.Text)
				appendToList(entry.Text)
				data = readFile()
				list.Refresh()
			}
		},
	}

	myWindow.SetContent(container.NewBorder(nil, form, nil, add, list))

	myWindow.Resize(fyne.NewSize(500, 320))
	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("2006-01-02 15:04")
	clock.SetText(formatted)
}

func readFile() []string {
	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	strArray := strings.Split(string(data), "\n")

	return strArray
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

func removeFromList(s string, data []string) bool {
	//TODO
	//go through each line
	//find if s is in file
	//remove it if it is
	//return bool from result of the operation

	newData := data
	for i, e := range data {
		if e == s {
			newData = append(data[:i], data[i+1:]...)

			file, err := os.Create("list.txt")
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()
			for _, d := range newData {
				file.Write([]byte(d))
				file.WriteString("\n")
				fmt.Println(d)
			}

			return true
		}
	}

	return false
}

func tidyUp() {
	fmt.Println("Exited")
}
