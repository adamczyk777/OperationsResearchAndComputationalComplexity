package main

import "fmt"

func main() {
	options := []string{"Enter Alternatives", "Show currently gathered data", "Exit"}
	var choice string
	printChoices(options)
	fmt.Scanf("Choice: %s", &choice)
	switch choice {
	case "0":
		readData()
		break
	case "1":
		showData()
		fmt.Printf("1")
		break
	case "2":
		fmt.Printf("2")
		break
	default:
		fmt.Printf("there is no %s choice, choose again!", choice)
	}
}

func readData() {

}

func showData() {

}

func printChoices(choices []string) {
	fmt.Printf("Hello in AHP CLI!\n")
	fmt.Printf("Choose one of options below:\n")
	for i, option := range choices {
		fmt.Printf("%d. %s \n", i, option)
	}
}
type Ahp struct {

}
