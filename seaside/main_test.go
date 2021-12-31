package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/vincentgong3mm/golanghttprest/seaside/mongowrap"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMain(t *testing.T) {
	fmt.Println("test hello world")

	config := mongowrap.LoadDBSetting("./db_setting.json")
	conn := mongowrap.NewMongoConn()
	conn.Connect(config)

	findCol := mongowrap.FindCollection{DBName: "seaside_user", CollectionName: "user", Query: bson.M{"user_name": "mbell"}}
	if err := conn.Find(&findCol); err != nil {
		log.Fatal(err)
		return
	}

	log.Println(findCol)
}

func TestLogPrint(t *testing.T) {
	NewSlog()
	sl.Info.Println("info log.")
	sl.Error.Println("error log.")
}
