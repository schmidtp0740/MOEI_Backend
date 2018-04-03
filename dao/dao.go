package dao

import (
	"strconv"
)

// Rx ...
type Rx struct {
	RXID         string `json:"RXID"`
	ID           string `json:"ID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	DOB          string `json:"DOB"`
	Prescription string `json:"Prescription"`
	Refills      string `json:"Refills"`
	Doctor       string `json:"Doctor"`
	License      string `json:"License"`
	Status       string `json:"Status"`
	TimeStamp    int    `json:"TimeStamp"`
}

var rxList []Rx
var rxLedger []Rx
var rxCounter int

// FindAll ...
func FindAll() []Rx {
	return rxLedger
}

// FindAllRxForPatient ...
func FindAllRxForPatient(id string) (rxResponse []Rx) {

	for _, rx := range rxList {
		if rx.ID == id {
			rxResponse = append(rxResponse, rx)
		}

	}
	return
}

// Insert ...
func (rx *Rx) Insert(id string) bool {
	rxCounter++
	rx.RXID = "RX" + strconv.Itoa(rxCounter)
	rx.ID = id
	rxList = append(rxList, *rx)
	rxLedger = append(rxLedger, *rx)
	return true
}

// Modify ...
func (rx *Rx) Modify() bool {
	for _, rxTemp := range rxList {
		if rxTemp.RXID == rx.RXID {
			rxT := Rx{
				RXID: rxTemp.RXID,
				ID: rxTemp.ID,
				FirstName: rxTemp.FirstName,
				LastName: rxTemp.LastName,
				DOB: rxTemp.DOB,
				Prescription: rxTemp.Prescription,
				Refills: rxTemp.Refills,
				Doctor: rxTemp.Doctor,
				License: rxTemp.License,
				Status: rx.Status,
				TimeStamp: rx.TimeStamp,
			}
			rxTemp.Status = rx.Status
			rxLedger = append(rxLedger, rxT)
		}
	}
	return true
}
