package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func writeToFile(ahp *Ahp) {
	stringToWrite := toJSONString(ahp)
	f, err := os.OpenFile("ahp.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

func toJSONString(ahp *Ahp) string {
	jsonString, err := json.Marshal(*ahp)
	if err != nil {
		fmt.Print(err)
	}
	result := string(jsonString)
	return result
}

func loadFromJsonFIle(filePath string, ahp *Ahp) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	ahpFromJson := Ahp{}
	json.Unmarshal(buf, &ahpFromJson)
	*ahp = ahpFromJson
}
