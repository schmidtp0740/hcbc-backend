package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/schmidtp0740/medbo-backend/pkg"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler).Methods("GET")

	// insert heart rate data for patient
	router.HandleFunc("/hr", pkg.InsertHeartRateMessage).Methods("POST")

	// // retrieve heart rate data for a patient
	router.HandleFunc("/hr/{patientID}", pkg.GetHeartRateHistoryForPatient).Methods("GET")

	// retreive hack status
	router.HandleFunc("/bcs", pkg.GetStatus).Methods("GET")

	// push hack status
	router.HandleFunc("/hack", pkg.SetStatus).Methods("GET")

	// Get All Patient Data
	router.HandleFunc("/pd", pkg.GetPeople).Methods("GET")

	// Get Patient Data
	router.HandleFunc("/pd/{patientID}", pkg.GetPerson).Methods("GET")

	//Get All Rx Data History
	router.HandleFunc("/rxledger", pkg.GetAllRx).Methods("GET")

	// Get Rx Data
	router.HandleFunc("/rx/{patientID}", pkg.GetRx).Methods("GET")

	// Insert Rx
	router.HandleFunc("/rx", pkg.InsertRx).Methods("POST")

	// Fill Rx
	router.HandleFunc("/rx", pkg.FillRx).Methods("PATCH")

	// Approve Rx
	router.HandleFunc("/rx", pkg.ApproveRx).Methods("PUT")

	// Get Insurance
	router.HandleFunc("/insurance/{patientID}", pkg.GetInsurance).Methods("GET")

	// New Insurance
	router.HandleFunc("/insurance", pkg.InsertInsurance).Methods("POST")

	// get blood pressure history
	router.HandleFunc("/bp/{patientID}", pkg.GetBloodPressure).Methods("GET")

	// insert blood pressure message
	router.HandleFunc("/bp", pkg.InsertBloodPressure).Methods("POST")

	fmt.Println("Listening on port: 8080")
	c := cors.AllowAll()
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`medbo-backend`))
}
