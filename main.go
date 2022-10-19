package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
)

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

var realm = flag.String("r", "", "sip realm")
var user = flag.String("u", "", "user")
var password = flag.String("p", "", "password")
var nonce = flag.String("n", "", "nonce")
var method = flag.String("m", "REGISTER", "sip method")

func main() {
	flag.Parse()

	var HA1 string
	var HA2 string
	var response string

	HA1 = MD5Hash(fmt.Sprintf("%s:%s:%s", *user, *realm, *password))
	HA2 = MD5Hash(fmt.Sprintf("%s:sip:%s", *method, *realm))

	response = MD5Hash(fmt.Sprintf("%s:%s:%s", HA1, *nonce, HA2))
	fmt.Println(response)
}
