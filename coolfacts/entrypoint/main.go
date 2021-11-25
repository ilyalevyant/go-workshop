package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	//app.StartApplication()

	//fmt.Println("Welcome to the go-workshop, hope you will have fun!!")

	facts, err := mentalfloss{}.Facts()
	if err != nil {
		fmt.Printf("can't reach mentalfloss: ", err)
	}

	//facts.add(fact{
	//	Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
	//	Description: "Did you know sonic is a hedgehog?!",
	//})
	//facts.add(fact{
	//	Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
	//	Description: "You won't believe what happened to Arya!",
	//})

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
			//factsList := facts.getAll()
			//result, err := json.Marshal(factsList)
			//if err != nil{
			//	http.Error(w, err.Error(), http.StatusInternalServerError)
			//	return
			//}
			//w.Write(result)
			tmpl, err := template.New("facts").Parse(newsTemplate)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//facts := facts.getAll()
			err = tmpl.Execute(w, facts)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return


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
			newFact := fact.Fact{
				Image:       factRequest.Image,
				Description: factRequest.Description,
			}
			facts = append(facts, newFact)
			w.Write([]byte("success"))
		default:
			http.Error(w, "page is not found", http.StatusNotFound)
			return
		}
	})

	http.ListenAndServe(":9002", nil)
}


