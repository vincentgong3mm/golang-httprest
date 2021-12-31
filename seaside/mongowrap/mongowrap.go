package mongowrap

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Acccount는 mongdb를 접속하기 위해 setting.json파일에서 읽을 정보 구조체입니다.
type Account struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// LoadSetting은 mongdb를 접속하기 위해한 계정정보를 setting.json파일에서 읽습니다.
func LoadSetting() Account {
	b, err := ioutil.ReadFile("./setting.json")
	if err != nil {
		log.Fatal(err)
	}

	// byte -> string
	str := string(b)

	var account Account
	json.Unmarshal([]byte(str), &account)

	return account
}
