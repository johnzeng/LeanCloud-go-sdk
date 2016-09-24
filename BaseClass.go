package lean

import (
	"time"
)

type ACLMap map[string]bool

type LeanClassesBase struct {
	ObjectId  string            `json:"objectId"`
	CreatedAt time.Time         `json:"createdAt"`
	updatedAt time.Time         `json:"updatedAt"`
	ACL       map[string]ACLMap `json:ACL`
}
