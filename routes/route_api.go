package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/kataras/iris"
	"github.com/zhuyanxi/goblockdemo/block"
)

// SHA256 :
func SHA256(ctx iris.Context) {
	postData := ctx.PostValue("postData")
	hashData := sha256.Sum256([]byte(postData))

	hashStr := hex.EncodeToString(hashData[:])
	ctx.JSON(hashStr)
}

// ComputeBlockHash :
func ComputeBlockHash(ctx iris.Context) {
	var block block.Block
	ctx.ReadForm(&block)

	block.SetHash()

	ctx.JSON(hex.EncodeToString(block.Hash[:]))
}

// MineBlock : the function to solve the puzzle
// func MineBlock(ctx iris.Context) {
// 	Height, _ := ctx.PostValueInt64("Height")
// 	Data := ctx.PostValue("Data")
// 	PrevHash := []byte(ctx.PostValue("PrevHash"))
// 	//Nouce, _ := ctx.PostValueInt64("Nouce")
// 	block := block.NewBlock(Height, Data, PrevHash)

// 	hash := block.Hash[:]
// 	hashStr := hex.EncodeToString(hash[:])
// 	ctx.JSON(hashStr)
// }

// GetBalance :
func GetBalance(address string) {
	c := block.NewBlockChain("blockchain", "")
	balance := 0
	UTXOs := c.FindUTXO(address)
	for _, out := range UTXOs {
		balance += out.Value
	}
}
