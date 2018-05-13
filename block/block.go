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
	Nouce     int
}

// SetHash :
func (b *Block) SetHash() {
	var headers []byte
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	height := []byte(strconv.FormatInt(b.Height, 10))
	nouce := []byte(strconv.FormatInt(int64(b.Nouce), 10))
	if b.Height == 0 {
		headers = bytes.Join([][]byte{b.PrevHash, b.Data, height, nouce}, []byte{})
	} else {
		headers = bytes.Join([][]byte{b.PrevHash, b.Data, timestamp, height, nouce}, []byte{})
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
	//block.SetHash()
	return block
}
