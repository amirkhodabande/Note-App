package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go.ir/note"
)

func main() {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(note.Title, note.Content, note.CreatedAt)

	_ = note.Save()

	fmt.Println("Note saved successfully!")
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	if value == "" {
		getUserInput(prompt)
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}
