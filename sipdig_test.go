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
