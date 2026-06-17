package main 

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Men struct {
	Id   string `json:"id"`
	Name string `json:"name"`//if you are not capitalising these fields you will not get them
}

func GetTrack(w http.ResponseWriter, r *http.Request) {
	track := Men {
		Id: "aalu",
		Name: "aakash",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(track)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
}

func CreateTrack(w http.ResponseWriter, r *http.Request) {
	var track Men

	err := json.NewDecoder(r.Body).Decode(&track)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	fmt.Println(track.Id)
}

