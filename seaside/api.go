package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	_ "github.com/vincentgong3mm/golanghttprest/seaside/docs"
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


type User struct {
}

type GetUser struct {
}

type SearchUser struct {
}

type ReqQuery struct {
	Query string `json:"query"`
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 {object} User
// @Router /user/{name} [get]
func (g *GetUser) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("GetUser"))
}

// @Summary Search User
// @Description Search User's Info
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 {object} User
// @Router /user2/{name} [get]
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
	sl.Info.Println("start NewServe" + ":8080")

	// 아래와 같이 해야지 /만 있는 경로가 아닌 /docs/ 경로설정 가능합니다.
	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs", fs))

	http.Handle("/user", new(GetUser))
	http.Handle("/user2", new(SearchUser))

	http.ListenAndServe(":8080", nil)
}
