package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AllSpots struct {
	AllSpots []Spot `json:"users"`
}
type Spot struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Social Social `json:"social"`
}
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["users"])

	var users AllSpots

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.AllSpots); i++ {
		fmt.Println("Spot Type: " + users.AllSpots[i].Type)
		fmt.Println("Spot social: " + users.AllSpots[i].Social.Facebook)
	}
}
