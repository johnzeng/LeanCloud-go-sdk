package lean

type User struct {
	LeanClassesBase
	Salt                string                 `json:"salt,omitempty"`
	Email               string                 `json:"email, omitempt"`
	SessionToken        string                 `json:"sessionToken, omitempt"`
	Passowrd            string                 `json:"password,omitempt"`
	Username            string                 `json:"username,omitempt"`
	EmailVerified       bool                   `json:"emailVerified, omitempt"`
	MobilePhoneNumber   string                 `json:"mobilePhoneNumber, omitempt"`
	AuthData            map[string]interface{} `json:"authData, omitempt"`
	MobilePhoneVerified bool                   `json:"mobilePhoneVerified, omitempty"`
}
