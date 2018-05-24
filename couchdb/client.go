package couchdb

import (
	"io"
	"log"
	"net/http"
)

// NewCouchClient :
func NewCouchClient(user, pwd, url string) CouchClient {
	return CouchClient{
		Username: user,
		Password: pwd,
		BaseURL:  url,
	}
}

// Request : request couchdb webapi
func (cc *CouchClient) Request(method, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(cc.Username, cc.Password)
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// DBInstance : return a database instance
func (cc *CouchClient) DBInstance(dbname string) Database {
	return Database{
		CouchClient: cc,
		Name:        dbname,
	}
}
