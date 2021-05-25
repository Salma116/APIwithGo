package main

//PARSING A JSON FILE + CREATING A RESTFUL API
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var records AllSpots

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
	MagicSeaweedLink string `json:"MagicSeaweedLink"`
}

func main() {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &records)
	handleRequests()
	defer jsonFile.Close()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json in homePage")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	for key, value := range result {
		fmt.Fprintf(w, "%s=\"%s\"\n", key, value)
	}

	for i := 0; i < len(records.AllSpots); i++ {
		//fmt.Println("Spot ID: " + records.AllSpots[i].Id)
		fmt.Fprintf(w, "Spot Destination: "+records.AllSpots[i].Fields.Destination)
		//fmt.Println("Spot DifficultyLevel: " + strconv.Itoa(records.AllSpots[i].Fields.DifficultyLevel))
		fmt.Println("Spot fields: " + records.AllSpots[i].Fields.SurfBreak)

	}
}
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
