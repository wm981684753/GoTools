package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	secret = "yuxuan3507"
)

func main() {
	fmt.Println(`FS线上密匙自动生成脚本ing...
	
	`)
	generate()
}

//FS密匙生成脚本
func generate() {
	//data生成
	nonce := randomString(16)
	timestamps := time.Now().Unix()
	payload := "[]"
	data := strconv.FormatInt(timestamps, 10) + nonce + payload

	//hash_hmac
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	hash := hex.EncodeToString(h.Sum(nil))

	//base64_encode
	var msg = []byte(hash)
	signature := base64.StdEncoding.EncodeToString(msg)
	fmt.Println("nonce：", nonce)
	fmt.Println("timestamps：", timestamps)
	fmt.Println("signature：", signature)
	fmt.Println("apiKey：", "yuxuanxuanpc")

	scale := "a"
	fmt.Println(`
回车重新生成密匙
	`)
	fmt.Scanln(&scale)
	if scale != "" {
		generate()
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
