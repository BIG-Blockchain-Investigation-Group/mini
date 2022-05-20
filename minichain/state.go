package minichain

import (
	"context"
	"io"

	_ "github.com/go-kivik/couchdb/v3"
	kivik "github.com/go-kivik/kivik/v3"
	"gopkg.in/ini.v1"
)

func GetConnection() (*kivik.Client, *kivik.DB) {
	cfg, _ := ini.Load("my.ini")
	connectionURL := cfg.Section("").Key("ConnectionURL").String()
	databaseName := cfg.Section("").Key("DatabaseName").String()
	client, err := kivik.New("couch", connectionURL)
	if err != nil {
		panic(err)
	}
	db := client.DB(context.TODO(), databaseName)
	return client, db
}

var client, db = GetConnection()

func PutState(key string, value []byte) string {
	rev, err := db.Put(context.TODO(), key, string(value))
	if err != nil {
		panic(err)
	}
	return rev
}

func GetState(key string) []byte {
	row := db.Get(context.TODO(), key) // Returns the content length and an io.ReadCloser(Body) to access the raw JSON content
	bytes, err := io.ReadAll(row.Body) // row.Body is a io.ReadCloser which is converted to bytes using io.ReadAll method
	if err != nil {
		panic(err)
	}
	return bytes
}
