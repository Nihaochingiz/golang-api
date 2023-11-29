package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type advertisement struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Photo       string `json:"Photo"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
	CreatedAt   string `json:"CreatedAt"`
}

type allAdvertisements []advertisement

var Advertisements = allAdvertisements{
	{
		ID:          "1",
		Title:       "Nike advertising",
		Photo:       "https://www.marketing91.com/wp-content/uploads/2020/06/Nike-Advertising-Example-Just-Do-It-Campaign.jpg",
		Description: "JUST DO IT ",
		Price:       1000,
		CreatedAt:   "2023/05/03",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to advertisement DB!")
}

func createAdvertisement(w http.ResponseWriter, r *http.Request) {
	var newAdvertisement advertisement
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Advertisement title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newAdvertisement)
	Advertisements = append(Advertisements, newAdvertisement)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newAdvertisement)
}

func getOneAdvertisement(w http.ResponseWriter, r *http.Request) {
	AdvertisementID := mux.Vars(r)["id"]

	for _, singleAdvertisement := range Advertisements {
		if singleAdvertisement.ID == AdvertisementID {
			json.NewEncoder(w).Encode(singleAdvertisement)
		}
	}
}

func getAllAdvertisements(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Advertisements)
}

func updateAdvertisement(w http.ResponseWriter, r *http.Request) {
	AdvertisementID := mux.Vars(r)["id"]
	var updatedAdvertisement advertisement

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Advertisement title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedAdvertisement)

	for i, singleAdvertisement := range Advertisements {
		if singleAdvertisement.ID == AdvertisementID {
			singleAdvertisement.Title = updatedAdvertisement.Title
			singleAdvertisement.Description = updatedAdvertisement.Description
			Advertisements = append(Advertisements[:i], singleAdvertisement)
			json.NewEncoder(w).Encode(singleAdvertisement)
		}
	}
}

func deleteAdvertisement(w http.ResponseWriter, r *http.Request) {
	AdvertisementID := mux.Vars(r)["id"]

	for i, singleAdvertisement := range Advertisements {
		if singleAdvertisement.ID == AdvertisementID {
			Advertisements = append(Advertisements[:i], Advertisements[i+1:]...)
			fmt.Fprintf(w, "The Advertisement with ID %v has been deleted successfully", AdvertisementID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/advertisement", createAdvertisement).Methods("POST")
	router.HandleFunc("/advertisements", getAllAdvertisements).Methods("GET")
	router.HandleFunc("/advertisements/{id}", getOneAdvertisement).Methods("GET")
	router.HandleFunc("/advertisements/{id}", updateAdvertisement).Methods("PATCH")
	router.HandleFunc("/advertisements/{id}", deleteAdvertisement).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
