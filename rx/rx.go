package rx

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schmidtp0740/moei_backend/dao"
)

type rxList struct {
	RX []dao.Rx `json:"RX"`
}

type response struct {
	Response string `json:"response"`
}


// GetAllRx ...
func GetAllRx(w http.ResponseWriter, r *http.Request) {
	rx := rxList{dao.FindAll()}

	rxJSON, err := json.Marshal(rx)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rxJSON)
}

// GetRx ...
func GetRx(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	rx := rxList{dao.FindAllRxForPatient(id)}
	rxJSON, err := json.Marshal(rx)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rxJSON)
}

// InsertRx ...
func InsertRx(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rx dao.Rx
	if err := json.NewDecoder(r.Body).Decode(&rx); err != nil {
		panic(err)
	}

	id := mux.Vars(r)["ID"]
	_ = rx.Insert(id)

	resp := response{
		"ok",
	}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)

}

//ModifyRx ...
func ModifyRx(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rx dao.Rx
	if err := json.NewDecoder(r.Body).Decode(&rx); err != nil {
		panic(err)
	}

	_ = rx.Modify()

	resp := response{
		"ok",
	}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)

}
