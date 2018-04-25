package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/schmidtp0740/moei_backend/hack"
	"github.com/schmidtp0740/moei_backend/ins"
	"github.com/schmidtp0740/moei_backend/iot"
	"github.com/schmidtp0740/moei_backend/people"
	"github.com/schmidtp0740/moei_backend/rx"
)

func main() {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/data", iot.SendData).Methods("POST")
	router.HandleFunc("/getData", iot.GetData).Methods("GET")
	router.HandleFunc("/bcs", hack.GetStatus).Methods("GET")
	router.HandleFunc("/hack", hack.SetStatus).Methods("GET")

	// Get All Patient Data
	router.HandleFunc("/pd", people.GetPeople).Methods("GET")

	// Get Patent Data
	router.HandleFunc("/pd/{FirstName}/{LastName}", people.GetPerson).Methods("GET")

	//Get All Rx Data
	router.HandleFunc("/rx", rx.GetAllRx).Methods("GET")

	// Get Rx Data
	router.HandleFunc("/rx/{ID}", rx.GetRx).Methods("GET")

	// Insert Rx
	router.HandleFunc("/rx/{ID}", rx.InsertRx).Methods("POST")

	// Fill Rx
	// TODO
	router.HandleFunc("/rx/{ID}", rx.ModifyRx).Methods("PATCH")

	// Get Insurance
	router.HandleFunc("/insurance/{ID}", ins.GetIns).Methods("GET")

	fmt.Println("Listening on port: 8000")
	c := cors.AllowAll()
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
