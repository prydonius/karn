package main

type identity struct {
	Name  string
	Email string
}

func (i *identity) hasName() bool {
	return i.Name != ""
}

func (i *identity) hasEmail() bool {
	return i.Email != ""
}

func (i *identity) String() string {
	str := ""
	if i.hasName() {
		str += "\tName: " + i.Name
	}
	if i.hasEmail() {
		str += "\tEmail: " + i.Email
	}
	return str
}
