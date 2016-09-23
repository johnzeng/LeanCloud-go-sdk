package lean

import (
	"os"
	"testing"
)

type Test struct {
	Hello     string `json:"hi"`
	notUpload string `json:"notUpload"`
	Ignore    string `json:"-"`
}

func (this Test) GetClassName() string {
	return "test"
}

func TestCreateObject(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	err := client.Create(
		Test{
			Hello:     "this is first message",
			notUpload: "nono",
			Ignore:    "ignore",
		}).Do()
	if nil != err {
		t.Error(err.Error())
	}
}

func TestGetObjectById(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	client.GetObjectById("test", "57e4fd355bbb50005d499f3e").Do()
}
