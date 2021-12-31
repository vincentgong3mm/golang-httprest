package mongowrap

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDBCloudDBURI = "mongodb+srv://cluster0.izgwf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

// Acccount는 mongdb를 접속하기 위해 setting.json파일에서 읽을 정보 구조체입니다.
type Account struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// LoadSetting은 mongdb를 접속하기 위해한 계정정보를 setting.json파일에서 읽습니다.
func loadSetting() Account {
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

func Connect() (*mongo.Client, error) {
	acc := loadSetting()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(MongoDBCloudDBURI).SetAuth(options.Credential{
		AuthSource: "",
		Username:   acc.UserName,
		Password:   acc.Password,
	})
	defer cancelFunc()

	// mongodb 연결
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 연결 검증
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connect Ok.", "User:", acc.UserName, "URI:", MongoDBCloudDBURI)
	return client, nil
}

// GetCollection은 mongdb접속 정보에서 지정한 db의 collection 정보를 구합니다.
func GetCollection(client *mongo.Client, db, collection string) *mongo.Collection {
	return client.Database(db).Collection(collection)
}
