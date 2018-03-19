package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getChoice() int {
	var i int
	fmt.Print("Your Choice:")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return -1
	} else {
		return i
	}
}

func enterData(ahp *Ahp) {
	options := []string{"Edit alternatives", "Edit Goals", "Back to Menu"}
	printOptions(options)
	choice := getChoice()

	switch choice {
	case 0:
		enterAlternative(ahp)
	case 1:
		editNode(&ahp.Goal)
	case 2:
		return
	}
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
			ahp.Alternatives = append(ahp.Alternatives, strings.TrimSpace(name))
		} else {
			addAnother = false
		}
	}
}

func editNode(node *SubAhp) {
	options := []string{"Edit Name", "Edit Preferences", "Edit Children", "Select Child", "Back"}
	editing := true
	for editing == true {
		printOptions(options)
		choice := getChoice()
		switch choice {
		case 0:
			editName(node)
		case 1:
			editPreferences(node)
		case 2:
			editChildren(node)
		case 3:
			selectChild(node)
		case 4:
			editing = false
		}
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

func editChildren(node *SubAhp) {
	printOptions([]string{"Add Child", "Remove Child", "Back"})
	choice := getChoice()
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

func selectChild(node *SubAhp) {
	clearConsole()
	fmt.Print("Choose node or go back:\n")
	for i, child := range node.Children {
		fmt.Printf("%d. %s\n", i, child.Name)
	}
	fmt.Printf("%d. Back\n", len(node.Children))
	fmt.Print("Your Choice:\n")
	choice := getChoice()
	if choice == len(node.Children) {
		return
	} else if choice < len(node.Children) && choice >= 0 {
		editNode(&node.Children[choice])
	}
}

func showData(ahp *Ahp) {
	fmt.Print(*ahp)
	fmt.Print("\n")
}
