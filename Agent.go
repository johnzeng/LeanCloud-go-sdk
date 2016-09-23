package lean

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type Agent struct {
	client       leanClient
	superAgent   *gorequest.SuperAgent
	useSignature bool
	useMasterKey bool
	ts           int64
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
	println(body)

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
