package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

