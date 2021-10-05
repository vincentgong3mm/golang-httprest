package main

import (
	"context"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func TestMongDBInsert(t *testing.T) {
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

	log.Println("connect ok")

	collection := GetCollection(client, "test_users")

	type User struct {
		Name  string `json:"name"`
		EMail string `json:"email"`
	}

	insertRet, err2 := collection.InsertOne(ctx, User{Name: "vincent2", EMail: "a@a.com"})
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(insertRet.InsertedID)
}
