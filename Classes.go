package lean

import (
	"github.com/parnurzeal/gorequest"
)

//create an object
func (client leanClient) Create(class string, obj interface{}) *Agent {
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

//you can also specialfy the query parameter, if you don't provide the id, you will delete the objects by query
func (client leanClient) DeleteObjectById(class, objectId string) *QueryAgent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + class
	if "" != objectId {
		classesUrl = classesUrl + "/" + objectId
	}
	superAgent := request.Delete(classesUrl).
		Set("X-LC-Id", client.appId)
	return &Agent{
		superAgent: superAgent,
		client:     client,
	}
}

//you can also specialfy the query parameter, if you don't provide the id, you will update the object by query
func (client leanClient) UpdateObjectById(class, id string, obj interface{}) *UpdateAgent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + obj.GetClassName()
	if "" != objectId {
		classesUrl = classesUrl + "/" + objectId
	}
	superAgent := request.Put(classesUrl).
		Set("X-LC-Id", client.appId).
		Send(obj)
	return &Agent{
		superAgent: superAgent,
		client:     client,
	}
}

func (client leanClient) Query(class string) *QueryAgent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + class
	superAgent := request.Get(classesUrl).
		Set("X-LC-Id", client.appId)
	agent := Agent{
		superAgent: superAgent,
		client:     client,
	}
	return &QueryAgent{agent}
}
