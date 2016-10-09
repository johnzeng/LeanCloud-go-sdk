package lean

import (
	"os"
	"testing"
)

var fileId string

func TestUploadPlainText(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	//	uploafile, openErr := os.Open("uploadtest.txt")
	//	if nil != openErr {
	//		t.Error(openErr.Error())
	//		return
	//	}
	ret, err := client.UploadPlainText("test.txt", "this is a test")

	if nil != err {
		t.Error(err.Error())
		return
	}

	fileId = ret.ObjectId

	t.Logf("%v", ret)
}

func TestUploadFile(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	uploafile, openErr := os.Open("uploadtest.txt")
	if nil != openErr {
		t.Error(openErr.Error())
		return
	}
	ret, err := client.UploadFile("test2.txt", "image/png", uploafile)

	if nil != err {
		t.Error(err.Error())
		return
	}

	t.Logf("%v", ret)
}
