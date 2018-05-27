package couchdb

import (
	"fmt"
	"testing"

	"github.com/zhuyanxi/goblockdemo/block"
)

const url = "http://127.0.0.1:5984"

// TestDBInfo : test func serverapi->DBInfo
func TestDBInfo(t *testing.T) {
	client := NewCouchClient("zhuyx", "zhuyx123", url)
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
	client := NewCouchClient("zhuyx", "zhuyx123", url)
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
	client := NewCouchClient("zhuyx", "zhuyx123", url)
	dbinfo, err := client.GetDbInfo("_users")
	if err != nil {
		t.Fatal(err)
	}
	if dbinfo.DbName != "users" {
		t.Errorf("error dbinfo: %s", dbinfo.DbName)
	}
}

func TestCreateDB(t *testing.T) {
	client := NewCouchClient("zhuyx", "zhuyx123", url)
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
	client := NewCouchClient("zhuyx", "zhuyx123", url)
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
	blk := block.NewBlock(0, "The First Block", []byte{})
	blkdb := blk.GenerateBlockMap()

	client := NewCouchClient("zhuyx", "zhuyx123", url)
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

func TestGetDocsByKeys(t *testing.T) {
	keys := []string{"0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc"}

	client := NewCouchClient("zhuyx", "zhuyx123", url)
	db := client.DBInstance("block")
	docs, err := db.GetDocsByKeys(keys)
	if err != nil {
		t.Fatal(err)
	}
	if docs.TotalRows != 4 {
		t.Errorf("error dbinfo: %s", docs.Rows[0].Doc)
	}
	// doc0 := docs.Rows[0].Doc

	// var blk block.Block
	// blk, ok := doc0.(map[string]interface{})
}
