package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
  m	"EPGWebController/EPGHandlers"
)

/**
Main Router file to handle all the request based on URI .
This file used gorilla/mux package for parameter substitution in the URL .
This Supports GET & POST Http Method .
Sample Urls for Reference
http://localhost:8081/EPG/csv
http://localhost:8081/EPG/csv/USA
http://localhost:8081/EPG/csv/USA/MTV

POST : http://localhost:8081/EPG/csv
{"ID": "120015","Name":" Test Series","Type":"Program","SubType":"Episode","Duration":1800,"SourceSystem":"P33","Country":"USA","Channel":"MTV"}

http://localhost:8081/EPG/mock
http://localhost:8081/EPG/mock/USA
http://localhost:8081/EPG/mock/USA/MTV

POST : http://localhost:8081/EPG/mock
{"ID": "120015","Name":" Test Series","Type":"Program","SubType":"Episode","Duration":1800,"SourceSystem":"P33","Country":"USA","Channel":"MTV"}
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
