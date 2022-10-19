package main

import (
	"testing"
)

func Test_Hash_OK(t *testing.T) {
	//values
	*user = "Q-_HE72tib"
	*password = "iCc2W5DsX4"
	*realm = "127.0.0.1"
	*nonce = "Y0/V82NP1MegreIpiGEe1pbqYHqFdons"

	response := calculdateDigets()

	if response != "4aa71ec40e1c6bf271c2d5badcde5404" {
		t.Error("Incorrect hash ")
	}
}

func Test_Hash_Not(t *testing.T) {
	//values
	*user = "HE72tib"
	*password = "iCc2W5DsX4"
	*realm = "127.0.0.1"
	*nonce = "Y0/V82NP1MegreIpiGEe1pbqYHqFdons"

	response := calculdateDigets()

	if response == "4aa71ec40e1c6bf271c2d5badcde5404" {
		t.Error("Correct hash? ")
	}
}
