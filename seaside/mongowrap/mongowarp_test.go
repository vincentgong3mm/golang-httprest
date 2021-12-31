package mongowrap

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestLoadConfig(t *testing.T) {
	acc := loadSetting()

	fmt.Println(acc)
}

func TestConnect(t *testing.T) {
	log.Println(Connect())
}

func TestSelectCollection(t *testing.T) {
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
		return
	}

	// 연결 검증
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Connect Ok.", "User:", acc.UserName, "URI:", MongoDBCloudDBURI)

	userCollection := GetCollection(client, "seaside_user", "user")
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

}
