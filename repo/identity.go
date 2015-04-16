package repo

type Identity struct {
	Name  string
	Email string
}

func (i *Identity) hasName() bool {
	return i.Name != ""
}

func (i *Identity) hasEmail() bool {
	return i.Email != ""
}

func (i *Identity) notEqual(j *Identity) bool {
	return (i.hasName() && i.Name != j.Name) ||
		(i.hasEmail() && i.Email != j.Email)
}

func (i *Identity) String() string {
	str := ""
	if i.hasName() {
		str += "  Name: " + i.Name + "\n"
	}
	if i.hasEmail() {
		str += "  Email: " + i.Email
	}
	return str
}
