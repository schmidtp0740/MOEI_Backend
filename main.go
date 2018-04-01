package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/schmidtp0740/MOEI_Backend/hack"
	"github.com/schmidtp0740/MOEI_Backend/iot"
	"github.com/schmidtp0740/MOEI_Backend/people"
	"github.com/schmidtp0740/MOEI_Backend/rx"
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
	//TODO
	router.HandleFunc("/rx/{ID}", rx.GetRx).Methods("GET")

	// Insert Rx
	// TODO
	router.HandleFunc("/rx/{ID}", rx.InsertRx).Methods("POST")

	// Fill Rx
	// TODO
	router.HandleFunc("/rx/{ID}", rx.ModifyRx).Methods("PATCH")

	fmt.Println("Listening on port: 8000")
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
