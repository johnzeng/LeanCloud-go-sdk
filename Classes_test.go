package lean

import (
	"os"
	"testing"
	"time"
)

type Test struct {
	LeanClassesBase
	Hello     string   `json:"hi"`
	TestBytes LeanByte `json:"bytess"`
	TestDate  LeanTime `json:"tester"`
	notUpload string   `json:"notUpload"`
	Ignore    string   `json:"-"`
}

func (this Test) GetClassName() string {
	return "test"
}

var id string

func TestCreateObject(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.Create(
		Test{
			Hello:     "this is first message",
			notUpload: "nono",
			Ignore:    "ignore",
			TestDate:  NewLeanTime(time.Now()),
		})

	if err := agent.Do(); nil != err {
		t.Error(err.Error())
	}
	ret := Test{}

	if err := agent.ScanResponse(&ret); err != nil {
		t.Error(err.Error())
	} else {
		id = ret.ObjectId
	}
}

func TestGetObjectById(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.GetObjectById("test", id)
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
	}

	ret := Test{}
	if err := agent.ScanResponse(&ret); nil != err {
		t.Error(err.Error())
	} else {
		if ret.Hello != "this is first message" {
			t.Error("message is wrong")
		}
		println(ret.UpdatedAt.String())
	}
}
