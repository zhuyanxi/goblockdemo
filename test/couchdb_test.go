package test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/zhuyanxi/goblockdemo/block"
	"github.com/zhuyanxi/goblockdemo/couchdb"
)

const url = "http://127.0.0.1:5984"

// TestDBInfo : test func serverapi->DBInfo
func TestDBInfo(t *testing.T) {
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	info, err := client.ServerInfo()
	fmt.Println(info)
	if err != nil {
		t.Fatal(err)
	}
	if info.Couchdb != "Welcome" {
		t.Errorf("error couchdb info: %s", info.Couchdb)
	}
	if info.Vendor.Name != "The Apache Software Foundation" {
		t.Errorf("error vendor name: %s", info.Vendor.Name)
	}
}

// TestGetAllDB : test serverapi->GetAllDB
func TestGetAllDB(t *testing.T) {
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	dbs, err := client.GetAllDB()
	if err != nil {
		t.Fatal(err)
	}
	if dbs[0] != "dummy" {
		t.Errorf("error dbs: %s", dbs)
	}
	if dbs[1] != "user" {
		t.Errorf("error dbs: %s", dbs)
	}
}

func TestGetDBinfo(t *testing.T) {
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	dbinfo, err := client.GetDbInfo("_users")
	if err != nil {
		t.Fatal(err)
	}
	if dbinfo.DbName != "users" {
		t.Errorf("error dbinfo: %s", dbinfo.DbName)
	}
}

func TestCreateDB(t *testing.T) {
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	dbok, dberror, err := client.CreateDB("block")
	if err != nil {
		t.Fatal(err)
	}
	if !dbok.OK {
		t.Errorf("error dbinfo: %t", dbok.OK)
		t.Errorf("error dbinfo: %s", dberror.Reason)
	}
}

func TestDeleteDB(t *testing.T) {
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	dbok, dberr, err := client.DeleteDB("block")
	if err != nil {
		t.Fatal(err)
	}
	if !dbok.OK {
		t.Errorf("error dbinfo: %t", dbok.OK)
		t.Errorf("error dbinfo: %s", dberr.Reason)
	}
}

func TestNewDocument(t *testing.T) {
	btt, _ := hex.DecodeString("0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc")
	//btt := []byte{}
	blk := block.NewBlock(1, "The Second Block", btt)
	blkdb := blk.GenerateBlockMap()

	// blkdb := make(map[string]string)
	// blkdb["_id"] = "tiphash"
	// blkdb["tiphash"] = "0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc"

	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	db := client.DBInstance("block")
	res, reserr, err := db.NewDocument(blkdb)
	if err != nil {
		t.Fatal(err)
	}
	if !res.OK {
		t.Errorf("error dbinfo: %s", res.ID)
		t.Errorf("error dbinfo: %s", reserr.Error)
	}
}

func TestGetAllDocsByKeys(t *testing.T) {
	keys := []string{"0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc"}

	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	db := client.DBInstance("block")
	docs, err := db.GetAllDocsByKeys(keys)
	if err != nil {
		t.Fatal(err)
	}
	if docs.TotalRows != 4 {
		t.Errorf("error dbinfo: %s", docs.Rows[0].Doc)
	}
}

func TestGetDoc(t *testing.T) {
	key := "0004d003fe2306112dba354ed28a3c6924572928d395471f417d36b89a7a365c"
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	db := client.DBInstance("block")
	docjson, reserr, err := db.GetDoc(key)
	if err != nil {
		t.Fatal(err)
	}

	var doc block.BDoc
	err = json.Unmarshal(docjson, &doc)

	if doc.ID != key {
		t.Errorf("error dbinfo: %s", doc.Blkjson)
		t.Errorf("error info: %s", reserr.Error)
	}
}

// func TestUpdateDoc(t *testing.T) {
// 	key := "0004d003fe2306112dba354ed28a3c6924572928d395471f417d36b89a7a365c"
// 	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
// 	db := client.DBInstance("block")
// 	docjson, reserr, err := db.GetDoc(key)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var doc block.BDoc
// 	err = json.Unmarshal(docjson, &doc)
// 	if doc.ID != key {
// 		t.Errorf("error dbinfo: %s", doc.Blkjson)
// 		t.Errorf("error info: %s", reserr.Error)
// 	}

// 	var block block.Block
// 	err = json.Unmarshal([]byte(doc.Blkjson), &block)

// 	payload := make(map[string]string)
// 	payload["_rev"] = doc.Rev
// 	//doc, reserr, err := db.UpdateDoc(key)
// }
