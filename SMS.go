package lean

import (
	"github.com/parnurzeal/gorequest"
)

const (
	VoiceSMS = "void"
	TextSMS  = "sms"
)

//Number is the phone number
//TTL is default to 10 minutes
//AppName is default to the app name of you lean cloud application
//Op is default to verifying
//SmsType is default to TextSMS,
//set SmsType to VoiceSMS if you wanna send a voice sms
type RequestMobilePhoneVerify struct {
	Number  string `json:"mobilePhoneNumber"`
	TTL     int    `json:"ttl,omitempty"`
	AppName string `json:"name,omitempty"`
	Op      string `json:"op,omitempty"`
	SmsType string `json:"smsType,omitempty"`
}

//request a mobile phone verify
func (client leanClient) RequestMobilVerify(
	verifyRequest RequestMobilePhoneVerify) error {

	request := gorequest.New()
	url := UrlBase + "/requestSmsCode"
	superAgent := request.Post(url).
		Set("X-LC-Id", client.appId).
		Send(verifyRequest)
	agent := Agent{
		superAgent: superAgent,
		client:     client,
	}

	err := agent.Do()
	if nil != err {
		return err
	} else {
		return nil
	}
}

//verfiy the code
func (client leanClient) VerifyCode(phone, code string) error {
	request := gorequest.New()
	url := UrlBase + "/verifySmsCode/" + code
	superAgent := request.Post(url).
		Set("X-LC-Id", client.appId).
		Query("mobilePhoneNumber=" + phone)

	agent := Agent{
		superAgent: superAgent,
		client:     client,
	}

	err := agent.UseSignature().Do()
	if nil != err {
		return err
	} else {
		return nil
	}
}
