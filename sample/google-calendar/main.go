package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
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

// GetCollection은 mongdb접속 정보에서 지정한 collection 정보를 구합니다. db이름은 만들어둔 gcalender로 임시로 설정했습니다.
func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
	dbName := "gcalendar"
	return client.Database(dbName).Collection(colName)
}

func main() {
	account := LoadSetting()
	log.Println(account)

}
