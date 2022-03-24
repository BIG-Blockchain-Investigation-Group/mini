package minichain

import (
    "gopkg.in/ini.v1"
    "io"
    "context"
    kivik "github.com/go-kivik/kivik/v3"
    _ "github.com/go-kivik/couchdb/v3" 
)

func GetConnection()(*kivik.DB){
    cfg, _ := ini.Load("my.ini")
    connectionURL := cfg.Section("").Key("ConnectionURL").String()
    databaseName := cfg.Section("").Key("DatabaseName").String()
    client, err := kivik.New("couch", connectionURL)
    if err != nil {
        panic(err)
    }
    db := client.DB(context.TODO(), databaseName)
    return db
}

func PutState(key string, value string)(string){
    db := GetConnection()
    rev, err := db.Put(context.TODO(), key, value)
    if err != nil {
        panic(err)
    }
    return rev
}

func GetState(key string)(string){
    db := GetConnection()
    row := db.Get(context.TODO(), key) // Returns the content length and an io.ReadCloser(Body) to access the raw JSON content
    if row.ContentLength <= 0 {
        return "Invalid Key"
    }    
    bytes, err := io.ReadAll(row.Body) // row.Body is a io.ReadCloser which is converted to bytes using io.ReadAll method
    if err != nil {
        panic(err)
    }
    return string(bytes)
}