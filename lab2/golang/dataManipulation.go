package main

import (
	"fmt"
	"bufio"
	"os"
)

func enterData(ahp *Ahp) {
	options := []string{"Edit alternatives", "Edit Goals", "Back to Menu"}
	printOptions(options)
	choice := getChoice()

	switch choice {
	case 0:
		enterAlternative(ahp)
	case 1:
		editGoals(ahp)
	case 2:
		return
	}
}

func editGoals(ahp *Ahp) {
	fmt.Print("edit goals")
}

func enterAlternative(ahp *Ahp) {
	options := []string{"Add alternative", "Enough"}
	addAnother := true
	for addAnother == true {
		printOptions(options)
		choice := getChoice()
		if choice == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter alternative name: ")
			name, _ := reader.ReadString('\n')
			ahp.Alternatives = append(ahp.Alternatives, name)
		} else {
			addAnother = false
		}
	}
}


