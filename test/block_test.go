package test

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"testing"

	"github.com/zhuyanxi/goblockdemo/block"
	"github.com/zhuyanxi/goblockdemo/couchdb"
)

func TestNewBlockChain(t *testing.T) {
	bc := block.NewBlockChain("blockchain")
	log.Println(bc)
	log.Println(hex.EncodeToString(bc.TipHash))

	if hex.EncodeToString(bc.TipHash) != "0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc" {
		t.Errorf("error blockchain create: %x", bc.TipHash)
	}
}

func TestGetFirstBlock(t *testing.T) {
	key := "0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc"
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	db := client.DBInstance("blockchain")
	docjson, reserr, err := db.GetDoc(key)
	if err != nil {
		t.Fatal(err)
	}

	var doc block.BDoc
	err = json.Unmarshal(docjson, &doc)

	// log.Println(doc.ID)
	// log.Println(doc.Blkjson)
	data, _ := hex.DecodeString(doc.Blkjson)
	log.Println(data)
	log.Println(string(data))

	if doc.ID != key {
		t.Errorf("error dbinfo: %s", doc.Blkjson)
		t.Errorf("error info: %s", reserr.Error)
	}

	var blk block.Block
	err = json.Unmarshal(data, &blk)
	if err != nil {
		log.Println(err)
	}
	bdata := string(blk.Data)
	log.Println(bdata)
	if bdata != "The First Block" {
		t.Errorf("error info: %s", bdata)
	}
}

func TestAddBlock(t *testing.T) {
	bc := block.NewBlockChain("blockchain")
	height := bc.Height
	log.Println(bc)
	log.Println(hex.EncodeToString(bc.TipHash))

	err := bc.AddBlock("The Second block data.")
	if err != nil {
		t.Errorf("error add block: %s", err)
	}

	bcNew := block.NewBlockChain("blockchain")
	if height != bcNew.Height-1 {
		t.Error(bc)
		t.Error(bcNew)
	}
}

func TestPrevDoc(t *testing.T) {
	c := block.NewBlockChain("blockchain")
	height := c.Height
	log.Println(height)
	log.Println(hex.EncodeToString(c.TipHash))

	prev, _ := hex.DecodeString("0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc")
	doc := c.PrevDoc(prev)
	if doc.Height != 0 {
		t.Error(doc)
	}
}

func TestAllBlock(t *testing.T) {
	c := block.NewBlockChain("blockchain")
	height := c.Height
	log.Println(height)
	log.Println(hex.EncodeToString(c.TipHash))

	allblock := c.AllBlock()
	if len(allblock) != 3 {
		t.Error(allblock)
	}
}
