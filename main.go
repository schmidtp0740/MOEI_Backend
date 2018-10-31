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

	// insert heart rate data for patient
	// TODO
	router.HandleFunc("/hr", insertHeartRateMessage).Methods("POST")

	// retrieve heart rate data for a patient
	// TODO
	router.HandleFunc("/hr/{patientID}", getHeartRateHistoryForPatient).Methods("GET")

	// retreive hack status
	// TODO
	router.HandleFunc("/bcs", GetStatus).Methods("GET")

	// push hack status
	// TODO
	router.HandleFunc("/hack", SetStatus).Methods("GET")

	// Get All Patient Data
	router.HandleFunc("/pd", GetPeople).Methods("GET")

	// Get Patient Data
	router.HandleFunc("/pd/{patientID}", GetPerson).Methods("GET")

	//Get All Rx Data History
	// TODO
	router.HandleFunc("/rx", GetAllRx).Methods("GET")

	// Get Rx Data
	router.HandleFunc("/rx/{patientID}", GetRx).Methods("GET")

	// Insert Rx
	router.HandleFunc("/rx", InsertRx).Methods("POST")

	// Modify Rx
	router.HandleFunc("/rx", ModifyRx).Methods("PATCH")

	// Get Insurance
	// TODO
	router.HandleFunc("/insurance/{patientID}", GetIns).Methods("GET")

	// New Insurance
	// TODO
	// router.HandleFunc("/insurance/{patientID}", NewIns).Methods("POST")

	fmt.Println("Listening on port: 8080")
	c := cors.AllowAll()
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
