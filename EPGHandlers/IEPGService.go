package EPGHandlers

/**
This structure is used to POST or GET information related to EPG
Json tags are used to map the json property name with strut properties.
[{"id":"120013","name":" Test Series","type":"Program","subtype":"Episode","duration":1800,"sourcesystem":"P33","country":"USA","channel":"MTV"}]
 */
type  EPGData  struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	SubType string `json:"subtype"`
	Duration int `json:"duration"`
	SourceSystem string `json:"sourcesystem"`
	Country string `json:"country"`
	Channel string `json:"channel"`
}

type EPGDataList []EPGData

/**
This interface is used to support the EPG Functionality
EPG only supports the functions mentioned in the Interface
 */
type EPGInterface interface {
	fetchAllData()  EPGDataList
	fetchEPGByCountry()  EPGDataList
	fetchEPGBychannel()  EPGDataList
	CreateEPGData(data EPGData) EPGDataList
}

/**
The mockreceiver implements all the epg interface functions .
The mockreceiver returns the mock values
 */
type mockReceiver struct {
	country string
	channel string
}
/**
The csvreceiver implements all the epg interface functions .
The csvreceiver returns the values based on reading the csv file .
 */
type csvReceiver struct {
	country string
	channel string
}

/**
this function is used call the concrete function(fetchAllData) based on Type i.e csvreceiver or mock receiver
 */
func IfetchAllDataProxy(epgInterface EPGInterface) EPGDataList{
	return  epgInterface.fetchAllData()
}
/**
this function is used call the concrete function (fetchEPGByCountry) based on Type i.e csvreceiver or mock receiver
 */
func IfetchEPGByCountryProxy(epgInterface EPGInterface) EPGDataList{
	return epgInterface.fetchEPGByCountry()
}
/**
this function is used call the concrete function (fetchEPGBychannel) based on Type i.e csvreceiver or mock receiver
 */
func IfetchEPGBychannelProxy(epgInterface EPGInterface) EPGDataList{
	return epgInterface.fetchEPGBychannel()
}
/**
this function is used call the concrete function (CreateEPGData) based on Type i.e csvreceiver or mock receiver
 */
func ICreateEPGDataProxy(epgInterface EPGInterface,data EPGData) EPGDataList{
	return epgInterface.CreateEPGData(data)
}

