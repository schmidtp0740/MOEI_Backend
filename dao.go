package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BlockchainRequest ...
type BlockchainRequest struct {
	Chaincode    string   `json:"chaincode"`
	Channel      string   `json:"channel"`
	ChaincodeVer string   `json:"chaincodeVer"`
	Method       string   `json:"method"`
	Args         []string `json:"args"`
}

// BlockchainResponse ...
type BlockchainResponse struct {
	ReturnCode string `json:"returnCode"`
	Result     string `json:"result"`
	Info       string `json:"info,omitempty"`
}

func queryBlockchain(hostname, chaincode, channel, chaincodeVer, method string, args []string) (BlockchainResponse, error) {
	url := hostname + "/bcsgw/rest/v1/transaction/query"

	payloadStruct := BlockchainRequest{
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

func invokeBlockchain(hostname, chaincode, channel, chaincodeVer, method string, args []string) (BlockchainResponse, error) {
	url := hostname + "/bcsgw/rest/v1/transaction/invocation"

	payloadStruct := BlockchainRequest{
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

func blockchainHandler(url string, payloadStruct BlockchainRequest) (BlockchainResponse, error) {

	payloadAsBytes, err := json.Marshal(payloadStruct)
	if err != nil {
		return BlockchainResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadAsBytes))
	if err != nil {
		return BlockchainResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("about to send request")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error with blockchain query:" + err.Error())
		return BlockchainResponse{}, err
	}
	defer resp.Body.Close()
	fmt.Println("got request")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return BlockchainResponse{}, err
	}

	// create struct from blockchain response
	responseStruct := BlockchainResponse{}

	fmt.Println("about to unmarshal")
	fmt.Println(responseStruct)
	if err := json.Unmarshal(body, &responseStruct); err != nil {
		fmt.Println("error with unmarshalling json: " + err.Error())
		return BlockchainResponse{}, err
	}
	fmt.Println("unmarshaled")
	fmt.Println(responseStruct)

	return responseStruct, nil
}
