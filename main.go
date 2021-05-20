package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AllSpots struct {
	AllSpots []Spot `json:"records"`
}
type Spot struct {
	Id     string `json:"id"`
	Fields Fields `json:"Fields"`
}
type Fields struct {
	SurfBreak        string `json:"SurfBreak"`
	DifficultyLevel  int    `json:"DifficultyLevel"`
	Destination      string `json:"Destination"`
	Geocode          string `json:"Geocode"`
	MagicSeaweedLink string `json:"MagicSeaweedLink"`
}

func main() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["records"])

	var records AllSpots

	json.Unmarshal(byteValue, &records)

	for i := 0; i < len(records.AllSpots); i++ {
		fmt.Println("Spot ID: " + records.AllSpots[i].Id)
		fmt.Println("Spot fields: " + records.AllSpots[i].Fields.SurfBreak)
	}
}
