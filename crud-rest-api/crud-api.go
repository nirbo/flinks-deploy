package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// The Person struct represents a person object
type Person struct {
	FirstName          string     `json:"first_name"`
	LastName           string     `json:"last_name"`
	PhysicalAttributes Attributes `json:"physical_attributes"`
}

// The Attributes struct specifies the physical attributes of a person
type Attributes struct {
	Height   string `json:"height"`
	EyeColor string `json:"eye_color"`
}

var persons []Person = []Person{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/persons", addPerson).Methods("POST")
	router.HandleFunc("/persons", getAllPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", getPersonByID).Methods("GET")
	router.HandleFunc("/persons/{id}", updatePersonByID).Methods("PUT")
	router.HandleFunc("/persons/{id}", deletePersonByID).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

// Delete a single person object by ID
func deletePersonByID(writer http.ResponseWriter, request *http.Request) {
	var idValue string = mux.Vars(request)["id"]
	id, err := strconv.Atoi(idValue)
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("Failed to convert ID to an integer"))
		return
	}

	if id >= len(persons) {
		writer.WriteHeader(404)
		writer.Write([]byte("Could not find a person with the specified ID"))
		return
	}

	persons = append(persons[:id], persons[id+1:]...)

	writer.WriteHeader(200)
}

// Update field values of a single person object by ID
func updatePersonByID(writer http.ResponseWriter, request *http.Request) {
	var idValue string = mux.Vars(request)["id"]
	id, err := strconv.Atoi(idValue)
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("Failed to convert ID to an integer"))
		return
	}

	if id >= len(persons) {
		writer.WriteHeader(404)
		writer.Write([]byte("Could not find a person with the specified ID"))
		return
	}

	var updatedPerson Person
	json.NewDecoder(request.Body).Decode(&updatedPerson)
	persons[id] = updatedPerson

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedPerson)
}

// Get a single person object by ID
func getPersonByID(writer http.ResponseWriter, request *http.Request) {
	var idValue string = mux.Vars(request)["id"]
	id, err := strconv.Atoi(idValue)
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("Failed to convert ID to an integer"))
		return
	}

	if id >= len(persons) {
		writer.WriteHeader(404)
		writer.Write([]byte("Could not find a person with the specified ID"))
		return
	}

	person := persons[id]

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(person)
}

// Get the full array with all person objects in it
func getAllPersons(writer http.ResponseWriter, request *http.Request) {
	if len(persons) <= 0 {
		writer.WriteHeader(404)
		writer.Write([]byte("No person objects found, the array is completely empty"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(persons)
}

// Add a person object to the array
func addPerson(writer http.ResponseWriter, request *http.Request) {
	var newPerson Person
	json.NewDecoder(request.Body).Decode(&newPerson)
	persons = append(persons, newPerson)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(persons)
}
