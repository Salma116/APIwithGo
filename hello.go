package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Spot struct {
	ID              string `json:"ID"`
	SurfBreak       string `json:"SurfBreak"`
	DifficultyLevel string `json:"DifficultyLevel"`
}
type allSpots []Spot

var Spots = allSpots{
	{
		ID:              "rec5aF9TjMjBicXCK",
		SurfBreak:       "Reef Break",
		DifficultyLevel: "4",
	},
}

func createSpots(w http.ResponseWriter, r *http.Request) {
	var newSpots Spot
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Spot title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newSpots)
	Spots = append(Spots, newSpots)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newSpots)
}
func getOneSpot(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleSpot := range Spots {
		if singleSpot.ID == eventID {
			json.NewEncoder(w).Encode(singleSpot)
		}
	}
}
func getAllSpots(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Spots)
}
func updateSpot(w http.ResponseWriter, r *http.Request) {
	SpotID := mux.Vars(r)["id"]
	var updatedSpot Spot

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Spot title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedSpot)

	for i, singleSpot := range Spots {
		if singleSpot.ID == SpotID {
			singleSpot.SurfBreak = updatedSpot.SurfBreak
			singleSpot.DifficultyLevel = updatedSpot.DifficultyLevel
			Spots = append(Spots[:i], singleSpot)
			json.NewEncoder(w).Encode(singleSpot)
		}
	}
}

func deleteSpot(w http.ResponseWriter, r *http.Request) {
	SpotID := mux.Vars(r)["id"]

	for i, singleSpot := range Spots {
		if singleSpot.ID == SpotID {
			Spots = append(Spots[:i], Spots[i+1:]...)
			fmt.Fprintf(w, "The Spot with ID %v has been deleted successfully", SpotID)
		}
	}
}

// func testPostSpots(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Test POST Endpoint worked")
// }
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
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
func main() {
	handleRequests()
}
