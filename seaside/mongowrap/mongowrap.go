package mongowrap

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI      string `json:"uri"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type MongoConn struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	client     *mongo.Client
}

func NewMongoConn() *MongoConn {
	conn := &MongoConn{}

	// context의 timeout 시간이 짧으면 context deadline exceeded 에러 발생합니다.
	conn.ctx, conn.cancelFunc = context.WithTimeout(context.Background(), 30*time.Second)
	conn.client = nil
	return conn
}

func (c *MongoConn) Close() {
	c.cancelFunc()
}

func (c *MongoConn) Connect(config *Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.URI).SetAuth(options.Credential{
		AuthSource: "",
		Username:   config.UserName,
		Password:   config.Password,
	})

	// mongodb 연결
	client, err := mongo.Connect(c.ctx, clientOptions)
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

	log.Println("Connect Ok.", "User:", config.UserName, "URI:", config.URI)

	c.client = client
	return client, nil
}

// GetCollection은 mongdb접속 정보에서 지정한 db의 collection 정보를 구합니다.
func (c *MongoConn) getCollection(db, collection string) *mongo.Collection {
	return c.client.Database(db).Collection(collection)
}

func (c *MongoConn) Find(findCol *FindCollection) error {
	collection := c.getCollection(findCol.DBName, findCol.CollectionName)
	cursor, err := collection.Find(c.ctx, findCol.Query)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err = cursor.All(c.ctx, &findCol.Document); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// LoadSetting은 mongdb를 접속하기 위해한 URI와 계정정보를 파일에서 읽습니다.
func LoadDBSetting(file string) *Config {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// byte -> string
	str := string(b)

	var config Config
	json.Unmarshal([]byte(str), &config)

	return &config
}

type FindCollection struct {
	DBName         string
	CollectionName string
	Query          bson.M
	Document       []bson.M
}
