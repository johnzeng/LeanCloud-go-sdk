package lean

import (
	"github.com/parnurzeal/gorequest"
)

type User struct {
	LeanClassesBase
	Salt                string                 `json:"salt,omitempty"`
	Email               string                 `json:"email, omitempt"`
	SessionToken        string                 `json:"sessionToken, omitempt"`
	Passowrd            string                 `json:"password,omitempt"`
	Username            string                 `json:"username,omitempt"`
	EmailVerified       bool                   `json:"emailVerified, omitempt"`
	MobilePhoneNumber   string                 `json:"mobilePhoneNumber, omitempt"`
	AuthData            map[string]interface{} `json:"authData, omitempt"`
	MobilePhoneVerified bool                   `json:"mobilePhoneVerified, omitempty"`
}

//will return nil if there are any error
func (this *leanClient) Login(userName, pwd string) (*User, error) {
	requestBody := map[string]string{
		"username": userName,
		"password": pwd,
	}
	url := UrlBase + "/login"
	request := gorequest.New()
	superAgent := request.Post(url).
		Send(requestBody)
	agent := &Agent{
		superAgent: superAgent,
		client:     this,
	}
	if err := agent.Do(); nil != err {
		return nil, err
	}
	ret := &User{}
	if err := agent.ScanResponse(ret); nil != err {
		return nil, err
	}
	return ret, nil
}
