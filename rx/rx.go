package rx

import (
	"encoding/json"
	"net/http"

	"github.com/schmidtp0740/MOEI_Backend/dao"
)

// GetAllRx ...
func GetAllRx(w http.ResponseWriter, r *http.Request) {
	rx := dao.FindAll()
	rxJSON, err := json.Marshal(rx)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rxJSON)
}

// GetRx ...
func GetRx(w http.ResponseWriter, r *http.Request) {
	rx := dao.FindRx()
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

	rx.Insert()

	rxJSON, err := json.Marshal(rx)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rxJSON)

}

//ModifyRx ...
func ModifyRx(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var rx dao.Rx
	if err := json.NewDecoder(r.Body).Decode(&rx); err != nil {
		panic(err)
	}
	rx.Modify()

	rxJSON, err := json.Marshal(rx)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rxJSON)

}
