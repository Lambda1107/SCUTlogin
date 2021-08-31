package server

import (
	"SCUTlogin/utils"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"

	"github.com/robertkrimen/otto"
)

func Login(serviceUrl string, studentID string, passwd string) *http.Client {
	curCookieJar, _ := cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: curCookieJar,
	}
	req, err := http.NewRequest("GET", serviceUrl, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	lt := utils.March(string(bodyBytes), "<input type=\"hidden\" id=\"lt\" name=\"lt\" value=\"", "\"")
	jsfile := "des.js"
	jsbytes, _ := ioutil.ReadFile(jsfile)
	vm := otto.New()
	//call js
	vm.Run(string(jsbytes))
	rsa, _ := vm.Call("strEnc", nil, studentID+passwd+lt, "1", "2", "3")

	postDict := map[string]string{}
	postDict["rsa"] = rsa.String()
	postDict["ul"] = strconv.Itoa(len(studentID))
	postDict["pl"] = strconv.Itoa(len(passwd))
	postDict["lt"] = lt
	postDict["execution"] = "e1s1"
	postDict["_eventId"] = "submit"

	postValues := url.Values{}

	for postKey, PostValue := range postDict {
		postValues.Set(postKey, PostValue)
	}

	req, err = http.NewRequest("POST", "https://sso.scut.edu.cn/cas/login?service="+url.PathEscape(serviceUrl), bytes.NewReader([]byte(postValues.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	if err != nil {
		log.Println(err)
	}
	_, err = httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	//登录完了，为所欲为
	return httpClient
}
