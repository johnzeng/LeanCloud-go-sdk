package lean

import (
	"time"
)

type IOSLocalAlert struct {
	Title          string   `json:"title,omitempty"`
	TitleLocKey    string   `json:"title-loc-key,omitempty"`
	SubTitle       string   `json:"sub-title,omitempty"`
	SubTitleLocKey string   `json:"sub=title-loc-key,omitempty"`
	Body           string   `json:"body,omitempty"`
	ActionLocKey   string   `json:"action-loc-key,omitempty"`
	LocKey         string   `json:"loc-key,omitempty"`
	LocArgs        []string `json:"loc-args,omitempty"`
	LauchImage     string   `json:"lauch-image,omitempty"`
}

//You can still add more keys into this struct using this struct as a anonymous field
type IOSData struct {
	//Alert can be eiterh a string or a IOSLocalAlert
	Alert    interface{} `json:"alter"`
	Category string      `json:"category,omitempty"`
	//Badge can be either a int or a "Increment"
	Badge            interface{} `json:"badge,omitempty"`
	Sound            string      `json:"sound,omitempty"`
	ContentAvailable int         `json:"content-available,omitempty"`
	MutableContent   int         `json:"mutable-content,omitempty"`
}

//You can still add more keys into this struct using this struct as a anonymous field
type IOSStandardData struct {
	//Alert in IOSData should be a IOSLocalAlert
	Aps IOSData `json:"aps"`
}

//You can still add more keys into this struct using this struct as a anonymous field
type AndroidData struct {
	Alert  string `json:"alert"`
	Title  string `json:"title"`
	Silent bool   `json:"silent,omitempty"`
	//this should be empty unless you have a custom Receiver
	Action string `json:"action,omitempty"`
}

//don't add custome key because wp doesn't support it
type WinPhoneData struct {
	Alert   string `json:"alert"`
	Title   string `json:"title"`
	WpParam string `json:"wp-param"`
}

type ComplexData struct {
	//the interface should be a IOSData or any struct with an anonymous file of IOSData
	IOS interface{} `json:"ios,omitempty"`
	//the interface should be a AndroidData or any struct with an anonymous file of AndroidData
	Android AndroidData  `json:"android,omitempty"`
	WP      WinPhoneData `json:"wp,omitempty"`
}

type Push struct {
	LeanClassesBase
	Channels           []string    `json:"channels,omitempty"`
	Data               interface{} `json:"data"`
	ExpirationInterval int         `json:"expirationInterval,omitempty"`
	ExpirationTime     int         `json:"expirationTime,omitempty"`
	Prod               string      `json:"prod,omitempty"`
	PushTime           time.Time   `json:"pushTime,omitempty"`
	Where              interface{} `json:"where,omitempty"`
}
