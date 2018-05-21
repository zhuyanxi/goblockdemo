package util

import (
	"math/rand"
	"strconv"
	"time"
)

// RandSleep :
func RandSleep() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(200)
	//fmt.Println(x)
	time.Sleep(time.Duration(x) * time.Millisecond)
}

// IntToHex :
func IntToHex(num int64) []byte {
	buffer := []byte(strconv.FormatInt(num, 10))
	// buffer := new(bytes.Buffer)
	// err := binary.Write(buffer, binary.BigEndian, num)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// return buffer.Bytes()
	return buffer
}
