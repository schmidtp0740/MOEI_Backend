package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type blockchainRequest struct {
	Chaincode    string   `json:"chaincode"`
	Channel      string   `json:"channel"`
	ChaincodeVer string   `json:"chaincodeVer"`
	Method       string   `json:"method"`
	Args         []string `json:"args"`
}

type blockchainResponse struct {
	ReturnCode string `json:"returnCode"`
	Result     string `json:"result"`
	Info       string `json:"info,omitempty"`
}

// var rxList []Rx
// var rxLedger []Rx
// var rxCounter int

// FindAll ...
func FindAll() []Rx {
	// return rxLedger
	return nil
}

// FindAllRxForPatient ...
func FindAllRxForPatient(id string) (rxResponse []Rx) {

	// for _, rx := range rxList {
	// 	if rx.ID == id {
	// 		rxResponse = append(rxResponse, rx)
	// 	}

	// }
	return
}

// Insert ...
// func (rx *Rx) Insert(patientID string) (string, error) {

// 	return "", nil
// }

// Modify ...
// func (rx *Rx) Modify() bool {
// 	// for key, rxTemp := range rxList {
// 	// 	if rxTemp.RXID == rx.RXID {
// 	// 		rxT := Rx{
// 	// 			RXID:         rxTemp.RXID,
// 	// 			ID:           rxTemp.ID,
// 	// 			FirstName:    rxTemp.FirstName,
// 	// 			LastName:     rxTemp.LastName,
// 	// 			DOB:          rxTemp.DOB,
// 	// 			Prescription: rx.Prescription,
// 	// 			Refills:      rxTemp.Refills,
// 	// 			Doctor:       rxTemp.Doctor,
// 	// 			License:      rxTemp.License,
// 	// 			User:         rx.User,
// 	// 			Status:       rx.Status,
// 	// 			Insurance:    &ins{Company: rxTemp.Insurance.Company, PolicyID: rxTemp.Insurance.PolicyID, ExpirationDate: rxTemp.Insurance.ExpirationDate},
// 	// 			TimeStamp:    rx.TimeStamp,
// 	// 		}
// 	// 		rxList[key].Status = rx.Status
// 	// 		rxList[key].Prescription = rx.Prescription
// 	// 		rxLedger = append(rxLedger, rxT)
// 	// 	}
// 	// }
// 	return true
// }

func queryBlockchain(hostname, chaincode, channel, chaincodeVer, method string, args []string) (blockchainResponse, error) {
	url := hostname + "/bcsgw/rest/v1/transaction/query"

	payloadStruct := blockchainRequest{
		Chaincode:    chaincode,
		Channel:      channel,
		ChaincodeVer: chaincodeVer,
		Method:       method,
		Args:         args,
	}

	responseFromBlockchain, err := blockchainHandler(url, payloadStruct)
	if err != nil {
		return responseFromBlockchain, err

	}

	return responseFromBlockchain, nil
}

func invokeBlockchain(hostname, chaincode, channel, chaincodeVer, method string, args []string) (blockchainResponse, error) {
	url := hostname + "/bcsgw/rest/v1/transaction/invocation"

	payloadStruct := blockchainRequest{
		Chaincode:    chaincode,
		Channel:      channel,
		ChaincodeVer: chaincodeVer,
		Method:       method,
		Args:         args,
	}
	fmt.Println("invoke blockchain start")
	responseFromBlockchain, err := blockchainHandler(url, payloadStruct)
	if err != nil {
		return responseFromBlockchain, err

	}

	return responseFromBlockchain, nil

}

func blockchainHandler(url string, payloadStruct blockchainRequest) (blockchainResponse, error) {

	payloadAsBytes, err := json.Marshal(payloadStruct)
	if err != nil {
		return blockchainResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadAsBytes))
	if err != nil {
		return blockchainResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("about to send request")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error with blockchain query:" + err.Error())
		return blockchainResponse{}, err
	}
	defer resp.Body.Close()
	fmt.Println("got request")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return blockchainResponse{}, err
	}

	// create struct from blockchain response
	responseStruct := blockchainResponse{}

	fmt.Println("about to unmarshal")
	fmt.Println(responseStruct)
	if err := json.Unmarshal(body, &responseStruct); err != nil {
		fmt.Println("error with unmarshalling json: " + err.Error())
		return blockchainResponse{}, err
	}
	fmt.Println("unmarshaled")
	fmt.Println(responseStruct)

	return responseStruct, nil
}
