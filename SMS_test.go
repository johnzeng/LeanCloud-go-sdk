package lean

import (
	"os"
	"testing"
)

func TestSMSVerifyRequest(t *testing.T) {
	println("now send request")
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))

	err := client.RequestMobilVerify(RequestMobilePhoneVerify{
		Number: os.Getenv("LEAN_TEST_PHONE_NUMBER"),
	})
	if nil != err {
		t.Error("Verify request error")
	}
	println("send request end")
}

func TestSMSVerify(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))

	err := client.VerifyCode(os.Getenv("LEAN_TEST_PHONE_NUMBER"), "")
	if nil != err {
		t.Error("Verify request error")
	}

}
