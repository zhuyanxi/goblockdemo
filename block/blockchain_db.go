package block

import (
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/zhuyanxi/goblockdemo/couchdb"
)

//const dbName = "blockchain"
const url = "http://127.0.0.1:5984"

// BlockChain :
type BlockChain struct {
	TipHash []byte
	Height  int64
	DB      couchdb.Database
}

// newGenesisBlock :
func newGenesisBlock() *Block {
	return NewBlock(0, "The First Block", []byte{}) //6929, 22949
}

func getTipDoc(db couchdb.Database) *BTipDoc {
	tipdoc, _, err := db.GetDoc("tiphash")
	if err != nil {
		log.Println(err)
		return nil
	}

	var doc BTipDoc
	err = json.Unmarshal(tipdoc, &doc)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &doc
}

// NewBlockChain :
func NewBlockChain(dbName string) *BlockChain {
	var tipHash []byte
	var height int64

	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	dbok, dberror, err := client.CreateDB(dbName)
	if err != nil {
		log.Println(err)
		return nil
	}

	db := client.DBInstance(dbName)
	if dbok.OK {
		gene := newGenesisBlock()
		db.NewDocument(gene.GenerateBlockMap())

		tipHash = gene.Hash
		r := make(map[string]interface{})
		r["_id"] = "tiphash"
		r["tiphash"] = hex.EncodeToString(tipHash)
		r["height"] = 0

		res, reserr, err := db.NewDocument(r)
		if err != nil {
			log.Println(err)
		}
		if !res.OK {
			log.Printf("error dbinfo: %s", res.ID)
			log.Printf("error dbinfo: %s", reserr.Error)
		}
	} else {
		log.Println(dberror.Reason)
		doc := getTipDoc(db)
		tipHash, err = hex.DecodeString(doc.Tiphash)
		if err != nil {
			log.Println(err)
		}
		height = doc.Height
	}

	bc := BlockChain{TipHash: tipHash, Height: height, DB: db}
	return &bc
}

// AddBlock :
func (bc *BlockChain) AddBlock(data string) error {
	newBlock := NewBlock(bc.Height+1, data, bc.TipHash)

	res, reserr, err := bc.DB.NewDocument(newBlock.GenerateBlockMap())
	if err != nil {
		log.Println(err)
		return err
	}
	if !res.OK {
		log.Println(reserr.Reason)
		return err
	}

	tipdoc := getTipDoc(bc.DB)
	r := make(map[string]interface{})
	r["_rev"] = tipdoc.Rev
	r["tiphash"] = hex.EncodeToString(newBlock.Hash)
	r["height"] = newBlock.Height
	resdoc, reserr, err := bc.DB.UpdateDoc(tipdoc.ID, r)
	if err != nil {
		log.Println(err)
		return err
	}
	if !resdoc.OK {
		log.Println(err)
		log.Println(reserr.Reason)
		return err
	}

	return nil
}
