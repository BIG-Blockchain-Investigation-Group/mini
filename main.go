package main

import (
    "gopkg.in/ini.v1"
    "io/ioutil"
    "fmt"
    "mini/minichain" 
    _ "github.com/go-kivik/couchdb/v3"
)

func main() {
    cfg, _ := ini.Load("my.ini")
    bcAssetID := cfg.Section("").Key("BCAssetId").String()
    bcAssetType := cfg.Section("").Key("BCAssetType").String()
    BCAssetJsonSchemaFile := cfg.Section("").Key("BCAssetJsonSchemaFile").String()
    BCAssetIdGet := cfg.Section("").Key("BCAssetIdGet").String()
    bcKey := bcAssetType + "~" + bcAssetID
    file, _ := ioutil.ReadFile(BCAssetJsonSchemaFile)
    fmt.Println(minichain.PutState(bcKey, string(file)))
    fmt.Println(minichain.GetState(BCAssetIdGet))
}




