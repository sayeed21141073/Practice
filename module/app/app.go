package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//create our own multiplexer
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	//now gorilla mux is introduced, don't need any changes

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/Addcustomer", createCustomers).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Fprint(w, vars["customer_id"])
}
func createCustomers(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Post Request Received")
}
