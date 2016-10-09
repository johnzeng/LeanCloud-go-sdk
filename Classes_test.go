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
	Hello        string        `json:"hi,omitempty"`
	TestBytes    *LeanByte     `json:"bytess,omitempty"`
	TestDate     *LeanTime     `json:"tester,omitempty"`
	TestRelation *LeanRelation `json:"user,omitempty"`
	//we have some problem on LeanFile API
	TestFile    *FullLeanFile `json:"filePtr,omitempty"`
	TestPointer *LeanPointer  `json:"userPtr,omitempty"`
	notUpload   string        `json:"notUpload,omitempty"`
	Ignore      string        `json:"-"`
	Array       []string      `json:"ss,omitempty"`
}

var id string

//remember to uncomment this if you run this test first time otherwise you will get error on scan and query
//func TestCreateObjectFirstRun(t *testing.T) {
//	client := NewClient(os.Getenv("LEAN_APPID"),
//		os.Getenv("LEAN_APPKEY"),
//		os.Getenv("LEAN_MASTERKEY"))
//	now := NewLeanTime(time.Now())
//	agent := client.Collection("test").Create(
//		Test{
//			Hello:     "this is first message",
//			notUpload: "nono",
//			Ignore:    "ignore",
//			TestDate:  &now,
//		})
//
//	if err := agent.Do(); nil != err {
//		t.Error(err.Error())
//	}
//}

func TestCreateObject(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	now := NewLeanTime(time.Now())
	bytes := NewLeanByte([]byte("hello"))
	agent := client.Collection("test").Create(
		Test{
			Hello:     "this is first message",
			TestBytes: &bytes,
			notUpload: "nono",
			Ignore:    "ignore",
			TestDate:  &now,
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
	t.Log("%v", agent.superAgent.Data)
}

func TestClassQuery(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.Collection("test").Query()
	q := query.Eq("hi", "this is first message")
	agent.WithQuery(q).Limit(1).Skip(1).Order("hi").WithCount().UseSignature()
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

	TestUploadPlainText(t)
	TestUserLogin(t)

	now := NewLeanTime(time.Now())
	pointer := LeanPointer{class: "_user", objectId: userId}
	//	filePtr := LeanFile{Id: fileId}
	test := Test{
		Array:       make([]string, 1),
		TestDate:    &now,
		TestPointer: &pointer,
		//		TestFile:    &filePtr,
	}
	test.Array[0] = "hello"
	agent := client.Collection("test").UpdateObjectById(id, test)
	if err := agent.Do(); nil != err {
		t.Log("%v", agent.superAgent.Data)
		t.Error(err.Error())
		return
	}

	t.Log("%v", agent.superAgent.Data)

}

func TestClassUpdateByPart(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))

	addObj := update.AddToArray("ss", "123", "456")
	addFile := update.Set("filePtr", LeanFile{Id: fileId})
	updateObj := update.AddRelation("user", LeanPointer{class: "_user", objectId: userId}).And(addObj).And(addFile)

	agent := client.Collection("test").UpdateObjectById(id, updateObj)
	//if you don't wanna update by master key, you need to specify the id in update object body
	agent.UseMasterKey()
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
		t.Log(agent.superAgent.Data)
	}

	ret := Test{}
	if err := agent.ScanResponse(&ret); nil != err {

		t.Error(err.Error())
		t.Log(agent.superAgent.Data)
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
