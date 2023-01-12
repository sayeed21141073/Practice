//do XML marshalling
//how to request the server to send us data in the JSON or XML format

package main

import (
	"encoding/xml"
	"net/http"
)

// this struct can handle both json and xml
type Customer struct {
	Name    string `json:"customer_name" xml:"customer_name"`
	City    string `json:"city/place" xml:"city/place"`
	ZipCode int    `json:"zip_code" xml:"zip_code"`
}

//only for xml, the srtuct should be this
// type Customer struct {
// 	Name    string `xml:"customer_name"`
// 	City    string `xml:"city/place"`
// 	ZipCode int    `xml:"zip_code"`
// }

func main() {

	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe(":8080", nil)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Sayeed", City: "Dhaka", ZipCode: 1206},
		{Name: "Ibne", City: "Rajshahi", ZipCode: 1000},
		{Name: "Zaman", City: "Comilla", ZipCode: 106},
	}

	w.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(customers)
	// like json.NewEncoder(w).Encode(customers)

	//now based on my Content-Type, I want xml/json based on content-type

	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	json.NewEncoder(w).Encode(customers)
	// }
	//so based on my content-type i can get json/xml
	//that's how we can tell the server how to send data in our appropriate format
	//using request header
}
