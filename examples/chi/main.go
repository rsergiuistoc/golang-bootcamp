package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rsergiuistoc/golang-bootcamp/models"
)

var contacts = models.Contacts

func main(){

	r := chi.NewRouter()
	r.Get("/contacts", GetAllContacts)
	r.Get("/contacts/{id}", GetContact)
	r.Post("/contacts", CreateContact)
	r.Put("/contacts/{id}", UpdateContact)
	r.Delete("/contacts/{id}", DeleteContact)

	_ = http.ListenAndServe(":3000", r)
}

func GetAllContacts(w http.ResponseWriter, r * http.Request) {
	_ = json.NewEncoder(w).Encode(contacts)
}

func GetContact (w http.ResponseWriter, r * http.Request){
	id := chi.URLParam(r, "id")
	if id == ""{
		http.Error(w, http.StatusText(404), 404)
		return
	}
	for _, c := range contacts{
		if c.Id == id {
			_ = json.NewEncoder(w).Encode(c)
		}
	}

}

func CreateContact(w http.ResponseWriter, r * http.Request){

	var contact models.Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)

	contacts = append(contacts, contact)

	_, _ = w.Write([]byte("Contact created successfully"))
	w.WriteHeader(201)


}

func UpdateContact(w http.ResponseWriter, r * http.Request){
	id := chi.URLParam(r, "id")
	if id == ""{
		http.Error(w, http.StatusText(404), 404)
		return
	}
	var contact models.Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)

	for _, c := range contacts{
		if c.Id == id {
			c = contact
		}
	}
	_, _ = w.Write([]byte("Contact updated successfully"))
	w.WriteHeader(201)

}

func DeleteContact(w http.ResponseWriter, r * http.Request){
	id := chi.URLParam(r, "id")
	if id == ""{
		http.Error(w, http.StatusText(404), 404)
		return
	}

	for i, c := range contacts{
		if c.Id == id {
			_ = json.NewEncoder(w).Encode(c)

			copy(contacts[i:], contacts[i+1:])
			contacts[len(contacts) - 1] = models.Contact{}
		}
	}
	_, _ = w.Write([]byte("Contact deleted successfully"))
	w.WriteHeader(201)
}