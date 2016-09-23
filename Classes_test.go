package lean

import (
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
	client := NewClient("L1rboIrylg7wJklCPV8v6TCO-gzGzoHsz", "", "")
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
	client := NewClient("L1rboIrylg7wJklCPV8v6TCO-gzGzoHsz", "", "")
	client.GetObjectById("test", "57e4fd355bbb50005d499f3e").Do()
}
