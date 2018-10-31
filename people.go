package main

import (
	"fmt"
	"net/http"

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

	blockVariable := getBlockchainVariables()

	result, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"getPeople",
		[]string{})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with querying blockchain for rx: " + result.Info)
		result.Result = "error querying the blockchain" + result.Info

	}
	fmt.Println("Result from blockchain: ")
	fmt.Println("returnCode: " + result.ReturnCode)
	fmt.Println("Result: " + result.Result)
	fmt.Println("Info: " + result.Info)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Result))
}

// GetPerson ...
// Input: id of a patient
// Output: All data of a patient
func GetPerson(w http.ResponseWriter, r *http.Request) {

	patientID := mux.Vars(r)["patientID"]

	blockVariable := getBlockchainVariables()

	result, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"getPerson",
		[]string{
			patientID,
		})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with querying blockchain for rx: " + result.Info)
		result.Result = "error querying the blockchain" + result.Info
	}

	fmt.Println("Result from blockchain: ")
	fmt.Println("returnCode: " + result.ReturnCode)
	fmt.Println("Result: " + result.Result)
	fmt.Println("Info: " + result.Info)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Result))
}
