package block

import (
	"bytes"
	"crypto/sha256"
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
	//block.SetHash()
	return block
}
