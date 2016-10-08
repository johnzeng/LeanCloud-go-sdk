package lean

import (
	"github.com/parnurzeal/gorequest"
	"os"
)

//be attention that EmailVerified and MobilePhoneVerified can be nil
type File struct {
	LeanClassesBase
	Name   string `json:"name,omitempt"`
	Type   string `json:"__type,omitempt"`
	Url    string `json:"url,omitempt"`
	Bucket string `json:"bucket, omitempt"`
}

//content-type can be text/plain,image/*, or empty
func (this *leanClient) UploadFile(fileName, contentType string, file *os.File) *Agent {
	url := UrlBase + "/files/" + fileName
	request := gorequest.New()
	superAgent := request.Post(url).
		Send(file)

	if "" != contentType {
		superAgent.Set("Content-Type", contentType)
	}
	agent := &Agent{
		superAgent: superAgent,
		client:     this,
	}

	return agent
}
