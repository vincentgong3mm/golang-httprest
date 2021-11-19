package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReqCalEvents는 지정한 calendar id에 해당되는 events(상세일정)을 가져옵니다.
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

// ReqCallList는 유저에게 등록되어 있는 모든 calendar의 list를 가져옵니다.
// 결과 : dump_callist.json
func ReqCalList(w http.ResponseWriter, r *http.Request) {
	//id := "vincentgong3mm@gmail.com"
	url := "https://www.googleapis.com/calendar/v3/users/me/calendarList"

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

// ReqCalListID는 ReqCals에서 get하는 것과 같은 결과를 주는 군요.
func ReqCalListID(w http.ResponseWriter, r *http.Request) {
	id := "vincentgong3mm@gmail.com"
	url := "https://www.googleapis.com/calendar/v3/users/me/calendarList/" + id

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

// ReqCals는 id에 해당하는 calendar의 정보를 가져옵니다.
// 결과 : dump_cals.json
func ReqCals(w http.ResponseWriter, r *http.Request) {
	id := "vincentgong3mm@gmail.com"
	url := "https://www.googleapis.com/calendar/v3/calendars/" + id

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
