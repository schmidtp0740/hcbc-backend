package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetInsurance ....
func GetInsurance(w http.ResponseWriter, r *http.Request) {

	patientID := mux.Vars(r)["patientID"]

	blockVariable := getBlockchainVariables()

	result, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"getInsurance",
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

// InsertInsurance ...
func InsertInsurance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("insertingrx")

	request := struct {
		PatientID      string `json:"patientID"`
		Name           string `json:"insuranceName,omitempty"`
		ExpirationDate int    `json:"expDate,omitempty"`
		PolicyID       string `json:"policyID,omitempty"`
	}{}

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
		"insertInsurance",
		[]string{
			request.PatientID,
			request.Name,
			strconv.Itoa(request.ExpirationDate),
			request.PolicyID,
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
