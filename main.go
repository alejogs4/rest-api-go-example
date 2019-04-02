package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type City struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	Population int32  `json:"population,omitempty"`
}

var cities []City

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cities", GetCities).Methods("GET")
	router.HandleFunc("/city/{id}", GetCity).Methods("GET")
	router.HandleFunc("/city/{id}", CreateCity).Methods("POST")
	router.HandleFunc("/city/{id}", DeleteCity).Methods("DELETE")

	log.Print("Probando")
	http.ListenAndServe(":8000", router)
}

func GetCities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cities)
}

func GetCity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, city := range cities {
		if city.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(city)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&City{})
}

func CreateCity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var city City
	_ = json.NewDecoder(r.Body).Decode(&city)
	city.ID = params["id"]
	cities = append(cities, city)

	json.NewEncoder(w).Encode(city)
}

// Todo
func DeleteCity(w http.ResponseWriter, r *http.Request) {

}
