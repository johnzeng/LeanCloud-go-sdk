package lean

import (
	"github.com/parnurzeal/gorequest"
)

type LeanClasses interface {
	GetClassName() string
}

//create an object
func (client leanClient) Create(obj LeanClasses) *Agent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + obj.GetClassName()
	superAgent := request.Post(classesUrl).
		Set("X-LC-Id", client.appId).
		Send(obj)
	return &Agent{
		superAgent: superAgent,
		client:     client,
	}
}

func (client leanClient) GetObjectById(class, objectId string) *Agent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + class + "/" + objectId
	superAgent := request.Get(classesUrl).
		Set("X-LC-Id", client.appId)
	return &Agent{
		superAgent: superAgent,
		client:     client,
	}
}
