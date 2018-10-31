package main

import "os"

// BlockchainVariables ...
type BlockchainVariables struct {
	Hostname     string   `json:"hostname"`
	Channel      string   `json:"channel"`
	Chaincode    string   `json:"chaincode"`
	ChaincodeVer string   `json:"chaincodeVer"`
	Method       string   `json:"method"`
	Args         []string `json:"args"`
}

func getBlockchainVariables() BlockchainVariables {
	t := BlockchainVariables{
		Hostname:     os.Getenv("hostname"),
		Chaincode:    os.Getenv("chaincode"),
		ChaincodeVer: os.Getenv("chaincodeVer"),
		Channel:      os.Getenv("channel"),
	}

	return t
}
