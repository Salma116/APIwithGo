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
func createSpots(w http.ResponseWriter, r *http.Request) {
	var newSpots Spot
	jsonFile, err := os.Open("data.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Spot title and description only in order to update")
	}

	json.Unmarshal(byteValue, &newSpots)
	var result map[string]interface{}
	//json.Unmarshal(byteValue, &records)
	for key, value := range result {
		fmt.Fprintf(w, "%s=\"%s\"\n", key, value)

	}
	records.AllSpots = append(records.AllSpots, newSpots)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newSpots)
}

func getOneSpot(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleSpot := range records.AllSpots {
		if singleSpot.Id == eventID {
			json.NewEncoder(w).Encode(singleSpot)
		}
	}
}
func getAllSpots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(records.AllSpots)
}

func updateSpot(w http.ResponseWriter, r *http.Request) {
	SpotID := mux.Vars(r)["id"]
	var updatedSpot Spot

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Spot title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedSpot)

	for i, singleSpot := range records.AllSpots {
		if singleSpot.Id == SpotID {
			singleSpot.Fields.SurfBreak = updatedSpot.Fields.SurfBreak
			singleSpot.Fields.DifficultyLevel = updatedSpot.Fields.DifficultyLevel
			records.AllSpots = append(records.AllSpots[:i], singleSpot)
			json.NewEncoder(w).Encode(singleSpot)
		}
	}
}
func deleteSpot(w http.ResponseWriter, r *http.Request) {
	SpotID := mux.Vars(r)["id"]

	for i, singleSpot := range records.AllSpots {
		if singleSpot.Id == SpotID {
			records.AllSpots = append(records.AllSpots[:i], records.AllSpots[i+1:]...)
			fmt.Fprintf(w, "The Spot with ID %v has been deleted successfully", SpotID)
		}
	}
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
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/spots", createSpots).Methods("POST")
	myRouter.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	myRouter.HandleFunc("/spots", getAllSpots).Methods("GET")
	myRouter.HandleFunc("/spots/{id}", updateSpot).Methods("PATCH")
	myRouter.HandleFunc("/spots/{id}", deleteSpot).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
