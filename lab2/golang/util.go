package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
)

func printOptions(options []string) {
	printLine()
	fmt.Print("Choose one:\n")
	for i, el := range options {
		fmt.Printf("%d. %s\n", i, el)
	}
	printLine()
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printLine() {
	fmt.Print("===================================================================================\n")
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

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value: ")
	result, _ := reader.ReadString('\n')
	result = strings.TrimSpace(result)
	return result
}
