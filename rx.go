package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Rx ...
type Rx struct {
	PatientID    string `json:"patientID,omitempty"`
	RXID         string `json:"rxid,omitempty"`      // id of the prescription
	Timestamp    int    `json:"timestamp,omitempty"` // timestamp of when prescription was prescribed and filled
	Doctor       string `json:"doctor,omitempty"`    // name of the doctor
	Pharmacist   string `json:"pharmacist,omitempty"`
	Prescription string `json:"prescription,omitempty"` // prescription name
	Refills      int    `json:"refills,omitempty"`      // number of refills
	ExpirateDate int    `json:"expDate,omitempty"`
	Status       string `json:"status,omitempty"` // current status of the prescription
	Approved     string `json:"approved,omitempty"`
}

// GetAllRx ...
// Input: none
// Output: list of a rx for all patients
func GetAllRx(w http.ResponseWriter, r *http.Request) {

	rxList := struct {
		RXList []Rx `json:"rx"`
	}{}

	blockVariable := getBlockchainVariables()

	// setup conn to mysql
	db, err := sql.Open("mysql", "dbuser:userpass@tcp(mysql:3306)/myimagedb")
	if err != nil {
		fmt.Println("err setting up connextion")
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("unable to ping db: " + err.Error())
		return
	}

	// create go func

	// call getPeople to get update of patientID
	getPeopleResult, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"getPeople",
		[]string{})
	if err != nil || getPeopleResult.ReturnCode == "Failure" {
		fmt.Println("error with querying blockchain for rx: " + getPeopleResult.Info)
		return
	}

	people := struct {
		People []struct {
			PatientID string `json:"patientID,omitempty"`
			FirstName string `json:"firstName,omitempty"`
			LastName  string `json:"lastName,omitempty"`
		} `json:"people,omitempty"`
	}{}

	if err := json.Unmarshal([]byte(getPeopleResult.Result), &people); err != nil {
		fmt.Println("Error:unmarshalling people")
		return
	}

	// call getRxHistoryOfPatient to get all rx history for each patient
	for _, person := range people.People {
		getRxHistoryResult, err := queryBlockchain(blockVariable.Hostname,
			blockVariable.Chaincode,
			blockVariable.Channel,
			blockVariable.ChaincodeVer,
			"getRxHistoryOfPatient",
			[]string{
				person.PatientID,
			})
		if err != nil || getRxHistoryResult.ReturnCode == "Failure" {
			fmt.Println("error with querying blockchain for rxHistory: " + getRxHistoryResult.Info)
			return
		}

		rxHistoryResponse := struct {
			PatientID string `json:"patientID"`
			RxHistory [][]Rx `json:"rxHistory"`
		}{}

		if err := json.Unmarshal([]byte(getRxHistoryResult.Result), &rxHistoryResponse); err != nil {
			fmt.Println("error with unmarshalling rxhistory: " + getRxHistoryResult.Info)
			return
		}

		// store rxid, timestamp to patientID in redis
		// fmt.Printf("PatientID: %s\n", person.PatientID)
		for _, rxList := range rxHistoryResponse.RxHistory {
			for _, rx := range rxList {
				// fmt.Printf("RXID: %s\n", rx.RXID)
				// fmt.Printf("timestamp: %d\n", rx.Timestamp)
				// store patientID, rxid, and timestamp
				_, err = db.Exec("INSERT INTO rxlist VALUES(?, ?, ?)", person.PatientID, rx.RXID, rx.Timestamp)
				if err != nil {
					fmt.Println("error: " + err.Error())
				}
			}
		}
	}

	rows, err := db.Query("SELECT * FROM rxlist ORDER BY timestamp ASC")
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}

	var col1, col2, col3 []byte
	for rows.Next() {
		err = rows.Scan(&col1, &col2, &col3)

		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}

		// Use the string value
		fmt.Printf("Row: %s %s %s\n", string(col1), string(col2), string(col3))
		timestamp, err := strconv.Atoi(string(col3))
		if err != nil {
			fmt.Println("error unparsing int from string for timestamp")
		}
		tempRx := Rx{
			PatientID: string(col1),
			RXID:      string(col2),
			Timestamp: timestamp,
		}
		rxList.RXList = append(rxList.RXList, tempRx)

	}

	rxListAsBytes, err := json.Marshal(rxList)
	if err != nil {
		fmt.Println("unable to marshal rxList to bytes")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rxListAsBytes)
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
			strconv.Itoa(request.ExpirateDate),
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
			strconv.Itoa(request.ExpirateDate),
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
