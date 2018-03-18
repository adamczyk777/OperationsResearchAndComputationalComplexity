package main

import "fmt"

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

func printLine() {
	fmt.Print("################################\n")
}

func showData(ahp *Ahp) {
	fmt.Print(*ahp)
}
