package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// Idea Struct
type Idea struct {
	Id                int
	Date, State, Idea string
}

// Ideas List
type Ideas []Idea

var filename = ".ideas.js"

// Create Initial File to store data
func Init() error {
	if !fileExist(filename) {
		file, err := os.Create(filename)
		file.Close()
		return err
	}
	return nil
}

// Load Ideas From File
func (ideas *Ideas) Load() error {
	var file []byte
	if fileExist(filename) {
		file, _ = ioutil.ReadFile(filename)
	} else {
		dir, _ := os.UserHomeDir()
		file, _ = ioutil.ReadFile(dir + "/" + filename)
	}
	err := json.Unmarshal([]byte(file), &ideas)
	return err
}

// Save ideas to file
func (ideas Ideas) Save() error {
	//data to json
	file, err := json.Marshal(ideas)
	if err != nil {
		return err
	}

	if fileExist(filename) {
		err = ioutil.WriteFile(filename, file, 0644)
		return err
	}

	dir, _ := os.UserHomeDir()
	err = ioutil.WriteFile(dir+"/"+filename, file, 0644)
	return err
}

// Create new ideas
func (ideas Ideas) Create(note string) error {
	id := 1
	if len(ideas) > 0 {
		id = ideas[len(ideas)-1].Id + 1
	}
	//get current date
	date := time.Now()

	idea := Idea{
		Id:    id,
		Date:  date.Format("Monday, 02 Jan 2006"),
		State: "OPEN",
		Idea:  note,
	}
	ideas = append(ideas, idea)

	return ideas.Save()
}

// Remove Ideas
func (ideas Ideas) Remove(id int) error {
	for i, v := range ideas {
		if v.Id == id {
			ideas = append(ideas[:i], ideas[i+1:]...)
		}
	}
	return ideas.Save()
}

// List all ideas
func (ideas Ideas) List() {
	// Format table colors
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	//create table
	tbl := table.New("ID", "Date", "State", "Idea")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, v := range ideas {
		// add table row
		tbl.AddRow(v.Id, v.Date, v.State, v.Idea)
	}

	tbl.Print()
}

// Solve idea
func (ideas Ideas) Solve(Id int) error {
	i := 0
	for k, v := range ideas {
		if Id == v.Id {
			i = k
		}
	}
	ideas[i].State = "SOLVED"
	return ideas.Save()
}

func (ideas Ideas) ListByState(state string) {
	// Format table colors
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	//create table
	tbl := table.New("ID", "Date", "State", "Idea")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, v := range ideas {
		if v.State == strings.ToUpper(state) {
			tbl.AddRow(v.Id, v.Date, v.State, v.Idea)
		}
	}
	tbl.Print()
}

func fileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
