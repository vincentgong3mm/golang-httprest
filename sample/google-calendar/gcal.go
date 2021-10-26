package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReqCalEvents(w http.ResponseWriter, r *http.Request) {
	id := "vincentgong3mm@gmail.com"
	url := "https://www.googleapis.com/calendar/v3/calendars/" + id + "/events"

	req, _ := http.NewRequest("GET", url, nil)

	token := "Bearer "
	token += MyToken.AccessToken
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("status : ", resp.StatusCode)

	if resp.StatusCode != 200 {
		// 결과 출력
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes) //바이트를 문자열로

		fmt.Println("ReqCallEvents get token ", resp)
		fmt.Println("ReqCallEvents get otken : req body : ", str)

		w.Write(bytes)

	} else {
		// 결과 출력
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes) //바이트를 문자열로

		fmt.Println("ReqCallEvents get token ", resp)
		fmt.Println("ReqCallEvents get otken : req body : ", str)

		w.Write(bytes)
	}
}
