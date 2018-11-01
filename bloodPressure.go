package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getBloodPressure(w http.ResponseWriter, r *http.Request) {

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

func insertBloodPressure(w http.ResponseWriter, r *http.Request) {

}
