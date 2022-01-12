package main

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// 참고 사이트 : building-microservices-with-go
// - https://github.com/building-microservices-with-go/chapter4/tree/master/handlers

func setUpAPITest(handler http.Handler, method, target string, d io.Reader) *httptest.ResponseRecorder {

	response := httptest.NewRecorder()

	if d == nil {
		request := httptest.NewRequest(method, target, nil)
		handler.ServeHTTP(response, request)
	} else {
		request := httptest.NewRequest(method, target, d)
		sl.Info.Println("test request")
		sl.Info.Println(request)

		handler.ServeHTTP(response, request)
	}

	return response
}

type SearchUser struct {
}

type ReqQuery struct {
	Query string `json:"query"`
}

func (s *SearchUser) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	sl.Info.Println(req.URL.Path)
	sl.Info.Println(req.URL.Query())
	sl.Info.Println("log req all")

	sl.Info.Println(req)

	w.Write([]byte("Hello http api"))

	len := req.ContentLength
	body := make([]byte, len)
	req.Body.Read(body)
	sl.Info.Println("body : ", string(body))

}

func NewServe() {
	http.Handle("/user", new(SearchUser))

	http.ListenAndServe(":9090", nil)
}
