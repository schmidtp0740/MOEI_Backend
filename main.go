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
	"github.com/rs/cors"
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
	router.HandleFunc("/getData", getData).Methods("GET")

	fmt.Println("Listening on port: 8000")
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
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
		"chaincodeVer": "v1", 
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

func getData(w http.ResponseWriter, r *http.Request) {
	url := "http://129.146.106.151:4001/bcsgw/rest/v1/transaction/query"

	m := []byte(`{
		"channel": "mychannel",
		"chaincode": "hrcc",
		"chaincodeVer": "v1",
		"method": "getHistory",
		"args": ["001"]
	}`)
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

	fmt.Println(string(body))

	json.NewEncoder(w).Encode(string(body))
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
