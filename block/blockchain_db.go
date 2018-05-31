package block

import (
	"encoding/hex"
	"log"

	"github.com/zhuyanxi/goblockdemo/couchdb"
)

const dbName = "blockchain"
const url = "http://127.0.0.1:5984"

// BlockChain :
type BlockChain struct {
	tipHash []byte
}

// newGenesisBlock :
func newGenesisBlock() *Block {
	return NewBlock(0, "The First Block", []byte{}) //6929, 22949
}

// NewBlockChain :
func NewBlockChain() *BlockChain {
	var tipHash []byte

	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	dbok, dberror, err := client.CreateDB(dbName)
	if err != nil {
		log.Println(err)
		return nil
	}
	if dbok.OK {
		gene := newGenesisBlock()
		tipHash = gene.Hash
		r := make(map[string]string)
		r["_id"] = "tiphash"
		r["tiphash"] = hex.EncodeToString(tipHash)

		db := client.DBInstance("block")
		db.NewDocument(r)
	} else {
		log.Println(dberror.Reason)
	}

	bc := BlockChain{tipHash: tipHash}
	return &bc
}
