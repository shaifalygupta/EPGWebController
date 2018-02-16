package EPGHandlers

/**
This GO file is used for Unit test cases .
It used the existing Go package : "testing"

 */

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"os"
)

func TestCSVfetchAllData(t *testing.T) {
	c := csvReceiver{}
	EPGDList := c.fetchAllData()
	if len(EPGDList) != 4 {
		t.Errorf("Expected EPG Data length 4, but got %v", len(EPGDList))
	}
	if EPGDList[0].ID != "120011" {
		t.Errorf("Expected EPG Data First value is wrong , which is %v", EPGDList[0].ID)
	}

}

func TestCSVfetchEPGBychannel(t *testing.T) {
	c := csvReceiver{"USA","MTV"}
	EPGDList := c.fetchEPGBychannel()
	//ShowCase of Fail case although we are getting data lengh as 1 which is correct
	if len(EPGDList) != 2 {
		t.Errorf("Expected EPG Data length 1, but got %v", len(EPGDList))
	}
	if EPGDList[0].ID != "120013" {
		t.Errorf("Expected EPG Data First value is wrong , which is %v", EPGDList[0].ID)
	}

}

func TestCreateEPGData(t *testing.T) {
	epData:=EPGData{}
	filename := "testEPGJSONData"
	bs,er:=ioutil.ReadFile(filename)
 if er!=nil{
	 t.Errorf("Error in Reading File")
	 os.Exit(1)
 }
	 if err := json.Unmarshal(bs, &epData);err != nil {
	 	t.Errorf("Error in Data")
		 os.Exit(1)
	 }
	c := mockReceiver{}
	EPGDList := c.CreateEPGData(epData)
	if len(EPGDList) != 1 {
		t.Errorf("Expected EPG Data length 1, but got %v", len(EPGDList))
	}

}
