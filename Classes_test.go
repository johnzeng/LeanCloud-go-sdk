package lean

import (
	"github.com/johnzeng/leancloud-go-sdk/query"
	"github.com/johnzeng/leancloud-go-sdk/update"
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
	Array     []string `json:"ss,omitempty"`
}

var id string

func TestCreateObject(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.Collection("test").Create(
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
	agent := client.Collection("test").GetObjectById(id)
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
	}
}

func TestClassQuery(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.Collection("test").Query()
	q := query.Eq("hi", "this is first message")
	agent.WithQuery(q).Limit(1)
	agent.Do()
	ret := TestResp{}
	agent.ScanResponse(&ret)
	if ret.Results[0].Hello != "this is first message" {
		t.Errorf("message is wrong:%v", ret)
		t.Log(agent.body)
	}

}

func TestClassUpdate(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))

	test := Test{
		Array:    make([]string, 1),
		TestDate: NewLeanTime(time.Now()),
	}
	test.Array[0] = "hello"
	agent := client.Collection("test").UpdateObjectById(id, test)
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
		return
	}

}

func TestClassUpdateByPart(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))

	test := Test{
		Array: make([]string, 1),
	}

	updateObj := update.AddToArray("ss", "123", "456")

	test.Array[0] = "hello"
	agent := client.Collection("test").UpdateObjectById(id, updateObj)
	agent.UseMasterKey()
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
		t.Log(agent.superAgent.Data)
	}
}
