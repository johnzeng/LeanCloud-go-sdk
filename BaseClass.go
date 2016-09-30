package lean

type ACLMap map[string]bool

type LeanClassesBase struct {
	ObjectId  string             `json:"objectId,omitempty"`
	CreatedAt *LeanTime          `json:"createdAt,omitempty"`
	UpdatedAt *LeanTime          `json:"updatedAt,omitempty"`
	ACL       *map[string]ACLMap `json:ACL,omitempty`
}
