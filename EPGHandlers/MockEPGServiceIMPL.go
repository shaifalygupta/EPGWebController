package EPGHandlers

/***
This Go file is the concrete implementation of EPGInteface functions
This implementation is used to return mock data.
 */

func (m mockReceiver) fetchAllData()  EPGDataList{

	EP := EPGDataList{
		EPGData{ "1234", "Test Asset","Program","Series",1800,"P123","AUS","MTV"},
		EPGData{ "4567", "Test Asset1","Program","Episode",1800,"P123","USA","Star"},
	}

	return EP
}



func (m mockReceiver) fetchEPGByCountry()  EPGDataList{

	EP := EPGDataList{
		EPGData{ "1234", "Test Asset","Program","Series",1800,"P123","AUS","MTV"},
		EPGData{ "4567", "Test Asset1","Program","Episode",1800,"P123","USA","Star"},
	}

	return EP
}


func (m mockReceiver) fetchEPGBychannel()  EPGDataList{

	EP := EPGDataList{
		EPGData{ "1234", "Test Asset","Program","Series",1800,"P123","AUS","MTV"},
		EPGData{ "4567", "Test Asset1","Program","Episode",1800,"P123","USA","Star"},
	}

	return EP
}


func (m mockReceiver) CreateEPGData(data EPGData) EPGDataList{
  EPList:=EPGDataList{}
	EPList = append(EPList, data)
	return EPList


	return EPList
}
