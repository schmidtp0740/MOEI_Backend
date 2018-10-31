package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Rx ...
type Rx struct {
	PatientID    string `json:"patientID"`
	RXID         string `json:"rxid"`             // id of the prescription
	Timestamp    int    `json:"timestamp"`        // timestamp of when prescription was prescribed and filled
	Doctor       string `json:"doctor,omitempty"` // name of the doctor
	Pharmacist   string `json:"pharmacist,omitempty"`
	Prescription string `json:"prescription,omitempty"` // prescription name
	Refills      int    `json:"refills,emitempty"`      // number of refills
	ExpirateDate int    `json:"expDate,omitempty"`
	Status       string `json:"status,emitempty"` // current status of the prescription
	Approved     string `json:"approved,omitempty"`
}

// GetAllRx ...
// Input: none
// Output: list of a rx for all patients
func GetAllRx(w http.ResponseWriter, r *http.Request) {
	// rx := rxList{dao.FindAll()}

	// rxJSON, err := json.Marshal(rx)
	// if err != nil {
	// 	panic(err)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(rxJSON)
}

// GetRx ...
// Input: id of a patient
// Output: All Rx for a patient
func GetRx(w http.ResponseWriter, r *http.Request) {
	patientID := mux.Vars(r)["patientID"]

	blockVariable := getBlockchainVariables()

	result, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"getRxForPatient",
		[]string{
			patientID,
		})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with querying blockchain for rx: " + result.Info)
		result.Result = "error querying the blockchain" + result.Info

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Result))
}

// InsertRx ...
// Input: rx data
// Output: success or failure
func InsertRx(w http.ResponseWriter, r *http.Request) {
	fmt.Println("insertingrx")
	request := Rx{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error decoding payload:" + err.Error())
		response := BlockchainResponse{}
		response.Result = "Error: incorrect payload"
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response.Result))
		return
	}
	defer r.Body.Close()

	blockVariable := getBlockchainVariables()

	result, err := invokeBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"insertRx",
		[]string{
			request.PatientID,
			request.RXID,
			strconv.Itoa(request.Timestamp),
			request.Doctor,
			request.Prescription,
			strconv.Itoa(request.Refills),
			request.Status,
			request.Approved,
		})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with invoking blockchain: " + result.Info)
	}

	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error marshalling response: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultAsBytes)

}

// ModifyRx ...
// Input: rx data (modified)
// Output: success or failure
func ModifyRx(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- starting modifyRx----")
	request := Rx{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error decoding payload:" + err.Error())
		response := BlockchainResponse{}
		response.Result = "Error: incorrect payload"
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response.Result))
		return
	}
	defer r.Body.Close()

	blockVariable := getBlockchainVariables()

	result, err := invokeBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"modifyRx",
		[]string{
			request.PatientID,
			request.RXID,
			strconv.Itoa(request.Timestamp),
			request.Doctor,
			request.Pharmacist,
			request.Prescription,
			strconv.Itoa(request.Refills),
			request.Status,
			request.Approved,
		})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with invoking blockchain: " + result.Result)

	}

	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error marshalling response: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultAsBytes)

}
