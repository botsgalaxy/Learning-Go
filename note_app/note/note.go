package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct { 
	Title string
	Content string
	Created_at time.Time
}

func (note Note) Save() error{ 
	fileName := strings.ReplaceAll(note.Title," ","_")
	fileName = strings.ToLower(fileName) + ".json"

	json,err := json.Marshal(note)
	if err != nil { 
		return err
	}
	return os.WriteFile(fileName,json,0644)
}

func (note Note) Display() { 
	fmt.Printf("\nYour note title: %v \n",note.Title)
	fmt.Printf("Your note content: %v \n",note.Content)
	fmt.Printf("Your note created at: %v \n\n",note.Created_at)

}


func New(title,content string) (Note,error){ 
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{ 
		Title: title,
		Content: content,
		Created_at: time.Now(),
	},nil 
}