package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Person struct{
	Name,City string
}


var person Person
var people []Person
func main(){

	// JSON string to parse, see below for example read in from file
	json_str := "{ \"Name\": \"Marcus\", \"City\": \"San Jose\"}"

	// JSON Unmarshal command takes []byte, so string needs to be cast
	// The second parameter is a pointer to the struct that matches format
	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	// output result
	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)
	// read json in from a file
	file, err := ioutil.ReadFile("./extras/names.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	// the names.json file has an array of person objects, so read into people
	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	// output result
	fmt.Println(people)

	// encoding a Go object into JSON is simply using the Marshal command
	json, err := json.Marshal(people)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println(string(json))
}
