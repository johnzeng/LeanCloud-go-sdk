package lean

import (
	"os"
	"testing"
)

var userId string

//to avoid repeating same account create and delete, you have to un-comment this bock to run this test.
//func TestUserRegister(t *testing.T) {
//	client := NewClient(os.Getenv("LEAN_APPID"),
//		os.Getenv("LEAN_APPKEY"),
//		os.Getenv("LEAN_MASTERKEY"))
//	user := User{
//		Username:          "alsdkfj",
//		Email:             "",
//		MobilePhoneNumber: "lasdkfj",
//		Passowrd:          "la;dsfj",
//	}
//	agent := client.User.Create(user)
//	if err := agent.Do(); nil != err {
//		t.Error(err.Error())
//		return
//	}
//	ret := User{}
//
//	if err := agent.ScanResponse(&ret); nil != err {
//		t.Error(err.Error())
//		return
//	}
//	userId := ret.ObjectId
//	t.Log("userId :", userId)
//
//}

func TestUserLogin(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	user, err := client.Login(os.Getenv("LEAN_TEST_USER"), os.Getenv("LEAN_TEST_PWD"))
	if nil != err {
		t.Error(err.Error())
	}

	userId = user.ObjectId

	//also test token
	_, err = client.UserMe(user.SessionToken)
	if nil != err {
		t.Error(err.Error())
	}

}
