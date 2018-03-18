package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

type SubAhp struct {
	Name        string
	Preferences [][]float64
	Children    []SubAhp
}

type Ahp struct {
	Alternatives []string
	Goal         SubAhp
}

func main() {
	var ahp Ahp = Ahp{
		Alternatives: []string{},
		Goal: SubAhp{
			Name:        "",
			Preferences: [][]float64{},
			Children:    []SubAhp{},
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
			resultString := toJSONString(&ahp)
			fmt.Print(resultString)
			writeToFile(resultString)
		case 3:
			fmt.Print("exiting program")
			programRunning = false
		default:
			fmt.Print("try again")
		}
	}
}

func toJSONString(ahp *Ahp) string {
	json, err := json.Marshal(*ahp)
	if err != nil {
		fmt.Print(err)
	}
	result := string(json)
	return result
}

func writeToFile(stringToWrite string) {
	f, err := os.OpenFile("ahp.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Print(err)
		return
	}
	bytes, err := f.WriteString(stringToWrite)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("wrote %d bytes\n", bytes)
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
			ahp.Alternatives = append(ahp.Alternatives, name)
		} else {
			addAnother = false
		}
	}

}

func showData(ahp *Ahp) {
	fmt.Print(*ahp)
}
