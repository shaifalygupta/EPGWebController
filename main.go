package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
  m	"EPGWebController/EPGHandlers"
)

/**
Main Router file to handle all the request based on URI .
This Supports GET & POST Http Method
 */
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/EPG/{receive}",m.FetchEPGData).Methods("GET")
	router.HandleFunc("/EPG/{receive}/{country}", m.FetchEPGByCountry).Methods("GET")
	router.HandleFunc("/EPG/{receive}/{country}/{channel}", m.FetchEPGBychannel).Methods("GET")
	router.HandleFunc("/EPG/{receive}", m.CreateEPGData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}