package block

import (
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/zhuyanxi/goblockdemo/couchdb"
)

//const dbName = "blockchain"
const url = "http://127.0.0.1:5984"

// Chain :
type Chain struct {
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

// NewBlockChain : Create the BlockChain database
// if the db is already exist, return the BlockChain entity;
// if the db is not exist, create the database and add the tip_doc and the genesis doc
func NewBlockChain(dbName string) *Chain {
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

	bc := Chain{TipHash: tipHash, Height: height, DB: db}
	return &bc
}

// AddBlock :
func (c *Chain) AddBlock(data string) error {
	newBlock := NewBlock(c.Height+1, data, c.TipHash)

	res, reserr, err := c.DB.NewDocument(newBlock.GenerateBlockMap())
	if err != nil {
		log.Println(err)
		return err
	}
	if !res.OK {
		log.Println(reserr.Reason)
		return err
	}

	tipdoc := getTipDoc(c.DB)
	r := make(map[string]interface{})
	r["_rev"] = tipdoc.Rev
	r["tiphash"] = hex.EncodeToString(newBlock.Hash)
	r["height"] = newBlock.Height
	resdoc, reserr, err := c.DB.UpdateDoc(tipdoc.ID, r)
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

// PrevDoc :
func (c *Chain) PrevDoc(prevHash []byte) *Block {
	docbyte, _, err := c.DB.GetDoc(hex.EncodeToString(prevHash))
	if err != nil {
		return nil
	}
	return DecodeJSONBlock(docbyte)
}

// AllBlock : return the array of all Block
func (c *Chain) AllBlock() []Block {
	var blkArr []Block
	doc := c.PrevDoc(c.TipHash)
	blkArr = append(blkArr, *doc)

	for {
		if hex.EncodeToString(doc.PrevHash) == "" {
			break
		}
		doc = c.PrevDoc(doc.PrevHash)
		blkArr = append(blkArr, *doc)
	}
	return blkArr
}
