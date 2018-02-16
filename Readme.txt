package EPGHandlers
    import "EPGWebController\EPGHandlers"

Main.go

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
 
 func handleRequests()
 
 


FUNCTIONS in EPGHandler.go

/***
It is used to provide the handler functions to process the API request .
It also prepare json response based on EPGData returned from Backend i.e. Mock , CSV or other implementation.
 */

func CreateEPGData(w http.ResponseWriter, r *http.Request)

func FetchEPGByCountry(w http.ResponseWriter, r *http.Request)

func FetchEPGBychannel(w http.ResponseWriter, r *http.Request)

func FetchEPGData(w http.ResponseWriter, r *http.Request)

Functions , Types in IEPGService.go

TYPES
/**
 This structure is used to POST or GET information related to EPG Json
    tags are used to map the json property name with strut properties.
    [{"id":"120013","name":" Test
    Series","type":"Program","subtype":"Episode","duration":1800,"sourcesystem":"P33","country":"USA","channel":"MTV"}]
**/
type EPGData struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Type         string `json:"type"`
    SubType      string `json:"subtype"`
    Duration     int    `json:"duration"`
    SourceSystem string `json:"sourcesystem"`
    Country      string `json:"country"`
    Channel      string `json:"channel"`
}

type EPGDataList []EPGData

/**
 This interface is used to support the EPG Functionality EPG only
  supports the functions mentioned in the Interface
    */
type EPGInterface interface {
	fetchAllData()  EPGDataList
	fetchEPGByCountry()  EPGDataList
	fetchEPGBychannel()  EPGDataList
	CreateEPGData(data EPGData) EPGDataList
}

/**
this function is used call the concrete function (CreateEPGData) based
    on Type i.e csvreceiver or mock receiver
    **/
func ICreateEPGDataProxy(epgInterface EPGInterface, data EPGData) EPGDataList

/**
this function is used call the concrete function(fetchAllData) based
    on Type i.e csvreceiver or mock receiver

**/
func IfetchAllDataProxy(epgInterface EPGInterface) EPGDataList

/**
this function is used call the concrete function (fetchEPGByCountry)
    based on Type i.e csvreceiver or mock receiver
* /
func IfetchEPGByCountryProxy(epgInterface EPGInterface) EPGDataList
 
 /**
   this function is used call the concrete function (fetchEPGBychannel)
    based on Type i.e csvreceiver or mock receiver
  */  

func IfetchEPGBychannelProxy(epgInterface EPGInterface) EPGDataList
   

Implementaion of Interfaces 

1. CSVEPGServiceIMPL.go : This Go file is the concrete implementation of EPGInteface functions
			  This implementation is used to read data from csv file.
 
2. MockEPGServiceIMPL.go : This Go file is the concrete implementation of EPGInteface functions
                           This implementation is used to return mock data .
 

Unit Test Sample 

1. PGServiceIMPL_test.go : This GO file is used for Unit test cases .
                           It used the existing Go package : "testing"


