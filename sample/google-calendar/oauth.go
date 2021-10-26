package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientCredential struct {
	Web struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	} `json:"web"`
}

var httpPort = 8080

// Google Cloud Platform > Credentials > OAuth 2.0 Client IDs > Client ID
var clientID = "YOUR_CLIENT_ID"

//Google Cloud Platform > Credentials > OAuth 2.0 Client IDs > Client Secret
var clientSecret = "YOUR_CLIENT_SECRET"

var AccessToken = ""

// Toke은 인증 후 결과를 받은 json에서 access_token, toke_type을 구하기 위한 sturct입니다.
// Step 5: Exchange authorization code for refresh and access tokens 단계 참고
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

var MyToken = Token{}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func DefaultMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Google OAuth"))
}

// RedirectGoogleOAuth는 google oauth로 인증요청합니다.
// 참고 : https://developers.google.com/identity/protocols/oauth2/web-server
//	Step 1: Set authorization parameters
//	Step 2: Redirect to Google's OAuth 2.0 server
//	// Step 3: Google prompts user for consent : 유저 동을 절차
//	// Step 4: Handle the OAuth 2.0 server response
func RedirectGoogleOAuth(w http.ResponseWriter, r *http.Request) {
	log.Println("RedirectGoogleOAuth ------------")

	q := make(map[string]string)

	// Google Cloud Platform에서 생성한 Client ID를 설정해야 합니다.
	//		Google Cloud Platform > Credentials > OAuth 2.0 Client IDs > Client ID
	q["client_id"] = clientID

	// Google Cloud Platform에 설정한 redirection uri를 설정해야 합니다.
	// 		Google Cloud Platform > Credentials > OAuth 2.0 Client IDs > Authorized redirect URIs
	q["redirect_uri"] = fmt.Sprintf("http://localhost:%v/redirectcode", httpPort)

	// 테스트로 google calendar에 access하기 위한 설정입니다.
	q["scope"] = "https://www.googleapis.com/auth/calendar"

	q["state"] = RandToken()
	q["response_type"] = "code"

	query := ""
	for k, v := range q {
		query += fmt.Sprintf("%v=%v&", k, v)
	}

	log.Println(query)

	// redirect 하면서 필요한 data를 추가합니다.
	// Step 2: Redirect to Google's OAuth 2.0 server, After you create the request URL, redirect the user to it.
	http.Redirect(w, r, "https://accounts.google.com/o/oauth2/v2/auth?"+query, http.StatusMovedPermanently) // 301
}

// RedirectCode 유저 동의 후 google server에서 redirect_uri로 redirectt합니다.
// 이때 ReqAccessToken를 호출해서 code로 access token을 구합니다.
// 참고 : https://developers.google.com/identity/protocols/oauth2/web-server
//		Step 5: Exchange authorization code for refresh and access tokens
func RedirectCode(w http.ResponseWriter, r *http.Request) {
	log.Println("RedirectCode ------------")

	log.Println("Query |", r.URL.Query())

	q := r.URL.Query()

	code := q["code"]
	log.Println("q |", q)
	log.Println("len |", len(code), "code |", code)

	if len(code) > 0 {
		t, _ := ReqAccessToken(code[0])

		// 테스트를 위해서 access_token는 보내줍니다.
		b, _ := json.Marshal(&t)
		w.Write(b)

		MyToken = t

		log.Println("MyToken : ", MyToken)

		//ReqCalEvents(t)
	}
}

func ReqAccessToken(code string) (Token, int) {
	// access token을 구하기 위한 uri입니다.
	url := "https://oauth2.googleapis.com/token"

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("client_id", clientID)

	// Google Cloud Platform에서 생성한 Client Secret를 설정해야 합니다.
	//		Google Cloud Platform > Credentials > OAuth 2.0 Client IDs > Client Secret
	q.Add("client_secret", clientSecret)

	// Google Cloud Platform에 설정한 redirection uri를 설정해야 합니다.
	// 		Google Cloud Platform > Credentials > OAuth 2.0 Client IDs > Authorized redirect URIs
	q.Add("redirect_uri", fmt.Sprintf("http://localhost:%v/redirectcode", httpPort))

	q.Add("grant_type", "authorization_code")
	q.Add("code", code)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("StatusCode :", resp.StatusCode)

	t := Token{}
	if resp.StatusCode == http.StatusOK {
		// 결과 출력
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes) //바이트를 문자열로

		log.Println("Resp :", resp)
		log.Println("Resp>Body :", str)

		err := json.Unmarshal([]byte(str), &t)
		if err != nil {
			panic(err)
		}

		log.Println("Token :", t)

		return t, resp.StatusCode
	} else {
		// 결과 출력
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes) //바이트를 문자열로

		log.Println("Resp :", resp)
		log.Println("Resp>Body :", str)
	}

	return t, resp.StatusCode
}

func Serve() {

	// Google Cloud Platform > Credentials에서 다운로드한 파일에서 client_id, client_secret를 읽습니다.
	// client_id, client_secret를 별도의 파일에서 읽고, commit할 때 설정이 있는 파일은 commit하지 않습니다.
	b, err := ioutil.ReadFile("./client_secret/client_secret.json")
	if err != nil {
		log.Fatal(err)
	}

	// byte -> string
	str := string(b)

	var credential ClientCredential
	json.Unmarshal([]byte(str), &credential)

	clientID = credential.Web.ClientId
	clientSecret = credential.Web.ClientSecret

	http.HandleFunc("/", DefaultMain)
	http.HandleFunc("/oauth", RedirectGoogleOAuth)
	http.HandleFunc("/redirectcode", RedirectCode)
	http.HandleFunc("/gcal", ReqCalEvents)

	log.Printf("Sever starting on port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", httpPort), nil))
}
