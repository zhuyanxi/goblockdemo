package couchdb

import (
	"encoding/json"
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

// NewDocument :
func (db *Database) NewDocument(e interface{}) (*ResponseDoc, error) {
	client := db.CouchClient
	url := client.BaseURL + "/" + db.Name
	res, err := client.Request(http.MethodPost, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var docres ResponseDoc
	err = json.Unmarshal(body, &docres)
	return &docres, err
}
