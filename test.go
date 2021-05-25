package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// type AllSpots struct {
// 	AllSpots []Spot `json:"records"`
// }
// type Spot struct {
// 	Id     string `json:"id"`
// 	Fields Fields `json:"fields"`
// }
// type Fields struct {
// 	SurfBreak        string `json:"Surf Break"`
// 	DifficultyLevel  int    `json:"Difficulty Level"`
// 	Destination      string `json:"Destination"`
// 	Geocode          string `json:"Geocode"`
// 	MagicSeaweedLink string `json:"MagicSeaweedLink"`
// }
type AllSpots struct {
	Records []struct {
		ID     string `json:"id"`
		Fields struct {
			SurfBreak        []string `json:"Surf Break"`
			DifficultyLevel  int      `json:"Difficulty Level"`
			Destination      string   `json:"Destination"`
			Geocode          string   `json:"Geocode"`
			Influencers      []string `json:"Influencers"`
			MagicSeaweedLink string   `json:"Magic Seaweed Link"`
			Photos           []struct {
				ID         string `json:"id"`
				URL        string `json:"url"`
				Filename   string `json:"filename"`
				Size       int    `json:"size"`
				Type       string `json:"type"`
				Thumbnails struct {
					Small struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"small"`
					Large struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"large"`
					Full struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"full"`
				} `json:"thumbnails"`
			} `json:"Photos"`
			PeakSurfSeasonBegins    string `json:"Peak Surf Season Begins"`
			DestinationStateCountry string `json:"Destination State/Country"`
			PeakSurfSeasonEnds      string `json:"Peak Surf Season Ends"`
			Address                 string `json:"Address"`
		} `json:"fields"`
		Createdtime time.Time `json:"createdTime"`
	} `json:"records"`
	Offset string `json:"offset"`
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

	for i := 0; i < len(records.AllSpots.Records); i++ {
		fmt.Println(records.AllSpots[i].Records[i].ID)
		//fmt.Println("Spot ID: " + records.AllSpots[i].Id)
		//fmt.Println("Spot Destination: " + records.AllSpots[i].Fields.Destination)
		//fmt.Println("Spot DifficultyLevel: " + strconv.Itoa(records.AllSpots[i].Fields.DifficultyLevel))
		//fmt.Println("Spot fields: " + records.AllSpots[i].Fields.SurfBreak)

	}

	defer jsonFile.Close()
}
