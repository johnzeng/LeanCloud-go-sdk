package lean

import (
	"os"
	"testing"
)

//any better way?
type TestResp struct {
	Results []Test `json:"results,omitepty"`
	Count   int64  `json:"count,omitempty"`
}

func TestCloudQuery(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.CloudQuery("select * from test")
	agent.Do()
	ret := TestResp{}
	agent.ScanResponse(&ret)
	if len(ret.Results) == 0 {
		t.Error("cloud query failed")
		return
	}
	if ret.Results[0].Hello == "" {
		t.Error("cloud query can not get value")
		return

	}
}
