package lean

import (
	"time"
)

type ACLMap map[string]bool

type LeanClassesBase struct {
	ObjectId  string            `json:"objectId,omitempty"`
	CreatedAt time.Time         `json:"createdAt,omitempty"`
	updatedAt time.Time         `json:"updatedAt,omitempty"`
	ACL       map[string]ACLMap `json:ACL,omitempty`
}
