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
	sl.Info.Println(req.Method)
	sl.Info.Println(req.URL.Query())

	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
		len := req.ContentLength
		body := make([]byte, len)
		req.Body.Read(body)
		sl.Info.Println("Body : ", string(body))
	case http.MethodDelete:
		len := req.ContentLength
		body := make([]byte, len)
		req.Body.Read(body)
		sl.Info.Println("Body : ", string(body))
	}

	//sl.Info.Println(req)

	w.Write([]byte("Hello http api"))
}

func NewServe() {
	http.Handle("/user", new(SearchUser))

	http.ListenAndServe(":9090", nil)
}
