package EPGHandlers

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

type EPGInterface interface {
	fetchAllData()  EPGDataList
	fetchEPGByCountry()  EPGDataList
	fetchEPGBychannel()  EPGDataList
	CreateEPGData(data EPGData) EPGDataList
}

type mockReceiver struct {
	country string
	channel string
}

type csvReceiver struct {
	country string
	channel string
}


func IfetchAllDataProxy(epgInterface EPGInterface) EPGDataList{
	return  epgInterface.fetchAllData()
}

func IfetchEPGByCountryProxy(epgInterface EPGInterface) EPGDataList{
	return epgInterface.fetchEPGByCountry()

}
func IfetchEPGBychannelProxy(epgInterface EPGInterface) EPGDataList{
	return epgInterface.fetchEPGBychannel()
}

func ICreateEPGDataProxy(epgInterface EPGInterface,data EPGData) EPGDataList{
	return epgInterface.CreateEPGData(data)
}

