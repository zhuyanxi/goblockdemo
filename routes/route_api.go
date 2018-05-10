package routes

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/kataras/iris"
)

// SHA256 :
func SHA256(ctx iris.Context) {
	postData := ctx.PostValue("postData")
	hashData := sha256.Sum256([]byte(postData))

	hashStr := hex.EncodeToString(hashData[:])
	ctx.JSON(hashStr)
}
