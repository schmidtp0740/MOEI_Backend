package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Rx ...
type Rx struct {
	PatientID    string `json:"patientID,omitempty"`
	RXID         string `json:"rxid,omitempty"`
	Prescription string `json:"prescription,omitempty"`
	Refills      int    `json:"refills,omitempty"`
	Doctor       string `json:"doctor,omitempty"`
	Pharmacist   string `json:"pharmacist,omitempty"`
	Status       string `json:"status,omitempty"`
	Timestamp    int    `json:"timestamp,omitempty"`
	Approved     string `json:"approved,omitempty"`
}

// ResponseFromBlockchain ...
type ResponseFromBlockchain struct {
	Status string `json:"status"`
	Info   string `json:"info,omitempty"`
}

type rxList struct {
	RX []Rx `json:"rx"`
}

// GetAllRx
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

// GetRx
// Input: id of a patient
// Output: All Rx for a patient
func GetRx(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["ID"]
	// rx := rxList{dao.FindAllRxForPatient(id)}
	// rxJSON, err := json.Marshal(rx)
	// if err != nil {
	// 	panic(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(rxJSON)
}

// InsertRx ...
// Input: rx data
// Output: success or failure
func InsertRx(w http.ResponseWriter, r *http.Request) {
	fmt.Println("insertingrx")
	request := Rx{}

	response := ResponseFromBlockchain{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error decoding payload:" + err.Error())

		response.Status = "Failure"
		response.Info = "Error: unable to decord request body"
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
			request.Pharmacist,
			request.Prescription,
			strconv.Itoa(request.Refills),
			request.Status,
			request.Approved,
		})

	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with invoking blockchain: " + result.Result)
		response.Info = result.Info
		response.Status = "Failure"
	}

	if response.Status != "Failure" {
		response.Status = "Success"
	}

	responseAsBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error marshalling response: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseAsBytes)

}

// ModifyRx
// Input: rx data (modified)
// Output: success or failure
func ModifyRx(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- starting modifyRx----")
	request := Rx{}

	response := ResponseFromBlockchain{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error decoding payload:" + err.Error())

		response.Status = "Failure"
		response.Info = "Error: unable to decord request body"
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
			request.Prescription,
			strconv.Itoa(request.Refills),
			request.Status,
			request.Approved,
		})

	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with invoking blockchain: " + result.Result)
		response.Info = result.Info
		response.Status = "Failure"
	}

	if response.Status != "Failure" {
		response.Status = "Success"
	}

	responseAsBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error marshalling response: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseAsBytes)

}
