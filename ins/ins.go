package ins

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/schmidtp0740/moei_backend/people"
)

// type ins struct {
// 	insuranceCompany string `json:"insuranceCompany"`
// 	policyID         string `json:"policyId"`
// 	expirationDate   string `json:"expirationDate"`
// }
type envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		Company struct {
			XMLName xml.Name `xml:"company"`
			Text    string   `xml:",chardata"`
		}
		PolicyID struct {
			XMLName xml.Name `xml:"policyId"`
			Text    string   `xml:",chardata"`
		}
		ExpirationDate struct {
			XMLName xml.Name `xml:"expirationDate"`
			Text    string   `xml:",chardata"`
		}
	}
}

type ins struct {
	Company        string `json:"insuranceCompany"`
	PolicyID       string `json:"policyId"`
	ExpirationDate string `json:"expirationDate"`
}

const soacsURL = "http://private-e5e0b-ironbankbcsapidoc.apiary-mock.com/insurancesoap"

// GetIns ....
func GetIns(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["ID"]
	var firstName, lastName string
	for _, person := range people.People {
		if person.ID == id {
			firstName = person.FirstName
			lastName = person.LastName
		}
	}
	body := `<?xml version="1.0" encoding="UTF-8"?>` +
		`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">` +
		`<soap:Header/>` +
		`<soap:Body>` +
		`<ns1:submit xmlns:ns1="http://webservice.com/">` +
		`<arg0>` + firstName + ` ` + lastName + `</arg1>` +
		`</ns1:submit>` +
		`</soap:Body>` +
		`<soap:Envelope>`
	fmt.Println("Body", body)
	req, err := http.NewRequest("POST", soacsURL, strings.NewReader(body))
	req.Method = "POST"
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", "submit")
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response: ", string(b))

	env := new(envelope)
	err = xml.Unmarshal(b, env)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("env.body.getInsurance", env.Body.PolicyID)

	ins := ins{env.Body.Company.Text, env.Body.PolicyID.Text, env.Body.ExpirationDate.Text}
	insJSON, err := json.Marshal(ins)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(insJSON)
}
