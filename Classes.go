package lean

import (
	"github.com/parnurzeal/gorequest"
)

type Collection struct {
	client leanClient
	class  string
}

func (this leanClient) Collection(collection string) Collection {
	return Collection{
		client: this,
		class:  collection,
	}
}

//create an object
func (this Collection) Create(obj interface{}) *Agent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + this.class
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
	classesUrl := ClasssesUrlBase + "/" + this.class + "/" + objectId
	superAgent := request.Get(classesUrl)
	return &Agent{
		superAgent: superAgent,
		client:     this.client,
	}
}

//you can also specialfy the query parameter, if you don't provide the id, you will delete the objects by query
func (this Collection) DeleteObjectById(objectId string) *QueryAgent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + this.class
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
	classesUrl := ClasssesUrlBase + "/" + this.class
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

func (this Collection) Query() *QueryAgent {
	request := gorequest.New()
	classesUrl := ClasssesUrlBase + "/" + this.class
	superAgent := request.Get(classesUrl)
	agent := Agent{
		superAgent: superAgent,
		client:     this.client,
	}
	return &QueryAgent{agent}
}
