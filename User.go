package lean

import (
	"github.com/parnurzeal/gorequest"
)

//be attention that EmailVerified and MobilePhoneVerified can be nil
type User struct {
	LeanClassesBase
	Salt                string                 `json:"salt,omitemptyy"`
	Email               string                 `json:"email, omitempty"`
	SessionToken        string                 `json:"sessionToken, omitempty"`
	Passowrd            string                 `json:"password,omitempty"`
	Username            string                 `json:"username,omitempty"`
	EmailVerified       *bool                  `json:"emailVerified, omitempty"`
	MobilePhoneNumber   string                 `json:"mobilePhoneNumber, omitempty"`
	AuthData            map[string]interface{} `json:"authData, omitempty"`
	MobilePhoneVerified *bool                  `json:"mobilePhoneVerified, omitempty"`
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

//will return nil if there are any error
func (this *leanClient) UserMe(token string) (*User, error) {
	url := UrlBase + "/users/me"
	request := gorequest.New()
	superAgent := request.Get(url)
	agent := &Agent{
		superAgent: superAgent,
		client:     this,
	}
	agent.UseSessionToken(token)

	if err := agent.Do(); nil != err {
		return nil, err
	}
	ret := &User{}
	if err := agent.ScanResponse(ret); nil != err {
		return nil, err
	}
	return ret, nil
}
