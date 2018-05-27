package couchdb

// ResponseError :
type ResponseError struct {
	Error  string
	Reason string
}

// ResponseOK :
type ResponseOK struct {
	OK bool
}

// ResponseDoc :
type ResponseDoc struct {
	ID  string
	OK  bool
	rev string
}

// CouchClient :
type CouchClient struct {
	Username string
	Password string
	BaseURL  string
}

// ServerInfo : root of a CouchDB instance.
// http://docs.couchdb.org/en/2.1.1/api/server/common.html#
type ServerInfo struct {
	Couchdb  string
	Version  string
	Features []string
	Vendor   struct {
		Name string
	}
}

// Database :
type Database struct {
	CouchClient *CouchClient
	Name        string
}

// DatabaseInfo :
type DatabaseInfo struct {
	DbName    string `json:"db_name"`
	UpdateSeq string `json:"update_seq"`
	Sizes     struct {
		File     int64 //`json:"file"`
		External int64 //`json:"external"`
		Active   int64 //`json:"active"`
	}
	PurgeSeq int64 `json:"purge_seq"`
	Other    struct {
		DataSize int64 `json:"data_size"`
	}
	DocDelCount       int64 `json:"doc_del_count"`
	DocCount          int64 `json:"doc_count"`
	DiskSize          int64 `json:"disk_size"`
	DiskFormatVersion int64 `json:"disk_format_version"`
	DataSize          int64 `json:"data_size"`
	CompactRunning    bool  `json:"compact_running"`
	Cluster           struct {
		Q int64
		N int64
		W int64
		R int64
	}
	InstanceStartTime string `json:"instance_start_time"`
}

// CouchDocument : the struct of documents returned form api:GET /{db}/_all_docs
type CouchDocument struct {
	TotalRows int           `json:"total_rows"`
	Rows      []CouchDocRow `json:"rows"`
}

// CouchDocRow : the document's rows
type CouchDocRow struct {
	ID    string                 `json:"id"`
	Key   interface{}            `json:"key"`
	Value interface{}            `json:"value"`
	Doc   map[string]interface{} `json:"doc"`
}
