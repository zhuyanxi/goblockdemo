package block

import (
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/zhuyanxi/goblockdemo/couchdb"
)

//const dbName = "blockchain"
const url = "http://127.0.0.1:5984"
const geneCoinBaseData = "the genesis coin"

// Chain :
type Chain struct {
	TipHash []byte
	Height  int64
	DB      couchdb.Database
}

// newGenesisBlock :
func newGenesisBlock(basecoin *Transaction) *Block {
	return NewBlock(0, []*Transaction{basecoin}, []byte{}) //6929, 22949
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
// if the db is already exist, return the Chain entity;
// if the db is not exist, create the database and add the tip_doc and the genesis doc
func NewBlockChain(dbName, address string) *Chain {
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
		trans := NewCoinBase(address, geneCoinBaseData)
		gene := newGenesisBlock(trans)
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
func (c *Chain) AddBlock(transactions []*Transaction) error {
	newBlock := NewBlock(c.Height+1, transactions, c.TipHash)

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

// FindUTX : find and return a list of unspent transactions
func (c *Chain) FindUTX(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXs := make(map[string][]int)
	bc := c.AllBlock()
	for _, block := range bc {
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXs[txID] != nil {
					for _, spentOut := range spentTXs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXs[inTxID] = append(spentTXs[inTxID], in.Vout)
					}
				}
			}
		}

		// if len(block.PrevHash) == 0 {
		// 	break
		// }
	}

	return unspentTXs
}

// FindUTXO : find and return the list of unspent transaction outputs
func (c *Chain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTXs := c.FindUTX(address)
	for _, tx := range unspentTXs {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}
	return UTXOs
}
