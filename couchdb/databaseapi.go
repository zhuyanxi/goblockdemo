package couchdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetDbInfo : Gets information about the specified database.
func (cc *CouchClient) GetDbInfo(dbname string) (*DatabaseInfo, error) {
	url := cc.BaseURL + "/" + dbname
	res, err := cc.Request(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var dbinfo DatabaseInfo
	err = json.Unmarshal(body, &dbinfo)
	return &dbinfo, err
}

// CreateDB : Creates a new database.
func (cc *CouchClient) CreateDB(dbname string) (*ResponseOK, *ResponseError, error) {
	url := cc.BaseURL + "/" + dbname
	res, err := cc.Request(http.MethodPut, url, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	var ok ResponseOK
	var reserr ResponseError
	log.Println(string(body))
	if res.StatusCode == 201 {
		err = json.Unmarshal(body, &ok)
	} else {
		err = json.Unmarshal(body, &reserr)
	}

	return &ok, &reserr, err
}

// DeleteDB : Deletes the specified database, and all the documents and attachments contained within it.
func (cc *CouchClient) DeleteDB(dbname string) (*ResponseOK, *ResponseError, error) {
	url := cc.BaseURL + "/" + dbname
	res, err := cc.Request(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var ok ResponseOK
	var reserr ResponseError
	log.Println(string(body))
	if res.StatusCode == 200 {
		err = json.Unmarshal(body, &ok)
	} else {
		err = json.Unmarshal(body, &reserr)
	}

	return &ok, &reserr, err
}

// NewDocument : Creates a new document in the specified database, using the supplied JSON document structure.
// s can be a struct, map or any other valid data structure
func (db *Database) NewDocument(s interface{}) (*ResponseDoc, *ResponseError, error) {
	client := db.CouchClient
	url := client.BaseURL + "/" + db.Name
	var payload bytes.Buffer
	if err := json.NewEncoder(&payload).Encode(s); err != nil {
		return nil, nil, err
	}
	res, err := client.Request(http.MethodPost, url, &payload)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	log.Println(string(body))
	var docres ResponseDoc
	var reserr ResponseError
	if res.StatusCode == 201 {
		err = json.Unmarshal(body, &docres)
	} else {
		err = json.Unmarshal(body, &reserr)
	}
	return &docres, &reserr, err
}

// GetDocsByKeys : Returns a JSON structure of all of the documents in a given database.
// The Get _all_docs allows to specify multiple keys to be selected from the database.
// Param (s []string) is the  string of keys, like ["key1","key2"]
func (db *Database) GetAllDocsByKeys(s []string) (*CouchDocument, error) {
	client := db.CouchClient
	//url := client.BaseURL + "/" + db.Name+"/_all_docs"
	keys, _ := json.Marshal(s)
	param := "include_docs=true&keys=" + string(keys)
	url := client.BaseURL + fmt.Sprintf("/%s/_all_docs?%s", db.Name, param)
	res, err := client.Request(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))
	var doc CouchDocument
	err = json.Unmarshal(body, &doc)

	return &doc, err
}
