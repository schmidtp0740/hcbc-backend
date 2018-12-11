package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type heartRateRequest struct {
	PatientID string `json:"patientID"`
	HeartRate int    `json:"heartRate,omitempty"` // heart rate of the patient
	Timestamp int    `json:"timestamp,omitempty"` // timestamp of the record
}

// InsertHeartRateMessage ...
func InsertHeartRateMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("insertHeartRateMessage")
	request := heartRateRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error decoding payload:" + err.Error())
		response := BlockchainResponse{}
		response.Result = "Error: incorrect payload"
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response.Result))
		return
	}
	defer r.Body.Close()

	blockVariable := getBlockchainVariables()

	result, err := invokeBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"newHeartRateMessage",
		[]string{
			request.PatientID,
			strconv.Itoa(request.HeartRate),
			strconv.Itoa(request.Timestamp),
		})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with invoking blockchain: " + result.Info)
	}

	resultAsBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println("error marshalling response: " + err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultAsBytes)
}

// GetHeartRateHistoryForPatient ...
func GetHeartRateHistoryForPatient(w http.ResponseWriter, r *http.Request) {

	patientID := mux.Vars(r)["patientID"]

	blockVariable := getBlockchainVariables()

	result, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"getHeartRateHistory",
		[]string{
			patientID,
		})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with querying blockchain for rx: " + result.Info)
		result.Result = "error querying the blockchain" + result.Info

	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Result))

}
