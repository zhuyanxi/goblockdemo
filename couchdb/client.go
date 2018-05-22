package couchdb

import (
	"io/ioutil"
	"log"
	"net/http"
)

// CouchClient :
type CouchClient struct {
	Username string
	Password string
	BaseURL  string
}

// NewCouchClient :
func NewCouchClient(user, pwd, url string) CouchClient {
	return CouchClient{
		Username: user,
		Password: pwd,
		BaseURL:  url,
	}
}

// DBInfo :
func (c CouchClient) DBInfo() string {
	url := c.BaseURL
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
