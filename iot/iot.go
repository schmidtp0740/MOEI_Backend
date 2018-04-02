package iot

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/schmidtp0740/MOEI_BACKEND/request"
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

	url := os.Getenv("URL") + "/bcsgw/rest/v1/transaction/invocation"

	id := data.ID
	heartRate := strconv.Itoa(data.HeartRate)
	unit := data.Unit
	timeStamp := strconv.Itoa(data.TimeStamp)

	m := []byte(`{ 
		"channel": "mychannel", 
		"chaincode": "hrcc", 
		"chaincodeVer": "v2", 
		"method": "insertData",	
		"args": ["` + id + `", "` +
		heartRate + `","` +
		unit + `","` +
		timeStamp + `"]}`,
	)

	body := request.Request(m, url)

	//fmt.Println(body)

	json.NewEncoder(w).Encode(body)
}

// GetData ...
func GetData(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("URL") + "/bcsgw/rest/v1/transaction/query"
	m := []byte(`{
		"channel": "mychannel",
		"chaincode": "hrcc",
		"chaincodeVer": "v3",
		"method": "getHistory",
		"args": ["001"]
	}`)

	body := request.Request(m, url)

	//fmt.Println(body)

	json.NewEncoder(w).Encode(body)
}
