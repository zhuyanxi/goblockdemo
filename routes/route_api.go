package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

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
	Height, _ := ctx.PostValueInt64("Height")
	Timestamp, _ := ctx.PostValueInt64("Timestamp")
	Data := []byte(ctx.PostValue("Data"))
	PrevHash := []byte(ctx.PostValue("PrevHash"))
	Hash := []byte(ctx.PostValue("Hash"))
	Nouce, _ := ctx.PostValueInt("Nouce")
	block := block.Block{Height, Timestamp, Data, PrevHash, Hash, Nouce}

	//var block block.Block
	//ctx.ReadJSON(&block)
	fmt.Printf("Height:%x\n", block.Height)
	fmt.Printf("Timestamp: %x\n", block.Timestamp)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Prev hash:%s\n", block.PrevHash)
	fmt.Println(block.PrevHash)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Nouce: %d\n", block.Nouce)
	fmt.Println()
}

// MineBlock : the function to solve the puzzle
func MineBlock(ctx iris.Context) {
	Height, _ := ctx.PostValueInt64("Height")
	Data := ctx.PostValue("Data")
	PrevHash := []byte(ctx.PostValue("PrevHash"))
	//Nouce, _ := ctx.PostValueInt64("Nouce")
	block := block.NewBlock(Height, Data, PrevHash)

	hash := block.Hash[:]
	hashStr := hex.EncodeToString(hash[:])
	ctx.JSON(hashStr)
}
