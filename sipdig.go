package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strings"
)

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

var realm = flag.String("r", "", "sip realm")
var uri = flag.String("ur", "", "sip uri, example: sip:127.0.0.1:5060")
var user = flag.String("u", "", "user")
var password = flag.String("p", "", "password")
var nonce = flag.String("n", "", "nonce")
var method = flag.String("m", "REGISTER", "sip method")
var checkString = flag.String("d", "", "validate your response Digest")
var checkResponse string

func main() {
	flag.Parse()

	if len(*checkString) > 0 && len(*password) > 0 {
		parseDigestResponce(*checkString)
	} else {
		flag.VisitAll(func(f *flag.Flag) {
			if f.Value.String() == "" && f.Name != "d" {
				fmt.Println(f.Name, "empty")
			}
		})
	}

	response := calculateDigets()

	if len(checkResponse) > 0 && checkResponse == response {
		fmt.Println("Responcse is valid")
	} else {
		fmt.Println("Responcse is invalid")
	}

	fmt.Println(response)
}

func calculateDigets() string {

	var HA1 string
	var HA2 string
	var response string

	HA1 = MD5Hash(fmt.Sprintf("%s:%s:%s", *user, *realm, *password))
	HA2 = MD5Hash(fmt.Sprintf("%s:%s", *method, *uri))

	response = MD5Hash(fmt.Sprintf("%s:%s:%s", HA1, *nonce, HA2))
	return response
}

func parseDigestResponce(string) {

	removeSpace := strings.ReplaceAll(*checkString, " ", "")
	removeQutos := strings.ReplaceAll(removeSpace, "\"", "")
	splitParams := strings.Split(removeQutos, ",")

	for _, v := range splitParams {
		switch params := strings.Split(v, "="); params[0] {
		case "nonce":
			*nonce = params[1]
		case "realm":
			*realm = params[1]
		case "uri":
			*uri = params[1]
		case "username":
			*user = params[1]
		case "response":
			checkResponse = params[1]
		}
	}
}
