package couchdb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetDoc : Returns document by the specified docid from the specified db.
// GET /{db}/{docid}
func (db *Database) GetDoc(id string) (*CouchDoc, *ResponseError, error) {
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

	var doc CouchDoc
	var reserr ResponseError
	log.Println(string(body))
	if res.StatusCode == 200 {
		err = json.Unmarshal(body, &doc)
	} else {
		err = json.Unmarshal(body, &reserr)
	}

	return &doc, &reserr, err
}
