package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	Util "github.com/zhuyanxi/goblockdemo/util"
)

var (
	maxNouce = math.MaxInt64
)

const difficultBits = 12

// ProofOfWork :
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//NewPOW :
func NewPOW(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficultBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nouce int) []byte {
	var data []byte

	if pow.block.Height == 0 {
		data = bytes.Join(
			[][]byte{
				pow.block.PrevHash,
				pow.block.Data,
				Util.IntToHex(pow.block.Height),
				Util.IntToHex(int64(difficultBits)),
				Util.IntToHex(int64(nouce)),
			},
			[]byte{},
		)
	} else {
		data = bytes.Join(
			[][]byte{
				pow.block.PrevHash,
				pow.block.Data,
				Util.IntToHex(pow.block.Timestamp),
				Util.IntToHex(pow.block.Height),
				Util.IntToHex(int64(difficultBits)),
				Util.IntToHex(int64(nouce)),
			},
			[]byte{},
		)
	}

	return data
}

// Run :
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nouce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nouce < maxNouce {
		data := pow.prepareData(nouce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		//fmt.Printf("\r%x", hash)
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			fmt.Println(nouce)
			nouce++
		}
	}
	fmt.Printf("\n\n")
	return nouce, hash[:]
}

// Validate :
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nouce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
