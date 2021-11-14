package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//app.StartApplication()

	//fmt.Println("Welcome to the go-workshop, hope you will have fun!!")
	facts := repo{}
	facts.add(fact{
		Image:       "image1",
		Description: "description1",
	})
	facts.add(fact{
		Image:       "image2",
		Description: "description2",
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet{
			http.Error(w, "page is not found", http.StatusNotFound)
			return
		}
		_, err := fmt.Fprintf(w,"PONG")
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.HandleFunc("/facts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			factsList := facts.getAll()
			result, err := json.Marshal(factsList)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(result)
		case http.MethodPost:
			bodyRequest, err := ioutil.ReadAll(r.Body)
			if err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err = json.Unmarshal(bodyRequest, &factRequest)
			if err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newFact := fact{
				Image:       factRequest.Image,
				Description: factRequest.Description,
			}
			facts.add(newFact)
			w.Write([]byte("success"))
		default:
			http.Error(w, "page is not found", http.StatusNotFound)
			return
		}
	})



	http.ListenAndServe(":9002", nil)

}
