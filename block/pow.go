package block

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"math"
	"math/big"
)

var (
	maxNouce = math.MaxInt64
)

// DifficultyBits : define the hash string of data starts with how many 0 bits
const DifficultyBits = 16

// ProofOfWork :
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewPOW :
func NewPOW(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DifficultyBits))

	pow := &ProofOfWork{b, target}
	return pow
}

// PrepareData :
func (pow *ProofOfWork) PrepareData(nouce int) []byte {
	block := pow.block
	block.Nouce = nouce
	data, _ := json.Marshal(block)

	// var data []byte
	// transhash := pow.block.GenerateTransactionHash()
	// if pow.block.Height == 0 {
	// 	data = bytes.Join(
	// 		[][]byte{
	// 			pow.block.PrevHash,
	// 			transhash,
	// 			Util.IntToHex(pow.block.Height),
	// 			Util.IntToHex(int64(DifficultyBits)),
	// 			Util.IntToHex(int64(nouce)),
	// 		},
	// 		[]byte{},
	// 	)
	// } else {
	// 	data = bytes.Join(
	// 		[][]byte{
	// 			pow.block.PrevHash,
	// 			transhash,
	// 			Util.IntToHex(pow.block.Timestamp),
	// 			Util.IntToHex(pow.block.Height),
	// 			Util.IntToHex(int64(DifficultyBits)),
	// 			Util.IntToHex(int64(nouce)),
	// 		},
	// 		[]byte{},
	// 	)
	// }
	//log.Println(hex.EncodeToString(data))
	//log.Println(data1)

	return data
}

// Run :
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nouce := 0

	//fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nouce < maxNouce {
		data := pow.PrepareData(nouce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		//fmt.Printf("\r%x", hash)
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			log.Println(nouce)
			nouce++
		}
	}
	//fmt.Printf("\n")
	return nouce, hash[:]
}

// Validate :
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nouce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
