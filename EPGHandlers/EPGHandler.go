package EPGHandlers

/***
This Go file  is used to provide the handler functions to process the API request .
This GO file also prepare json response based on EPGData returned from Backend i.e. Mock , CSV or other implementation.
 */

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
)

func FetchEPGData(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	epglist:=EPGDataList{}
	if vars["receive"]=="csv" {
		s:=csvReceiver{}
		epglist = IfetchAllDataProxy(s);

	}else if vars["receive"]=="mock" {
		s:=mockReceiver{}
		epglist = IfetchAllDataProxy(s);

	}else{
		fmt.Fprintln(w,"Format does not supported")}

	if len(epglist)==0{
		fmt.Fprint(w,"No Data Found!")
	}else {
		json.NewEncoder(w).Encode(epglist)
	}


}

func FetchEPGByCountry(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	epglist:=EPGDataList{}
	if vars["receive"]=="csv" {
		s:=csvReceiver{vars["country"],""}
		epglist = IfetchEPGByCountryProxy(s);

	}else if vars["receive"]=="mock" {
		s:=mockReceiver{vars["country"],""}
		epglist = IfetchEPGByCountryProxy(s);

	}else{
		fmt.Fprint(w,"Format does not supported")}

	if len(epglist)==0{
		fmt.Fprint(w,"No Data Found!")
	}else {
		json.NewEncoder(w).Encode(epglist)
	}

}

func FetchEPGBychannel(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	epglist:=EPGDataList{}
	if vars["receive"]=="csv" {
		s:=csvReceiver{vars["country"],vars["channel"]}
		epglist = IfetchEPGBychannelProxy(s);

	}else if vars["receive"]=="mock"{
		s := mockReceiver{vars["country"], vars["channel"]}
		epglist = IfetchEPGBychannelProxy(s);

	}else {
		fmt.Fprint(w,"Format does not supported")
	}

	if len(epglist)==0{
		fmt.Fprint(w,"No Data Found!")
	}else{
		json.NewEncoder(w).Encode(epglist)
	}


}

func CreateEPGData(w http.ResponseWriter, r *http.Request){

	var EPData EPGData
	vars := mux.Vars(r)
	//body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	body, err := ioutil.ReadAll(io.MultiReader(r.Body))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &EPData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	epglist:=EPGDataList{}
	if vars["receive"]=="csv" {
		s:=csvReceiver{}
		epglist = ICreateEPGDataProxy(s,EPData);

	}else if vars["receive"]=="mock"{
		s:=mockReceiver{}
		epglist = ICreateEPGDataProxy(s,EPData);
	}else {
		fmt.Fprint(w,"Format does not supported")
	}
	//fmt.Println(epglist)
	//json.NewEncoder(w).Encode(epglist)
	//t := ICreateEPGData(CreateEPGData)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(epglist); err != nil {
		panic(err)
	}
}
