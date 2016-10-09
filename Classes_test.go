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
	Hello       string    `json:"hi,omitempty"`
	TestBytes   *LeanByte `json:"bytess,omitempty"`
	TestDate    *LeanTime `json:"tester,omitempty"`
	TestPointer *Pointer  `json:"user,omitempty"`
	notUpload   string    `json:"notUpload,omitempty"`
	Ignore      string    `json:"-"`
	Array       []string  `json:"ss,omitempty"`
}

var id string

//remember to uncomment this if you run this test first time otherwise you will get error on scan and query
/*func TestCreateObjectFirstRun(t *testing.T) {
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
}*/

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
		t.Logf("%x", ret)
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
	if len(ret.Results) == 0 {
		t.Errorf("message is wrong:%v", ret)
		t.Log(agent.body)
		return
	}
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

	updateObj := update.AddToArray("ss", "123", "456")

	agent := client.Collection("test").UpdateObjectById(id, updateObj)
	//if you don't wanna update by master key, you need to specify the id in update object body
	agent.UseMasterKey()
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
		t.Log(agent.superAgent.Data)
	}
}

func TestClassDelete(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.Collection("test").DeleteObjectById(id)
	agent.UseMasterKey()
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
	}
}

func TestClassScan(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.Collection("test").Scan("", "")
	q := query.Eq("hi", "this is first message")
	agent.WithQuery(q).Limit(1)
	agent.Do()
	ret := TestResp{}
	agent.ScanResponse(&ret)
	if len(ret.Results) == 0 {
		t.Errorf("message is wrong:%v", ret)
		t.Log(agent.body)
		return
	}
	if ret.Results[0].Hello != "this is first message" {
		t.Errorf("message is wrong:%v", ret)
		t.Log(agent.body)
	}

}
