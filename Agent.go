package lean

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/johnzeng/leancloud-go-sdk/query"
	"github.com/parnurzeal/gorequest"
	"time"
)

type Agent struct {
	client       leanClient
	superAgent   *gorequest.SuperAgent
	body         string
	useSignature bool
	useMasterKey bool
	ts           int64
}

type QueryAgent struct {
	Agent
}

type UpdateAgent struct {
	QueryAgent
}

func (this *QueryAgent) WithKeys(key string) *QueryAgent {
	this.superAgent.QueryData.Add("keys", key)
	return this
}

func (this *QueryAgent) WithCount() *QueryAgent {
	this.superAgent.QueryData.Add("count", "1")
	return this
}

func (this *QueryAgent) WithQuery(q *query.Query) *QueryAgent {
	this.superAgent.QueryData.Add("where", q.String())
	return this
}

func (this *QueryAgent) WithCql(cql string) *QueryAgent {
	this.superAgent.QueryData.Add("cql", cql)
	return this
}

func (this *QueryAgent) Limit(l int) *QueryAgent {
	this.superAgent.QueryData.Add("limit", fmt.Sprintf("%d", l))
	return this
}

func (this *QueryAgent) Skip(s int) *QueryAgent {
	this.superAgent.QueryData.Add("skip", fmt.Sprintf("%d", s))
	return this
}

func (this *QueryAgent) Order(s string) *QueryAgent {
	this.superAgent.QueryData.Add("order", fmt.Sprintf("%d", s))
	return this
}

func (this *Agent) ScanResponse(ret interface{}) error {
	if err := json.Unmarshal([]byte(this.body), ret); nil != err {
		println(this.body)
		return err
	}
	return nil
}

//if you wanna use signature instead of key directlly , use this
func (this *Agent) UseSignature() *Agent {
	this.useSignature = true
	this.ts = time.Now().UnixNano() / 1000
	return this
}

//if you wanna use master key, use it
func (this *Agent) UseMasterKey() *Agent {
	this.useMasterKey = true
	return this
}

func (this *Agent) Do() error {
	if this.useSignature {
		this.superAgent = this.superAgent.
			Set("X-LC-Sign", this.getSignature())
	} else if this.useMasterKey {
		this.superAgent = this.superAgent.
			Set("X-LC-Key", this.client.masterKey+",master")
	} else {
		this.superAgent = this.superAgent.
			Set("X-LC-Key", this.client.appKey)
	}
	resp, body, err := this.superAgent.End()
	this.body = body

	if resp.StatusCode >= 400 {
		return errors.New(body)
	}
	if len(err) != 0 {
		for i := range err {
			println(err[i].Error())
		}
		println("response bod:" + body)
		return err[0]
	}
	return nil
}

func (this Agent) getSignature() string {

	toSign := ""
	append := ""
	if this.useMasterKey {
		toSign = fmt.Sprintf("%d%s", this.ts, this.client.masterKey)
		append = ",master"
	} else {
		toSign = fmt.Sprintf("%d%s", this.ts, this.client.appKey)
	}

	sign := fmt.Sprintf("%x", md5.Sum([]byte(toSign)))
	return fmt.Sprintf("%s,%d%s", sign, this.ts, append)
}
