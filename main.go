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

type blockchainCall struct {
	channel      string
	chaincode    string
	chaincodeVer string
	method       string
	args         []string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/data", SendData).Methods("POST")

	fmt.Println("Listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// SendData ...
func SendData(w http.ResponseWriter, r *http.Request) {

	handler(w, r, "insertData")

}

func blockchainRequest(m blockchainCall, url string) string {

	b, err := json.Marshal(m)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func handler(w http.ResponseWriter, r *http.Request, action string) {
	var data data

	json.NewDecoder(r.Body).Decode(&data)

	url :=  "http://129.146.106.151:4001/bcsgw/rest/v1/transaction/invocation"

	id := data.ID
	heartRate := strconv.Itoa(data.HeartRate)
	unit := data.Unit
	timeStamp := strconv.Itoa(data.TimeStamp)

	m := blockchainCall{
		"mychannel",
		"emrCC",
		"v1",
		action,
		[]string{
			id, heartRate,
			unit, timeStamp,
		},
	}

	body := blockchainRequest(m, url)

	json.NewEncoder(w).Encode(body)

	fmt.Printf("Response from blockchain: %s\n\n", body)

}
