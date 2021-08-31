package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func HealthReport(studentID string, passwd string) {
	Url := "https://enroll.scut.edu.cn/door/health/h5/get"
	httpClient := Login(Url, studentID, passwd)

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	dec := json.NewDecoder(bytes.NewReader(bodyBytes))
	var jsonDecode map[string]interface{}
	dec.Decode(&jsonDecode)
	if jsonDecode["code"].(float64) != 1 {
		fmt.Println("panic code!=1")
		return
	}
	jsonDecode = jsonDecode["data"].(map[string]interface{})

	//填充表单信息，筛选出get中有用的信息填入
	jsonDecode = jsonDecode["healthRptInfor"].(map[string]interface{})
	PostDict := url.Values{}
	PostDict.Set("sPersonName", fmt.Sprint(jsonDecode["sPersonName"]))
	PostDict.Set("sPersonCode", fmt.Sprint(jsonDecode["sPersonCode"]))
	PostDict.Set("dRptDate", fmt.Sprint(jsonDecode["dRptDate"]))
	PostDict.Set("sPhone", fmt.Sprint(jsonDecode["sPhone"]))
	PostDict.Set("sParentPhone", fmt.Sprint(jsonDecode["sParentPhone"]))
	PostDict.Set("iIsGangAoTai", fmt.Sprint(jsonDecode["iIsGangAoTai"]))
	PostDict.Set("iIsOversea", fmt.Sprint(jsonDecode["iIsOversea"]))
	PostDict.Set("sHomeProvName", fmt.Sprint(jsonDecode["sHomeProvName"]))
	PostDict.Set("sHomeProvCode", fmt.Sprint(jsonDecode["sHomeProvCode"]))
	PostDict.Set("sHomeCityName", fmt.Sprint(jsonDecode["sHomeCityName"]))
	PostDict.Set("sHomeCityCode", fmt.Sprint(jsonDecode["sHomeCityCode"]))
	PostDict.Set("sHomeCountyName", fmt.Sprint(jsonDecode["sHomeCountyName"]))
	PostDict.Set("sHomeCountyCode", fmt.Sprint(jsonDecode["sHomeCountyCode"]))
	PostDict.Set("sHomeAddr", fmt.Sprint(jsonDecode["sHomeAddr"]))
	PostDict.Set("iSelfState", fmt.Sprint(jsonDecode["iSelfState"]))
	PostDict.Set("iFamilyState", fmt.Sprint(jsonDecode["iFamilyState"]))
	PostDict.Set("sNowProvName", fmt.Sprint(jsonDecode["sNowProvName"]))
	PostDict.Set("sNowProvCode", fmt.Sprint(jsonDecode["sNowProvCode"]))
	PostDict.Set("sNowCityName", fmt.Sprint(jsonDecode["sNowCityName"]))
	PostDict.Set("sNowCityCode", fmt.Sprint(jsonDecode["sNowCityCode"]))
	PostDict.Set("sNowCountyName", fmt.Sprint(jsonDecode["sNowCountyName"]))
	PostDict.Set("sNowCountyCode", fmt.Sprint(jsonDecode["sNowCountyCode"]))
	PostDict.Set("sNowAddr", fmt.Sprint(jsonDecode["sNowAddr"]))
	PostDict.Set("iNowGoRisks", fmt.Sprint(jsonDecode["iNowGoRisks"]))
	PostDict.Set("iRctRisks", fmt.Sprint(jsonDecode["iRctRisks"]))
	PostDict.Set("iRctKey", fmt.Sprint(jsonDecode["iRctKey"]))
	PostDict.Set("iRctOut", fmt.Sprint(jsonDecode["iRctOut"]))
	PostDict.Set("iRctTouchKeyMan", fmt.Sprint(jsonDecode["iRctTouchKeyMan"]))
	PostDict.Set("iRctTouchBackMan", fmt.Sprint(jsonDecode["iRctTouchBackMan"]))
	PostDict.Set("iRctTouchDoubtMan", fmt.Sprint(jsonDecode["iRctTouchDoubtMan"]))
	PostDict.Set("iVaccinState", fmt.Sprint(jsonDecode["iVaccinState"]))
	PostDict.Set("iHealthCodeState", fmt.Sprint(jsonDecode["iHealthCodeState"]))
	PostDict.Set("iRptState", fmt.Sprint(jsonDecode["iRptState"]))
	PostDict.Set("sDegreeCode", fmt.Sprint(jsonDecode["sDegreeCode"]))
	PostDict.Set("iSex", fmt.Sprint(jsonDecode["iSex"]))
	PostDict.Set("sCollegeName", fmt.Sprint(jsonDecode["sCollegeName"]))
	PostDict.Set("sCampusName", fmt.Sprint(jsonDecode["sCampusName"]))
	PostDict.Set("sDormBuild", fmt.Sprint(jsonDecode["sDormBuild"]))
	PostDict.Set("sDormRoom", fmt.Sprint(jsonDecode["sDormRoom"]))
	PostDict.Set("sMajorName", fmt.Sprint(jsonDecode["sMajorName"]))
	PostDict.Set("sClassName", fmt.Sprint(jsonDecode["sClassName"]))
	PostDict.Set("iInSchool", fmt.Sprint(jsonDecode["iInSchool"]))

	req, err = http.NewRequest("POST", "https://enroll.scut.edu.cn/door/health/h5/add", bytes.NewReader([]byte(PostDict.Encode())))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err = httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes), time.Now(), studentID)
}
