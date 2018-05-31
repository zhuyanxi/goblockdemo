package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	Util "github.com/zhuyanxi/goblockdemo/util"
)

// Block :
type Block struct {
	Height    int64
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nouce     int
}

// BDoc : the doc
type BDoc struct {
	ID      string `json:"_id"`
	Rev     string `json:"_rev"`
	Prev    string `json:"prev"`
	Blkjson string `json:"blkjson"`
}

// BTipDoc :
type BTipDoc struct {
	ID      string `json:"_id"`
	Rev     string `json:"_rev"`
	Tiphash string `json:"tiphash"`
}

// SetHash :
func (b *Block) SetHash() {
	var headers []byte
	timestamp := Util.IntToHex(b.Timestamp)
	height := Util.IntToHex(b.Height)
	nouce := Util.IntToHex(int64(b.Nouce))
	difficult := Util.IntToHex(int64(DifficultyBits))
	if b.Height == 0 {
		headers = bytes.Join(
			[][]byte{
				b.PrevHash,
				b.Data,
				height,
				difficult,
				nouce,
			},
			[]byte{},
		)
	} else {
		headers = bytes.Join(
			[][]byte{
				b.PrevHash,
				b.Data,
				timestamp,
				height,
				difficult,
				nouce,
			},
			[]byte{},
		)
	}
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock :
func NewBlock(height int64, data string, prevBlockHash []byte) *Block {
	block := &Block{height, time.Now().UnixNano(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewPOW(block)
	nouce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nouce = nouce
	return block
}

// ToJSONByte : return []byte of the block's json string
func (b *Block) ToJSONByte() []byte {
	jsonBlock, _ := json.Marshal(b)
	return jsonBlock
}

// DecodeJSONBlock :
func DecodeJSONBlock(d []byte) *Block {
	var block Block
	err := json.Unmarshal(d, &block)
	if err != nil {
		log.Println(err)
	}
	return &block
}

// GenerateBlockMap :
func (b *Block) GenerateBlockMap() map[string]string {
	r := make(map[string]string)
	blkjson, _ := json.Marshal(b)
	r["_id"] = hex.EncodeToString(b.Hash) //customize couchdb _id
	r["prev"] = hex.EncodeToString(b.PrevHash)
	r["blkjson"] = string(blkjson)
	return r
}
