package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Saint struct {
	ID    string
	Name  string
	Cloth string
	Type  string
	Power int
}

var data = []Saint{
	Saint{"01", "Seiya", "Pegasus", "Bronze", 80},
	Saint{"02", "Hyoga", "Cygnus", "Bronze", 75},
	Saint{"03", "Shiryu", "Dragon", "Bronze", 82},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {

		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/saints", users)
	http.HandleFunc("/saint", user)

	fmt.Println("starting web server at http://localhost:8088/")
	http.ListenAndServe(":8088", nil)
}
