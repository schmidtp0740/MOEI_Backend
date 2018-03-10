package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type data struct {
	ID        string `json:"id"`
	HeartRate int    `json:"heartRate"`
	Unit      string `json:"unit"`
	TimeStamp int    `json:"timeStamp"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/data", SendData).Methods("POST")

	fmt.Println("Listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// SendData ...
func SendData(w http.ResponseWriter, r *http.Request) {

	var data data

	json.NewDecoder(r.Body).Decode(&data)

	url := "http://129.146.106.151:4001/bcsgw/rest/v1/transaction/invocation"

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

	body := request(m, url)

	fmt.Println(string(body))

	json.NewEncoder(w).Encode(body)
}

func request(m []byte, url string) string {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(m))
	if err != nil {
		panic(err)
	}

	fmt.Println("payload: ", string(m))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
