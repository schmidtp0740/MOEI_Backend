package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// insert iot data
	// TODO
	router.HandleFunc("/data", SendData).Methods("POST")

	// retrieve iot data
	// TODO
	router.HandleFunc("/getData", GetData).Methods("GET")

	// retreive hack status
	// TODO
	router.HandleFunc("/bcs", GetStatus).Methods("GET")

	// push hack status
	// TODO
	router.HandleFunc("/hack", SetStatus).Methods("GET")

	// Get All Patient Data
	router.HandleFunc("/pd", GetPeople).Methods("GET")

	// Get Patent Data
	router.HandleFunc("/pd/{patientID}", GetPerson).Methods("GET")

	//Get All Rx Data History
	// TODO
	router.HandleFunc("/rx", GetAllRx).Methods("GET")

	// Get Rx Data
	// TODO
	router.HandleFunc("/rx/{patientID}", GetRx).Methods("GET")

	// Insert Rx
	router.HandleFunc("/rx", InsertRx).Methods("POST")

	// Modify Rx
	// TESTED FAILED
	router.HandleFunc("/rx", ModifyRx).Methods("PATCH")

	// Get Insurance
	// TODO
	router.HandleFunc("/insurance/{ID}", GetIns).Methods("GET")

	fmt.Println("Listening on port: 8080")
	c := cors.AllowAll()
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
