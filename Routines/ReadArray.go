package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadArray is a function which reads array of ints from file
func ReadArray(filename string) []int {
	result := make([]int, 5)
	configFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("opening config file" + err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&result); err != nil {
		fmt.Println("parsing config file" + err.Error())
	}
	return result
}
