package models


type Contact struct {
	Id string `json:"id"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone"`
}

var Contacts = []Contact{
	{"1", "John Doe", "0735667891"},
	{"2","Constantin Mihaita", "0723554012"},
	{"3", "Farcas Mihai", "0744236035"},
}