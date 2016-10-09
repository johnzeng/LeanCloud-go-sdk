package lean

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"os"
)

//be attention that EmailVerified and MobilePhoneVerified can be nil
type FullLeanFile struct {
	LeanClassesBase
	Name     string                 `json:"name,omitempty"`
	FileType string                 `json:"mime_type,omitempty"`
	Key      string                 `json:"key,omitempty"`
	Type     string                 `json:"__type,omitempty"`
	Url      string                 `json:"url,omitempty"`
	Bucket   string                 `json:"bucket,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
}

func (this *leanClient) UploadPlainText(fileName, content string) (*FullLeanFile, error) {
	url := UrlBase + "/files/" + fileName
	request := gorequest.New()
	superAgent := request.Post(url).
		Send(content)

	superAgent.Set("Content-Type", "text/plain")
	agent := &Agent{
		superAgent: superAgent,
		client:     this,
	}

	if err := agent.Do(); nil != err {
		return nil, err
	}

	ret := &FullLeanFile{}

	if err := agent.ScanResponse(ret); nil != err {
		return nil, err
	}

	return ret, nil
}

//content-type can be text/plain,image/*, or empty
func (this *leanClient) UploadFile(fileName, contentType string, file *os.File) (*FullLeanFile, error) {
	//as gorequest doesn't have a good support for binary request, so I have to make the request by native http.request api,this is bad
	url := UrlBase + "/files/" + fileName

	if "" != contentType {
		contentType = "application/json"
	}
	httpRequest, err := http.NewRequest("POST", url, file)
	if nil != err {
		return nil, err
	}
	httpRequest.Header.Add("X-LC-Key", this.appKey)
	httpRequest.Header.Add("X-LC-Id", this.appId)
	client := http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if nil != err {
		return nil, err
	}

	ret := &FullLeanFile{}
	retBytes, err := ioutil.ReadAll(httpResponse.Body)
	if nil != err {
		return nil, err
	}
	if jsonErr := json.Unmarshal(retBytes, ret); nil != jsonErr {
		return nil, jsonErr
	}

	return ret, nil
}
