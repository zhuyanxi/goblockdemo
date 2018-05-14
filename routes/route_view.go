package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

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
	Util.RandSleep()
	bc.AddBlock("Send 2 btc to zhuyx")
	Util.RandSleep()
	bc.AddBlock("Send 0.5 btc to John")

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

	fmt.Printf("Height: %x\n", bc.Blocks[1].Height)
	fmt.Printf("Timestamp: %x\n", bc.Blocks[1].Timestamp)
	fmt.Printf("Data: %s\n", bc.Blocks[1].Data)
	fmt.Printf("Prev hash:%x\n", bc.Blocks[1].PrevHash)
	fmt.Println(bc.Blocks[1].PrevHash)
	fmt.Printf("Hash: %x\n", bc.Blocks[1].Hash)
	fmt.Printf("Nouce: %d\n", bc.Blocks[1].Nouce)
	fmt.Println()

	ctx.View("block.html", iris.Map{
		"block":  bc.Blocks[1],
		"active": ctx.Path(),
	})
}
