package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func (note Note) Save() (err error) {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	encodedNote, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, encodedNote, 0644)
}
