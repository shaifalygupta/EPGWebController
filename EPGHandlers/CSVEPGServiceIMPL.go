package EPGHandlers

/***
This Go file is the concrete implementation of EPGInteface functions
This implementation is used to read data from csv file.
 */

import (
	"os"
	"encoding/csv"
	"strconv"
	"path/filepath"
	"runtime"
	"path"
)

// Fetch all the data for EPG
func (c csvReceiver) fetchAllData()  EPGDataList{

	_, filename, _, _ := runtime.Caller(1)
	EPList := EPGDataList{}
	EPData:=EPGData{}

	if len(filename)>15 {
		filename = path.Join(path.Dir(filename), "EPG.csv")
	}else {
		filename,_ = filepath.Abs("../EPGWebController/EPGHandlers/EPG.csv")
	}
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

    //var n_value int
	// Loop through lines & turn into object
	for i, line := range lines {
		if i > 0 {
			n_value, _ := strconv.Atoi(line[4]);
			EPData = EPGData{line[0], line[1], line[2], line[3], n_value, line[5], line[6], line[7],}
			EPList=append(EPList,EPData)
		}
	}

	return EPList
}

// Fetch all the data for EPG based on country
func (c csvReceiver) fetchEPGByCountry()  EPGDataList{

	_, filename, _, _ := runtime.Caller(1)
	EPList := EPGDataList{}
	EPData:=EPGData{}

	if len(filename)>15 {
		filename = path.Join(path.Dir(filename), "EPG.csv")
	}else {
		filename,_ = filepath.Abs("../EPGWebController/EPGHandlers/EPG.csv")
	}
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	//var n_value int
	// Loop through lines & turn into object
	for i, line := range lines {
		if i > 0 {
			n_value, _ := strconv.Atoi(line[4]);
			if c.country==line[6] {
				EPData = EPGData{line[0], line[1], line[2], line[3], n_value, line[5], line[6], line[7],}
				EPList = append(EPList, EPData)
			}
		}
	}

	return EPList
}

// Fetch all the data for EPG based on Channel & Channel
func (c csvReceiver) fetchEPGBychannel()  EPGDataList{

	_, filename, _, _ := runtime.Caller(1)
	EPList := EPGDataList{}
	EPData:=EPGData{}

	if len(filename)>15 {
		filename = path.Join(path.Dir(filename), "EPG.csv")
	}else {
		filename,_ = filepath.Abs("../EPGWebController/EPGHandlers/EPG.csv")
	}
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	//var n_value int
	// Loop through lines & turn into object
	for i, line := range lines {
		if i > 0 {
			n_value, _ := strconv.Atoi(line[4]);
			if c.country==line[6] && c.channel==line[7] {
				EPData = EPGData{line[0], line[1], line[2], line[3], n_value, line[5], line[6], line[7],}
				EPList = append(EPList, EPData)
			}
		}
	}

	return EPList
}

// This is used to post the data
func (m csvReceiver) CreateEPGData(data EPGData) EPGDataList{
	EPList:=m.fetchAllData()
	EPList = append(EPList, data)
	return EPList


	return EPList
}
