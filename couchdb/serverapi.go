package couchdb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ServerInfo : Accessing the root of a CouchDB instance returns meta information about the instance.
// http://docs.couchdb.org/en/2.1.1/api/server/common.html#
func (c CouchClient) ServerInfo() (*ServerInfo, error) {
	url := c.BaseURL
	res, err := c.Request(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var server ServerInfo
	err = json.Unmarshal(body, &server)
	return &server, err
}

// GetAllDB : Returns a list of all the databases in the CouchDB instance.
// http://docs.couchdb.org/en/2.1.1/api/server/common.html#all-dbs
func (c CouchClient) GetAllDB() ([]string, error) {
	url := c.BaseURL + "/_all_dbs"
	res, err := c.Request(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var dbs []string
	err = json.Unmarshal(body, &dbs)
	return dbs, nil
}
