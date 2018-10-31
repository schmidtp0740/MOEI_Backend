package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type heartRateRequest struct {
	PatientID string `json:"patientID"`
	HeartRate int    `json:"heartRate,omitempty"` // heart rate of the patient
	Timestamp int    `json:"timestamp,omitempty"` // timestamp of the record
}

// insertHeartRateMessage ...
func insertHeartRateMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("insertHeartRateMessage")
	request := heartRateRequest{}

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
			strconv.Itoa(request.HeartRate),
			strconv.Itoa(request.Timestamp),
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

// GetData ...
func GetData(w http.ResponseWriter, r *http.Request) {

	url := "http://129.213.52.239:4001/bcsgw/rest/v1/transaction/query"
	m := []byte(`{
		"channel": "mychannel",
		"chaincode": "iotcc",
		"chaincodeVer": "v3",
		"method": "getHistory",
		"args": ["001"]
	}`)

	body := Request(m, url)

	//fmt.Println(body)

	json.NewEncoder(w).Encode(body)
}
