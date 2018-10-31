package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type data struct {
	ID        string `json:"id"`
	HeartRate int    `json:"heartRate"`
	Unit      string `json:"unit"`
	TimeStamp int    `json:"timeStamp"`
}

// SendData ...
func SendData(w http.ResponseWriter, r *http.Request) {

	var data data

	json.NewDecoder(r.Body).Decode(&data)

	url := "http://129.213.52.239:4001/bcsgw/rest/v1/transaction/invocation"

	id := data.ID
	heartRate := strconv.Itoa(data.HeartRate)
	unit := data.Unit
	timeStamp := strconv.Itoa(data.TimeStamp)

	m := []byte(`{
		"channel": "mychannel",
		"chaincode": "iotcc",
		"chaincodeVer": "v3",
		"method": "insertData",
		"args": ["` + id + `", "` +
		heartRate + `","` +
		unit + `","` +
		timeStamp + `"]}`,
	)

	body := Request(m, url)

	//fmt.Println(body)

	json.NewEncoder(w).Encode(body)
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
