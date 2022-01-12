package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/vincentgong3mm/golanghttprest/seaside/mongowrap"
	"go.mongodb.org/mongo-driver/bson"
)

func TestHTTPRequest() {
	NewServe()
}

func TestMongoDB() {
	config := mongowrap.LoadDBSetting("./db_setting.json")
	conn := mongowrap.NewMongoConn()
	conn.Connect(config)

	findCol := mongowrap.FindCollection{DBName: "seaside_user", CollectionName: "user", Query: bson.M{"user_name": "mbell"}}
	if err := conn.Find(&findCol); err != nil {
		log.Fatal("aaa", err)
		return
	}

	log.Println(findCol)
}

func main() {
	NewSlog()
	sl.Info.Println("Hello World!")

	//NewServe()

	//TestMongoDB()

	handler := &SearchUser{}

	// response := setUpAPITest(handler, "GET", "/user", nil)
	// sl.Info.Println(response)

	param := url.Values{}
	param.Add("mykey", "myvalue")
	param.Add("mykey2222", "myvalue2222")

	response2 := setUpAPITest(handler, "POST", "/user", strings.NewReader("test body string..."))
	sl.Info.Println(response2)

}
