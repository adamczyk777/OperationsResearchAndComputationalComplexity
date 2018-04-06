package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

// Node edit mode selection
func getNodeEditMode(node *SubAhp) {
	options := []string{"Edit Name", "Edit Preferences", "Edit Children", "Select Child", "Back"}
	editing := true
	for editing == true {
		printOptions(options)
		choice := readChoice()
		switch choice {
		case 0:
			//exit editing
			editing = false
		case 1:
			// edit name
			editName(node)
		case 2:
			// edit preferences
			editPreferences(node)
		case 3:
			// edit children list
			getEditChildrenMode(node)
		case 4:
			// select next child to edit
			selectChild(node)
		}
	}
}

// Read string from command line
func readString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	result, _ := reader.ReadString('\n')
	result = strings.TrimSpace(result)
	return result
}

// Choice reader from console input
func readChoice() int {
	var i int
	fmt.Print("Your Choice:")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return -1
	} else {
		return i
	}
}

// edit data mode selection
func getEditDataMode(ahp *Ahp) {
	options := []string{"Edit alternatives", "Edit Goals", "Back to Menu"}
	printOptions(options)
	choice := readChoice()

	switch choice {
	case 1:
		getAlternativeMode(ahp)
	case 2:
		getNodeEditMode(&ahp.Goal)
	case 0:
		return
	}
}

// alternative manipulation mode selection
func getAlternativeMode(ahp *Ahp) {
	options := []string{"Add alternative", "Enough"}
	addAnother := true
	for addAnother == true {
		printOptions(options)
		choice := readChoice()
		if choice == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter alternative name: ")
			name, _ := reader.ReadString('\n')
			ahp.Alternatives = append(ahp.Alternatives, strings.TrimSpace(name))
		} else {
			addAnother = false
		}
	}
}

//
func selectChild(node *SubAhp) {
	clearConsole()
	fmt.Print("Choose node or go back:\n")
	for i, child := range node.Children {
		fmt.Printf("%d. %s\n", i, child.Name)
	}
	fmt.Printf("%d. Back\n", len(node.Children))
	fmt.Print("Your Choice:\n")
	choice := readChoice()
	if choice == len(node.Children) {
		return
	} else if choice < len(node.Children) && choice >= 0 {
		getNodeEditMode(&node.Children[choice])
	}
}

func editName(node *SubAhp) {
	fmt.Printf("Current name: %s \n", node.Name)
	fmt.Print("New name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	node.Name = strings.TrimSpace(name)
}

func editPreferences(node *SubAhp) {
	fmt.Print("Provide me with a MATRIX, syntax: 1,2,3;1,2,3;1,2,3 where selicolon notifies that next row is entered:\n")
	reader := bufio.NewReader(os.Stdin)
	matrix, _ := reader.ReadString('\n')
	matrix = strings.TrimSpace(matrix)
	rows := strings.Split(matrix, ";")
	var resultMatrix [][]float64
	for _, row := range rows {
		var floatRow []float64
		for _, element := range strings.Split(row, ",") {
			floatElement, _ := strconv.ParseFloat(element, 64)
			floatRow = append(floatRow, floatElement)
		}
		resultMatrix = append(resultMatrix, floatRow)
	}
	node.Preferences = resultMatrix
}

func getEditChildrenMode(node *SubAhp) {
	printOptions([]string{"Add Child", "Remove Child", "Back"})
	choice := readChoice()
	switch choice {
	case 0:
		clearConsole()
		fmt.Print("name:\n")
		name := readString()
		node.Children = append(node.Children, SubAhp{strings.TrimSpace(name), [][]float64{}, []SubAhp{}})
	case 1:
		fmt.Print("not available")
	case 2:
		return
	}
}
