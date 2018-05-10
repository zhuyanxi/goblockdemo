package util

import (
	"fmt"
	"math/rand"
	"time"
)

// RandSleep :
func RandSleep() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(200)
	fmt.Println(x)
	time.Sleep(time.Duration(x) * time.Millisecond)
}
