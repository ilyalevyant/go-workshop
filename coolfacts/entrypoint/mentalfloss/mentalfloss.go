package mentalfloss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type mentalfloss struct{}

var items []struct {
	FactText     string `json:"fact"`
	PrimaryImage string `json:"primaryImage"`
}


func (mf mentalfloss) Facts() ([]fact.Fact, error) {
	resp, err := http.Get("http://mentalfloss.com/api/facts")
	if err != nil {
		fmt.Printf("error to get result from mentalFloss ", err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyRequest, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Printf("error to read results from mentalFloss ", err)
		return nil, err
	}
	err = json.Unmarshal(bodyRequest, &items)
	if err != nil{
		fmt.Printf("error to unmarshall results from mentalFloss ", err)
		return nil, err
	}
	var facts []fact

	for _, element := range items {
		newFact := fact{
			Image:       element.PrimaryImage,
			Description: element.FactText,
		}
		facts = append(facts, newFact)
	}
	return facts, nil
}

