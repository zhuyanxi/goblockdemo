package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/kataras/iris"
	"github.com/zhuyanxi/goblockdemo/block"
)

// Index :
func Index(ctx iris.Context) {
	bc := block.NewBlockChain("blockchain")
	log.Println(bc)
	var blkArr []block.Block

	ctx.View("index.html", iris.Map{
		"blocks": blkArr,
		"active": ctx.Path(),
	})
}

// Hash :
func Hash(ctx iris.Context) {
	hash := sha256.Sum256([]byte{})
	ctx.View("hash.html", iris.Map{
		"hash":   hex.EncodeToString(hash[:]),
		"active": ctx.Path(),
	})
}

// Block :
func Block(ctx iris.Context) {
	bc := block.NewBlockChain("blockchain")
	bc.AddBlock("Send 1 btc to Alice")
	var blkArr []block.Block

	ctx.View("block.html", iris.Map{
		"block":  blkArr[1],
		"active": ctx.Path(),
	})
}
