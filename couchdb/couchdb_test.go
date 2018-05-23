package couchdb

import (
	"fmt"
	"testing"
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
	dbinfo, err := client.GetDbInfo("user")
	if err != nil {
		t.Fatal(err)
	}
	if dbinfo.DbName != "users" {
		t.Errorf("error dbinfo: %s", dbinfo.DbName)
	}
}

func TestCreateDB(t *testing.T) {
	client := NewCouchClient("zhuyx", "zhuyx123", url)
	dbok, dberror, err := client.CreateDB("user2")
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
	dbok, dberr, err := client.DeleteDB("user3")
	if err != nil {
		t.Fatal(err)
	}
	if !dbok.OK {
		t.Errorf("error dbinfo: %t", dbok.OK)
		t.Errorf("error dbinfo: %s", dberr.Reason)
	}
}
