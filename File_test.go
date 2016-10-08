package lean

import (
	"os"
	"testing"
)

func TestFileUpload(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	uploafile, openErr := os.Open("uploadtest.txt")
	if nil != openErr {
		t.Error(openErr.Error())
		return
	}
	agent := client.UploadFile("test.txt", "text/plain", uploafile)
	if err := agent.Do(); nil != err {
		t.Error(err.Error())
		return
	}
}
