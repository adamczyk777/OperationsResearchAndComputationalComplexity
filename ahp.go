package main

import (
	"fmt"
)

func main() {
	var ahp = Ahp{
		Alternatives: []string{},
		Goal: SubAhp{
			Name:        "",
			Preferences: [][]float64{},
			Children:    []SubAhp{},
		},
	}

	var choice int
	options := []string{"Exit Program", "Enter Data", "Print Data", "Export as JSON", "Load from JSON file", "Calculate Using Eigenvalue Method", "Calcualate using method"}
	programRunning := true

	for programRunning {
		printOptions(options)
		choice = readChoice()
		clearConsole()
		switch choice {
		case 0:
			fmt.Print("exiting program")
			programRunning = false
		case 1:
			getEditDataMode(&ahp)
		case 2:
			showData(&ahp)
		case 3:
			writeToFile(&ahp)
			fmt.Print("Written to a file")
		case 4:
			file := readString()
			loadFromJsonFIle(file, &ahp)
		case 5:
			calculateWithEigenvalue()
		case 6:
			calculateUsing()
		default:
			fmt.Print("try again")
		}
	}
}
