package main

import (
	"net/http"
)

type ins struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	SSN            string `json:"ssn"`
	Company        string `json:"company"`
	PolicyID       string `json:"policyId"`
	ExpirationDate string `json:"expirationDate"`
}

// GetIns ....
func GetIns(w http.ResponseWriter, r *http.Request) {
	// var soacsURL string

	// if os.Getenv("SOA") != "" {
	// 	soacsURL = os.Getenv("SOA")
	// } else {
	// 	soacsURL = "http://private-e5e0b-ironbankbcsapidoc.apiary-mock.com/insurancesoap"
	// }
	// id := mux.Vars(r)["ID"]
	// var firstName, lastName string

	// body := `<soapenv:Envelope     xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	// <soap:Header     xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	// </soap:Header>
	// <soapenv:Body>
	// <web:getFromDB     xmlns:web="http://webserviceapp/"/>
	// </soapenv:Body>
	// </soapenv:Envelope>`

	// //fmt.Println("Body", body)

	// req, err := http.NewRequest("POST", soacsURL, strings.NewReader(body))
	// req.Method = "POST"
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// req.Header.Add("Content-Type", "text/xml;charset=UTF-8")

	// client := http.Client{}

	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// b, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("response: ", string(b))

	// re := regexp.MustCompile(`policy+`)
	// responseParsedString := re.ReplaceAllString(string(b), `policyId`)
	// fmt.Println("\nresponse:", responseParsedString)

	// re = regexp.MustCompile(`exp+`)
	// responseParsedString = re.ReplaceAllString(responseParsedString, `expirationDate`)
	// fmt.Println("\nResponse:", responseParsedString)

	// re = regexp.MustCompile(`({[":\s\,0-9\-\.A-Za-z\/]*})+`)
	// responseParsedStrings := re.FindAllString(responseParsedString, -1)
	// var insuranceJSON []byte
	// for _, insuranceIterator := range responseParsedStrings {
	// 	var insurance ins
	// 	err = json.Unmarshal([]byte(insuranceIterator), &insurance)
	// 	if err != nil {
	// 		println(err)
	// 	}
	// 	var breakValue = false

	// 	for _, person := range people.People {
	// 		if person.ID == id {
	// 			firstName = person.FirstName
	// 			lastName = person.LastName
	// 		}

	// 		if strings.ToLower(firstName+" "+lastName) == strings.ToLower(insurance.Name) {
	// 			fmt.Println(insurance)
	// 			insuranceJSON, err = json.Marshal(insurance)
	// 			if err != nil {
	// 				println(err)
	// 			}
	// 			breakValue = true
	// 			break

	// 		}
	// 	}
	// 	if breakValue {
	// 		break
	// 	}

	// }

	// if insuranceJSON == nil {
	// 	errStruct := map[string]string{"error": "insurance not found"}
	// 	insuranceJSON, err = json.Marshal(errStruct)
	// 	if err != nil {
	// 		println(err)
	// 	}
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(insuranceJSON)

}
