// json encoding means making anything to json

package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"customer_name"`
	City    string `json:"city/place"`
	ZipCode int    `json:"zip_code"`
}

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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// soja kotha customers ke json korlam, then w te nilam karon w http server
//er response or output
/* the json.NewEncoder function is used to create a new JSON encoder that writes
to the http.ResponseWriter interface.
The Encode method of the encoder is then called with the customers slice as an argument, which encodes the slice as a JSON array
and writes it to the response writer. */
//json e output asole eta plain text. tai 37 no line likhe json type banano hoyeche
// json tag deoa hoy key ta rename korar jonno, value eki thake.
