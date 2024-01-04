package main

import (
	"fmt"

	lesson "github.com/acround/gl/lessons"
	"github.com/acround/gl/lessons/files"
	"github.com/acround/gl/lessons/record"
	"github.com/acround/gl/lessons/roots2"
	"github.com/acround/gl/lessons/routine"
	"github.com/acround/gl/lessons/routine2"
)

func main() {
	var iAction int64
	actions, actionNames := initVars()
	fmt.Println("Please choose an action:")
	for index, action := range actionNames {
		fmt.Printf("\t%d. %s", index+1, action)
		fmt.Println()
	}
	fmt.Scan(&iAction)
	actions[iAction-1].Run()
}

func initVars() ([]lesson.Lesson, []string) {
	return []lesson.Lesson{
			roots2.New(),
			record.New(),
			files.New(),
			routine.New(),
			routine2.New(),
		},
		[]string{
			"Square roots",
			"Ages",
			"Files",
			"Routines",
			"Routines & channels",
		}
}
