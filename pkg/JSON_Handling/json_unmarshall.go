// Copyright (c) avijit bhattacharjee 2024

package JSON_Handling

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Name struct {
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`

}
type Person struct {
	Name Name `json: "name"`
	Gender string `json: "gender"`
	Company string `json: "company"`
	Exp int `json: "exp"`
}

func Unmarshall() {
	var person Person

	// Open the JSON file
	file, err := os.Open("test-data/person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file contents
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("Error marshalling the json ", err)
	} else {
		fmt.Println("this is the person details", person)
	}

}