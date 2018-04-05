package people

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Ethnicity string `json:"ethnicity"`
	Phone     string `json:"phone"`
}

type persons struct {
	Persons []person `json:"persons"`
}

// People ...
var People = []person{
	{
		ID:        "001",
		FirstName: "John",
		LastName:  "Doe",
		DOB:       "01/01/1987",
		Address:   "999 Denver Rd, Portland, OR 98765",
		Ethnicity: "Asian",
		Phone:     "111-111-1111",
	},
	{
		ID:        "002",
		FirstName: "Mary",
		LastName:  "Jane",
		DOB:       "05/05/1997",
		Address:   "111 Denver Rd, Portland, OR 98765",
		Ethnicity: "Caucasion",
		Phone:     "123-123-1234",
	},
}

// GetPeople ...
func GetPeople(w http.ResponseWriter, r *http.Request) {

	personJSON, err := json.Marshal(persons{People})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(personJSON)
}

// GetPerson ...
func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["FirstName"])
	var response person
	for _, value := range People {
		if strings.ToLower(vars["FirstName"]) == strings.ToLower(value.FirstName) && strings.ToLower(vars["LastName"]) == strings.ToLower(value.LastName) {
			response = value
		}
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
