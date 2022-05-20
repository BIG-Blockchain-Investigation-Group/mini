package main

import (
	"fmt"
	"io/ioutil"
	"mini/minichain"

	_ "github.com/go-kivik/couchdb/v3"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, _ := ini.Load("my.ini")
	bcAssetID := cfg.Section("").Key("BCAssetId").String()
	bcAssetType := cfg.Section("").Key("BCAssetType").String()
	BCAssetJsonSchemaFile := cfg.Section("").Key("BCAssetJsonSchemaFile").String()
	bcKey := "Gateway~bc~" + bcAssetType + "~" + bcAssetID
	file, _ := ioutil.ReadFile(BCAssetJsonSchemaFile)
	fmt.Println(minichain.PutState(bcKey, file))
	fmt.Println(string(minichain.GetState((bcKey))))
}
