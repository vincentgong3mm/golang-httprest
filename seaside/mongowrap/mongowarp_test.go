package mongowrap

import (
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

var MongDBTestURI = "mongodb+srv://cluster0.izgwf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

func TestConnect(t *testing.T) {
	config := LoadDBSetting("./for_test_setting.json")

	conn := NewMongoConn()
	conn.Connect(config)
}

func TestFindCollection(t *testing.T) {
	config := LoadDBSetting("./for_test_setting.json")

	conn := NewMongoConn()
	conn.Connect(config)

	//findCol := FindCollection{DBName: "seaside_user", CollectionName: "user"}
	//findCol.Query = bson.M{"user_name": "vcassel"}

	findCol := FindCollection{DBName: "seaside_user", CollectionName: "user", Query: bson.M{"user_name": "vcassel"}}
	if err := conn.Find(&findCol); err != nil {
		log.Fatal(err)
		return
	}

	log.Println(findCol)
}
