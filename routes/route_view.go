package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/kataras/iris"
	"github.com/zhuyanxi/goblockdemo/block"
	Util "github.com/zhuyanxi/goblockdemo/util"
)

// Index :
func Index(ctx iris.Context) {
	bc := block.NewBlockchain()
	Util.RandSleep()
	bc.AddBlock("Send 1 btc to Alice")
	Util.RandSleep()
	bc.AddBlock("Send 1.1 btc to Bob")

	ctx.View("index.html", iris.Map{
		"blocks": bc.Blocks,
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
	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 btc to Alice")
	ctx.View("block.html", iris.Map{
		"block":  bc.Blocks[1],
		"active": ctx.Path(),
	})
}