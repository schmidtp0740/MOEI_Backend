package hack

import (
	"bytes"
	"encoding/json"
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
	var payload payload

	if status {
		payload.RXID = "RX001"
		payload.Blockchain = "Doctor"
		payload.Status = "True"
	} else {
		payload.Status = "False"

	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payloadJSON)

}

// SetStatus ...
func SetStatus(w http.ResponseWriter, r *http.Request) {

	if status {
		status = false
	} else {
		status = true
	}

	var payload = []byte(`{}`)
	b := bytes.NewBuffer(payload)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}
