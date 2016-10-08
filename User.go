package lean

import (
	"github.com/parnurzeal/gorequest"
)

//be attention that EmailVerified and MobilePhoneVerified can be nil
type User struct {
	LeanClassesBase
	Salt                string                 `json:"salt,omitempty"`
	Email               string                 `json:"email, omitempt"`
	SessionToken        string                 `json:"sessionToken, omitempt"`
	Passowrd            string                 `json:"password,omitempt"`
	Username            string                 `json:"username,omitempt"`
	EmailVerified       *bool                  `json:"emailVerified, omitempt"`
	MobilePhoneNumber   string                 `json:"mobilePhoneNumber, omitempt"`
	AuthData            map[string]interface{} `json:"authData, omitempt"`
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
