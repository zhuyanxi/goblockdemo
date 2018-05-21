package block

import (
	"encoding/hex"
	"log"

	"github.com/zemirco/couchdb"
	"github.com/zhuyanxi/goblockdemo/util"
)

const dbName = "blockchain"

// BlockChain :
type BlockChain struct {
	tipHash []byte
}

type BlockChainDB struct {
	couchdb.Document
	BCMap map[string]BlockChain
}

// newGenesisBlock :
func newGenesisBlock() *Block {
	return NewBlock(0, "The First Block", []byte{}) //6929, 22949
}

// NewBlockChain :
func NewBlockChain() *BlockChain {
	var tipHash []byte
	client := util.DBClient()
	db, err := client.Get(dbName)
	if db == nil {
		log.Println(err)
		client.Create(dbName)
		_db := client.Use(dbName)

		gene := newGenesisBlock()
		bcDB := &BlockChainDB{BCMap: map[string]BlockChain{
			"tip_hash": BlockChain{tipHash: gene.Hash},
		}}
		_db.Post(bcDB)

		blkDB := &BlockDB{
			BlockMap: map[string]Block{
				hex.EncodeToString(gene.Hash): *gene,
			},
		}
		_db.Post(blkDB)

		tipHash = gene.Hash
	} else {

	}

	bc := BlockChain{tipHash: tipHash}
	return &bc
}
