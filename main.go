package main

import (
	"fmt"
	"github.com/acround/gl/roots2"
	"github.com/acround/gl/record"
)

func main() {
	var iAction int64
	actions := []string{
		"Square roots",
		"Ages",
	}
	fmt.Println("Please choose an action:")
	for index, action := range actions {
		fmt.Printf("\t%d. %s", index+1, action)
		fmt.Println()
	}
	fmt.Scan(&iAction)
	switch iAction {
	case 1:
		roots2.run()
	case 2:
		record.run()
	}
}
