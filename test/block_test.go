package test

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/zhuyanxi/goblockdemo/block"
	"github.com/zhuyanxi/goblockdemo/couchdb"
)

func TestNewBlockChain(t *testing.T) {
	t1 := time.Now()
	bc := block.NewBlockChain("blockchain", "39ZVV99CGTUnWHGuCNqGxb4EAeqTV1S92n")
	elapsed := time.Since(t1)
	log.Println("Execute time: ", elapsed)
	log.Println(bc)
	log.Println(hex.EncodeToString(bc.TipHash))

	if hex.EncodeToString(bc.TipHash) != "0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc" {
		t.Errorf("error blockchain create: %x", bc.TipHash)
	}
}

func TestGetFirstBlock(t *testing.T) {
	key := "000018f5fa26b88c1b4d7beeab7e4dfd0476e5f26ea43cf2c16c081832eb19e9"
	client := couchdb.NewCouchClient("zhuyx", "zhuyx123", url)
	db := client.DBInstance("blockchain")
	docjson, reserr, err := db.GetDoc(key)
	if err != nil {
		t.Fatal(err)
	}

	var doc block.BDoc
	err = json.Unmarshal(docjson, &doc)

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
	bdata := blk.Transactions
	log.Println(bdata)
	vin := bdata[0].Vin[0]
	vout := bdata[0].Vout[0]
	log.Println(vin)
	log.Println(vout)
	if vin.ScriptSig != "the genesis coin" {
		t.Errorf("error info: %s", vin.ScriptSig)
	}
	if vout.ScriptPubKey != "39ZVV99CGTUnWHGuCNqGxb4EAeqTV1S92n" {
		t.Errorf("error info: %s", vout.ScriptPubKey)
	}
}

// func TestAddBlock(t *testing.T) {
// 	bc := block.NewBlockChain("blockchain", "")
// 	height := bc.Height
// 	log.Println(bc)
// 	log.Println(hex.EncodeToString(bc.TipHash))

// 	err := bc.AddBlock("The Second block data.")
// 	if err != nil {
// 		t.Errorf("error add block: %s", err)
// 	}

// 	bcNew := block.NewBlockChain("blockchain")
// 	if height != bcNew.Height-1 {
// 		t.Error(bc)
// 		t.Error(bcNew)
// 	}
// }

// func TestPrevDoc(t *testing.T) {
// 	c := block.NewBlockChain("blockchain")
// 	height := c.Height
// 	log.Println(height)
// 	log.Println(hex.EncodeToString(c.TipHash))

// 	prev, _ := hex.DecodeString("0006cedcd17986875c9b28a4ed2c3e7f415d6e263e9bae13b1c9ed44663479cc")
// 	doc := c.PrevDoc(prev)
// 	if doc.Height != 0 {
// 		t.Error(doc)
// 	}
// }

// func TestAllBlock(t *testing.T) {
// 	c := block.NewBlockChain("blockchain")
// 	height := c.Height
// 	log.Println(height)
// 	log.Println(hex.EncodeToString(c.TipHash))

// 	allblock := c.AllBlock()
// 	if len(allblock) != 3 {
// 		t.Error(allblock)
// 	}
// }
