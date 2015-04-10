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

func (i *Identity) String() string {
	str := ""
	if i.hasName() {
		str += "\tName: " + i.Name
	}
	if i.hasEmail() {
		str += "\tEmail: " + i.Email
	}
	return str
}
