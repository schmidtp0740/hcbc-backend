package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var status = false

type payload struct {
	RXID       string `json:"rxid"`
	Status     string `json:"status"`
	Blockchain string `json:"blockchain"`
}

// GetStatus ...
func GetStatus(w http.ResponseWriter, r *http.Request) {

	blockVariable := getBlockchainVariables()

	result, err := queryBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"isHacked",
		[]string{})
	if err != nil || result.ReturnCode == "Failure" {
		fmt.Println("error with querying blockchain for rx: " + result.Info)
		result.Result = "error querying the blockchain" + result.Info
	}

	fmt.Println("Result from blockchain: ")
	fmt.Println("returnCode: " + result.ReturnCode)
	fmt.Println("Result: " + result.Result)
	fmt.Println("Info: " + result.Info)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result.Result))

}

// SetStatus ...
func SetStatus(w http.ResponseWriter, r *http.Request) {

	blockVariable := getBlockchainVariables()

	result, err := invokeBlockchain(blockVariable.Hostname,
		blockVariable.Chaincode,
		blockVariable.Channel,
		blockVariable.ChaincodeVer,
		"hack",
		[]string{})
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
