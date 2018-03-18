package main

import (
	"fmt"

	"bufio"
	"os"
)

type SubAhp struct {
	name        string
	preferences [][]float64
	children    []SubAhp
}

type Ahp struct {
	alternatives []string
	goal         SubAhp
}

func main() {
	var ahp Ahp = Ahp{
		alternatives: []string{},
		goal: SubAhp{
			name:        "",
			preferences: [][]float64{},
			children:    []SubAhp{},
		},
	}

	var choice int
	options := []string{"Enter Data", "Print Data", "Export as JSON", "Exit Program"}
	programRunning := true

	for programRunning {
		printOptions(options)
		choice = getChoice()
		switch choice {
		case 0:
			enterData(&ahp)
		case 1:
			showData(&ahp)
		case 2:
			saveAsJSON()
		case 3:
			fmt.Print("exiting program")
			programRunning = false
		default:
			fmt.Print("try again")
		}
	}
}

func saveAsJSON() {
	
}

func printOptions(options []string) {
	printLine()
	fmt.Print("Choose one:\n")
	for i, el := range options {
		fmt.Printf("%d. %s\n", i, el)
	}
	printLine()
}

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
	options := []string{"Back to Menu", "Enter alternatives"}
	printOptions(options)
	choice := getChoice()

	switch choice {
	case 0:
		return
	case 1:
		enterAlternative(ahp)
	}
}

func printLine() {
	fmt.Print("################################\n")
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
			ahp.alternatives = append(ahp.alternatives, name)
		} else {
			addAnother = false
		}
	}

}

func showData(ahp *Ahp) {
	fmt.Print(*ahp)
}
