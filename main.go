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
	router.HandleFunc("/hr", insertHeartRateMessage).Methods("POST")

	// retrieve heart rate data for a patient
	router.HandleFunc("/hr/{patientID}", getHeartRateHistoryForPatient).Methods("GET")

	// retreive hack status
	router.HandleFunc("/bcs", GetStatus).Methods("GET")

	// push hack status
	router.HandleFunc("/hack", SetStatus).Methods("GET")

	// Get All Patient Data
	router.HandleFunc("/pd", GetPeople).Methods("GET")

	// Get Patient Data
	router.HandleFunc("/pd/{patientID}", GetPerson).Methods("GET")

	//Get All Rx Data History
	router.HandleFunc("/rx", GetAllRx).Methods("GET")

	// Get Rx Data
	router.HandleFunc("/rx/{patientID}", GetRx).Methods("GET")

	// Insert Rx
	router.HandleFunc("/rx", InsertRx).Methods("POST")

	// Fill Rx
	router.HandleFunc("/rx", FillRx).Methods("PATCH")

	// Approve Rx
	router.HandleFunc("/rx", ApproveRx).Methods("PUT")

	// Get Insurance
	router.HandleFunc("/insurance/{patientID}", GetInsurance).Methods("GET")

	// New Insurance
	router.HandleFunc("/insurance", insertInsurance).Methods("POST")

	// get blood pressure history
	router.HandleFunc("/bp/{patientID}", getBloodPressure).Methods("GET")

	// insert blood pressure message
	router.HandleFunc("/bp", insertBloodPressure).Methods("POST")

	fmt.Println("Listening on port: 8080")
	c := cors.AllowAll()
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
