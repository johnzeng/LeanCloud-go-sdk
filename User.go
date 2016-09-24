package lean

type User struct {
	LeanClassesBase
}

func (this *User) GetClassName() string {
	return "user"
}
