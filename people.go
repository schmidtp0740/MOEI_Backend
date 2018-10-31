package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// type person struct {
// 	PatientID        string    `json:"patientID"`
// 	FirstName string `json:"firstName,omitempty"`
// 	LastName  string `json:"lastName,omitempty"`
// 	DOB       string `json:"dob,omitempty"`
// 	Address   string `json:"address,omitempty"`
// 	Ethnicity string `json:"ethnicity,omitempty"`
// 	Phone     string `json:"phone,omitempty"`
// }

// GetPeople ...
// Input: none
// Output: id, first name and lastname for all patients
func GetPeople(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("hostname")
	chaincode := os.Getenv("chaincode")
	chaincodeVer := os.Getenv("chaincodeVer")
	channel := os.Getenv("channel")

	// initialize a struct of people
	// people := struct {
	// 	People []person `json:"people"`
	// }{}

	result, err := queryBlockchain(hostname, chaincode, channel, chaincodeVer, "getPeople", []string{})
	if err != nil {
		w.Write([]byte("error query blockchain" + err.Error()))
	}
	fmt.Println(string(result.Result))

	type person struct {
		PatientID string `json:"patientID"`
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
	}

	people := struct {
		People []person `json:"people"`
	}{}

	if err := json.Unmarshal([]byte(result.Result), &people); err != nil {
		fmt.Println(err.Error())
	}

	peeopleJSON, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(peeopleJSON)
}

// GetPerson ...
// Input: id of a patient
// Output: All data of a patient
func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	patientID := vars["patientID"]

	hostname := os.Getenv("hostname")
	chaincode := os.Getenv("chaincode")
	chaincodeVer := os.Getenv("chaincodeVer")
	channel := os.Getenv("channel")

	result, err := queryBlockchain(hostname, chaincode, channel, chaincodeVer, "getPerson", []string{patientID})
	if err != nil {
		w.Write([]byte("error query blockchain" + err.Error()))
	}
	fmt.Println(string(result.Result))

	person := struct {
		PatientID string `json:"patientID"`
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		DOB       string `json:"dob,omitempty"`
		Address   string `json:"address,omitempty"`
		Ethnicity string `json:"ethnicity,omitempty"`
		Phone     string `json:"phone,omitempty"`
	}{}

	if err := json.Unmarshal([]byte(result.Result), &person); err != nil {
		fmt.Println(err.Error())
	}

	personAsBytes, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(personAsBytes)
}
