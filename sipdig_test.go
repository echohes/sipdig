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
	*uri = "sip:127.0.0.1"

	response := calculateDigets()

	if response != "4aa71ec40e1c6bf271c2d5badcde5404" {
		t.Error("Incorrect hash ")
	}
}

func Test_Hash_Not(t *testing.T) {
	//values
	*user = "HE72tib"
	*password = "iCc2W5DsX4"
	*realm = "sip:127.0.0.1"
	*nonce = "Y0/V82NP1MegreIpiGEe1pbqYHqFdons"

	response := calculateDigets()

	if response == "4aa71ec40e1c6bf271c2d5badcde5404" {
		t.Error("Correct hash? ")
	}
}

func Test_Hash_INVITE_Response_OK(t *testing.T) {
	//values
	*checkString = `realm="127.0.0.1", nonce="Y1D5w2NQ+Jfd3lJixhK5CmoAy38b/VRs", username="Q-_HE72tib", uri="sip:121@127.0.0.1", response="66c17f995f15d80581f86d634076cf7f"`
	*password = "iCc2W5DsX4"
	*method = "INVITE"
	parseDigestResponce(*checkString)
	response := calculateDigets()

	if response != "66c17f995f15d80581f86d634076cf7f" {
		t.Error("Incorrect hash")
	}
}

func Test_Hash_REGISTER_Response_OK(t *testing.T) {
	//values
	*checkString = `realm="127.0.0.1", nonce="Y1D5w2NQ+Jfd3lJixhK5CmoAy38b/VRs", username="Q-_HE72tib", uri="sip:127.0.0.1", response="bc567020760b1f48982cbff3a0086318"`
	*password = "iCc2W5DsX4"
	*method = "REGISTER"
	parseDigestResponce(*checkString)
	response := calculateDigets()

	if response != "bc567020760b1f48982cbff3a0086318" {
		t.Error("Incorrect hash")
	}
}

func Test_Hash_GET_and_QOP_Response_OK(t *testing.T) {
	//values
	*checkString = `username="Mufasa",realm="testrealm@host.com",nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",uri="/dir/index.html",qop=auth,nc=00000001,cnonce="0a4f113b",response="6629fae49393a05397450978507c4ef1",opaque="5ccc069c403ebaf9f0171e9517f40e41"`
	*password = "Circle Of Life"
	*method = "GET"
	parseDigestResponce(*checkString)
	response := calculateDigets()

	if response != "6629fae49393a05397450978507c4ef1" {
		t.Error("Incorrect hash")
	}
}



