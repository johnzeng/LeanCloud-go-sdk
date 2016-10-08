package lean

import (
	"github.com/parnurzeal/gorequest"
)

type Collection struct {
	client      *leanClient
	class       string
	classSubfix string
}

func (this *leanClient) Collection(collection string) Collection {
	return Collection{
		client:      this,
		class:       collection,
		classSubfix: "/classes/" + collection,
	}
}

//create an object
func (this Collection) Create(obj interface{}) *Agent {
	request := gorequest.New()
	classesUrl := UrlBase + this.classSubfix
	superAgent := request.Post(classesUrl).
		Send(obj)
	return &Agent{
		superAgent: superAgent,
		client:     this.client,
	}
}

//get an object by objectId
func (this Collection) GetObjectById(objectId string) *Agent {
	request := gorequest.New()
	classesUrl := UrlBase + this.classSubfix + "/" + objectId
	superAgent := request.Get(classesUrl)
	return &Agent{
		superAgent: superAgent,
		client:     this.client,
	}
}

//you can also specialfy the query parameter, if you don't provide the id, you will delete the objects by query
func (this Collection) DeleteObjectById(objectId string) *QueryAgent {
	request := gorequest.New()
	classesUrl := UrlBase + this.classSubfix
	if "" != objectId {
		classesUrl = classesUrl + "/" + objectId
	}
	superAgent := request.Delete(classesUrl)
	return &QueryAgent{Agent{
		superAgent: superAgent,
		client:     this.client,
	}}
}

//you can also specialfy the query parameter, if you don't provide the id, you will update the object by query
func (this Collection) UpdateObjectById(objectId string, obj interface{}) *UpdateAgent {
	request := gorequest.New()
	classesUrl := UrlBase + this.classSubfix
	if "" != objectId {
		classesUrl = classesUrl + "/" + objectId
	}
	superAgent := request.Put(classesUrl).
		Send(obj)
	return &UpdateAgent{QueryAgent{Agent{
		superAgent: superAgent,
		client:     this.client,
	}}}
}

//leave cursor and key empty if they are empty
func (this Collection) Scan(cursor, key string) *QueryAgent {
	request := gorequest.New()
	classesUrl := UrlBase + "/scan" + this.classSubfix
	superAgent := request.Get(classesUrl)
	if "" != cursor {
		superAgent.QueryData.Add("cursor", cursor)
	}

	if "" != key {
		superAgent.QueryData.Add("scan_key", key)
	}
	agent := Agent{
		superAgent: superAgent,
		client:     this.client,
	}
	return &QueryAgent{agent}
}

func (this Collection) Query() *QueryAgent {
	request := gorequest.New()
	classesUrl := UrlBase + this.classSubfix
	superAgent := request.Get(classesUrl)
	agent := Agent{
		superAgent: superAgent,
		client:     this.client,
	}
	return &QueryAgent{agent}
}
