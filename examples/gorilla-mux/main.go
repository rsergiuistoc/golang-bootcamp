package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rsergiuistoc/golang-bootcamp/models"
)

var contacts = models.Contacts

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/contacts", getAllContacts).Methods("GET")
	router.HandleFunc("/contacts", createContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", getContact).Methods("GET")
	router.HandleFunc("/contacts/{id}", updateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", deleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAllContacts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, c := range contacts{
		if c.Id == params["id"]{
			_ = json.NewEncoder(w).Encode(c)
			break
		}
	}


}

func createContact(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	var contact models.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	_ = json.NewEncoder(w).Encode(&contact)

}

func updateContact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, c := range contacts{
		if c.Id == params["id"]{
			contacts = append(contacts[:i], contacts[i+1:]...)

			var contact models.Contact
			_ = json.NewDecoder(r.Body).Decode(contact)
			contact.Id = params["id"]
			contacts = append(contacts, contact)
			_ = json.NewEncoder(w).Encode(&contact)
			return
		}
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, c := range contacts {
		if c.Id == params["id"] {
			contacts = append(contacts[:i], contacts[i+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(contacts)

}