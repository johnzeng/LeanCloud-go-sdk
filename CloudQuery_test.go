package lean

import (
	"os"
	"testing"
)

//any better way?
type TestResp struct {
	Results []Test `json:"results,omitempty"`
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

func TestCloudQueryWithValue(t *testing.T) {
	client := NewClient(os.Getenv("LEAN_APPID"),
		os.Getenv("LEAN_APPKEY"),
		os.Getenv("LEAN_MASTERKEY"))
	agent := client.CloudQuery("select * from test where hi != ?", "hello")
	if err := agent.Do(); nil != err {
		t.Log(agent.superAgent.QueryData)
		t.Error(err.Error())
		return
	}

	ret := TestResp{}
	agent.ScanResponse(&ret)
	if len(ret.Results) == 0 {
		t.Log(agent.superAgent.Data)
		t.Error("cloud query failed")
		return
	}
	if ret.Results[0].Hello == "" {
		t.Error("cloud query can not get value")
		return
	}
	t.Log(agent.superAgent.QueryData)
}
