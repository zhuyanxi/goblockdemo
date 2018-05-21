package util

import (
	"net/url"

	"github.com/zemirco/couchdb"
)

//DBClient :
func DBClient() *couchdb.Client {
	u, err := url.Parse("http://127.0.0.1:5984/")
	if err != nil {
		panic(err)
	}
	// create a new client
	client, err := couchdb.NewAuthClient("zhuyx", "zhuyx123", u)
	if err != nil {
		panic(err)
	}
	return client
}
