package robotserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range params {
		fmt.Fprintf(w, "(GetData) Key: %v, value: %v", i, item)
	}
}

func Change(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range params {
		fmt.Fprintf(w, "(GetData) Key: %v, value: %v", i, item)
	}
}

func RobotServerMain() {
	router := mux.NewRouter()
	router.HandleFunc("/data/{req}", GetData).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8800", router))
}
