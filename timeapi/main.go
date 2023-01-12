package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/time", timeNow)
	http.ListenAndServe(":8080", router)
}
func timeNow(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Asia/Dhaka")
	//fmt.Println(time.Now().In(loc))
	fmt.Fprint(w, time.Now().In(loc))

}
