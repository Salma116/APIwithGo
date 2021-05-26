package main

//PARSING A JSON FILE
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type AllSpots struct {
	AllSpots []Spot `json:"records"`
}
type Spot struct {
	Id     string `json:"id"`
	Fields Fields `json:"fields"`
}
type Fields struct {
	SurfBreak        string `json:"Surf Break"`
	DifficultyLevel  int    `json:"Difficulty Level"`
	Destination      string `json:"Destination"`
	Geocode          string `json:"Geocode"`
	MagicSeaweedLink string `json:"Magic Seaweed Link"`
}

func main() {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var records AllSpots

	json.Unmarshal(byteValue, &records)

	for i := 0; i < len(records.AllSpots); i++ {
		fmt.Println("Spot ID: " + records.AllSpots[i].Id)
		fmt.Println("Spot Destination: " + records.AllSpots[i].Fields.Destination)
		fmt.Println("Spot Difficulty Level: " + strconv.Itoa(records.AllSpots[i].Fields.DifficultyLevel))
		fmt.Println("Spot fields: " + records.AllSpots[i].Fields.SurfBreak)
		fmt.Println("Spot Geocode: " + records.AllSpots[i].Fields.Geocode)
		fmt.Println("Spot MagicSeaweedLink: " + records.AllSpots[i].Fields.MagicSeaweedLink)

	}

	defer jsonFile.Close()
}
