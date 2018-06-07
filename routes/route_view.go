package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/kataras/iris"
	"github.com/zhuyanxi/goblockdemo/block"
)

// Index :
func Index(ctx iris.Context) {
	c := block.NewBlockChain("blockchain")
	allblock := c.AllBlock()

	ctx.View("index.html", iris.Map{
		"blocks": allblock,
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
	c := block.NewBlockChain("blockchain")
	allblock := c.AllBlock()
	ctx.View("block.html", iris.Map{
		"block":  allblock[0],
		"active": ctx.Path(),
	})
}
