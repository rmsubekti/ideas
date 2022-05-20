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

// Load Ideas From File
func (ideas *Ideas) Load() error {
	dir, _ := os.UserHomeDir()
	file, err := ioutil.ReadFile(dir + "/.ideas.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(file), &ideas)
	return err
}

// Save ideas to file
func (ideas Ideas) Save() error {
	//data to json
	file, err := json.Marshal(ideas)
	if err != nil {
		return err
	}

	dir, _ := os.UserHomeDir()
	//writing to file
	err = ioutil.WriteFile(dir+"/.ideas.json", file, 0644)
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
