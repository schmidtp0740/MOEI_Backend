package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Rx ...
type Rx struct {
	RXID         string `json:"RXID"`
	ID           string `json:"ID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	DOB          string `json:"DOB"`
	Prescription string `json:"Prescription"`
	Refills      int    `json:"Refills"`
	Doctor       string `json:"Doctor"`
	License      string `json:"License"`
	Status       string `json:"Status"`
	TimeStamp    string `json:"TimeStamp"`
}

// FindAll ...
func FindAll() []Rx {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	c := session.DB("RxService").C("Rx")
	var rx []Rx
	err = c.Find(bson.M{}).All(&rx)
	return rx
}

// Insert ...
func (rx *Rx) Insert() bool {
	session, err := mgo.Dial("Localhost:27017")
	if err != nil {
		log.Println("Could no connect to Mongo: ", err.Error())
		return false
	}
	defer session.Close()

	c := session.DB("RxService").C("Rx")
	_, err = c.UpsertId(rx.RXID, rx)
	if err != nil {
		log.Println("Error creating Profile: ", err.Error())
		return false
	}
	return true
}
