package lean

/*
这个测试用例需要分两次进行，第一次跑VerifyRequest，得到验证码之后跑SMSVerify
另外，需要在设置->应用选项->其他中，打开启用通用的短信验证码服务（开放 requestSmsCode 和 verifySmsCode 接口） 选项。
这个接口独立于区别于VerfiyMobilePhone
*/

//import (
//	"os"
//	"testing"
//)

//func TestSMSVerifyRequest(t *testing.T) {
//	println("now send request")
//	client := NewClient(os.Getenv("LEAN_APPID"),
//		os.Getenv("LEAN_APPKEY"),
//		os.Getenv("LEAN_MASTERKEY"))
//
//	err := client.RequestMobilVerify(RequestMobilePhoneVerify{
//		Number: os.Getenv("LEAN_TEST_PHONE_NUMBER"),
//	})
//	if nil != err {
//		t.Error("Verify request error")
//	}
//	println("send request end")
//}

//func TestSMSVerify(t *testing.T) {
//	client := NewClient(os.Getenv("LEAN_APPID"),
//		os.Getenv("LEAN_APPKEY"),
//		os.Getenv("LEAN_MASTERKEY"))
//
//	err := client.VerifyCode(os.Getenv("LEAN_TEST_PHONE_NUMBER"), "062537")
//	if nil != err {
//		t.Error("Verify request error")
//	}
//
//}
