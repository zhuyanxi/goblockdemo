package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block :
type Block struct {
	Height    int64
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
}

// SetHash :
func (b *Block) SetHash() {
	//timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	height := []byte(strconv.FormatInt(b.Height, 10))
	//headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp, height}, []byte{})
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, height}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock :
func NewBlock(height int64, data string, prevBlockHash []byte) *Block {
	block := &Block{height, time.Now().UnixNano(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
