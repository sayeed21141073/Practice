package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"customer_name" xml:"customer_name"`
	City    string `json:"city/place" xml:"city/place"`
	ZipCode int    `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Sayeed", City: "Dhaka", ZipCode: 1206},
		{Name: "Ibne", City: "Rajshahi", ZipCode: 1000},
		{Name: "Zaman", City: "Comilla", ZipCode: 106},
	}

	//w.Header().Add("Content-Type", "application/xml")
	//now based on my Content-Type, I want xml/json based on content-type

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		json.NewEncoder(w).Encode(customers)
	}
}
