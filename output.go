package main

import (
	"fmt"
	"os/exec"
	"os"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printOptions(options []string) {
	printLine()
	fmt.Print("Choose one:\n")
	for i, el := range options {
		fmt.Printf("%d. %s\n", i, el)
	}
	printLine()
}

func printLine() {
	fmt.Print("===================================================================================\n")
}

func showData(ahp *Ahp) {
	fmt.Print(*ahp)
	fmt.Print("\n")
}
