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
	options := []string{"Enter Data", "Print Data", "Export as JSON", "Exit Program"}
	programRunning := true

	for programRunning {
		printOptions(options)
		choice = getChoice()
		clearConsole()
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
