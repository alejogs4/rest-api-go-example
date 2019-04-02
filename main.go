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

var cities = []City{
	{
		ID:         "01",
		Name:       "Medellín",
		State:      "Antioquia",
		Country:    "Colombia",
		Population: 3731447,
	},
	{
		ID:         "02",
		Name:       "Boston",
		State:      "Massachusetts",
		Country:    "United States",
		Population: 685094,
	},
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cities", GetCities).Methods("GET")
	router.HandleFunc("/cities/{id}", GetCity).Methods("GET")
	router.HandleFunc("/cities", CreateCity).Methods("POST")
	router.HandleFunc("/cities/{id}", DeleteCity).Methods("DELETE")

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
	var city City
	_ = json.NewDecoder(r.Body).Decode(&city)
	cities = append(cities, city)

	w.WriteHeader(http.StatusOk)
	json.NewEncoder(w).Encode(cities)
}

// Todo
func DeleteCity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, city := range cities {
		if city.ID == params["id"] {
			cities[index] = cities[len(cities)-1]
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cities)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(cities)

}
