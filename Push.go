package lean

import (
	"time"
)

type IOSPushData struct {
	Data struct {
		Alter            string
		Category         string
		Badge            int
		Sound            string
		ContentAvailable int
		MutableContent   int
		Custom           interface{}
	} `json:data`
}

type Push struct {
	LeanClassesBase
	Channels           []string    `json:"channels"`
	Data               interface{} `json:"data"`
	ExpirationInterval int         `json:"expirationInterval"`
	ExpirationTime     int         `json:"expirationTime"`
	Prod               string      `json:"prod"`
	PushTime           time.Time   `json:"pushTime"`
	Where              interface{} `json:"where"`
}

func (this *Push) GetClassName() string {
	return "push"
}
