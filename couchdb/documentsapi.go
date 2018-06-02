package couchdb

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetDoc : Returns document by the specified docid from the specified db.
// GET /{db}/{docid}
// return []byte-docjson
func (db *Database) GetDoc(id string) ([]byte, *ResponseError, error) {
	client := db.CouchClient
	url := client.BaseURL + "/" + db.Name + "/" + id
	res, err := client.Request(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var doc []byte
	var reserr ResponseError
	log.Println(string(body))
	if res.StatusCode == 200 {
		doc = body
	} else {
		err = json.Unmarshal(body, &reserr)
	}

	return doc, &reserr, err
}

// UpdateDoc : The PUT method creates a new named document, or creates a new revision of the existing document.
// s is the json of data that is to be updated
// s must have the key "_rev" of the specified doc
func (db *Database) UpdateDoc(id string, s interface{}) (*ResponseDoc, *ResponseError, error) {
	client := db.CouchClient
	url := client.BaseURL + "/" + db.Name + "/" + id
	var payload bytes.Buffer
	if err := json.NewEncoder(&payload).Encode(s); err != nil {
		return nil, nil, err
	}
	res, err := client.Request(http.MethodPut, url, &payload)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	var doc ResponseDoc
	var reserr ResponseError
	log.Println(string(body))
	if res.StatusCode == 201 {
		err = json.Unmarshal(body, &doc)
	} else {
		err = json.Unmarshal(body, &reserr)
	}

	return &doc, &reserr, err
}
