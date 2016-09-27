package lean

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type ACLMap map[string]bool

type LeanClassesBase struct {
	ObjectId  string            `json:"objectId,omitempty"`
	CreatedAt *LeanTime         `json:"createdAt,omitempty"`
	UpdatedAt *LeanTime         `json:"updatedAt,omitempty"`
	ACL       map[string]ACLMap `json:ACL,omitempty`
}

type LeanTime time.Time

func NewLeanTime(t time.Time) LeanTime {
	return LeanTime(t)
}

func (this *LeanTime) GetTime() *time.Time {
	return (*time.Time)(this)
}

//2015-07-14T02:31:50.100Z
//convet TimeStamp into json, we take only the unix timestamp seconds
func (t LeanTime) MarshalJSON() ([]byte, error) {
	stamp := t.GetTime().UTC().Format("2006-01-02T15:04:05.999Z")
	println(stamp)
	str := fmt.Sprintf(`{
		"__type": "Date",
		"iso": "%s"
	}`, stamp)
	return []byte(str), nil
}

func (t *LeanTime) typeName() string {
	return "Date"
}

func (t *LeanTime) fillByMap(m map[string]string) error {
	timeStr := m["iso"]
	if ret, err := str2Date(timeStr); nil != err {
		return err
	} else {
		t = ret
	}
	return nil
}

//convet json into TimeStamp, we take only the unix timestamp seconds
func (t *LeanTime) UnmarshalJSON(i []byte) error {
	//do your serializing here
	var timeStr string
	if err := json.Unmarshal(i, &timeStr); err == nil {
		if t, err := str2Date(timeStr); err != nil {
			if converErr := str2AvObject(timeStr, t); nil != converErr {
				return converErr
			}
		}
		return nil
	} else {
		return err
	}
}

func str2Date(str string) (*LeanTime, error) {
	if ret, err := time.Parse("2006-01-02T15:04:05.000Z", str); nil != err {
		return nil, err
	} else {
		return (*LeanTime)(&ret), nil
	}
}

type AVObject interface {
	typeName() string
	fillByMap(map[string]string) error
}

func str2AvObject(str string, obj AVObject) error {
	object := map[string]string{}
	if err := json.Unmarshal([]byte(str), &object); nil != err {
		return err
	} else {
		if object["__type"] != obj.typeName() {
			return errors.New("type wrong!")
		}
		if err := obj.fillByMap(object); nil != err {
			return err
		}
		return nil
	}
}
