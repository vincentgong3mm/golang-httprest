package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

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

func TestMongoDBConnect(t *testing.T) {
	account := LoadSetting()
	log.Println(account)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://cluster0.izgwf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority").SetAuth(options.Credential{
		AuthSource: "",
		Username:   account.UserName,
		Password:   account.Password,
	})
	defer cancelFunc()

	// mongodb 연결
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)

	}

	// 연결 검증
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ok")
}
