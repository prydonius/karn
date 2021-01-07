package repo

type Identity struct {
	Name       string
	Email      string
	SigningKey string
}

func (i *Identity) hasName() bool {
	return i.Name != ""
}

func (i *Identity) hasEmail() bool {
	return i.Email != ""
}

func (i *Identity) hasSigningKey() bool {
	return i.SigningKey != ""
}

func (i *Identity) NameNotEqual(name string) bool {
	return i.hasName() && i.Name != name
}

func (i *Identity) EmailNotEqual(email string) bool {
	return i.hasEmail() && i.Email != email
}

func (i *Identity) SigningKeyNotEqual(signingKey string) bool {
	return i.hasSigningKey() && i.SigningKey != signingKey
}

func (i *Identity) String() string {
	str := ""
	if i.hasName() {
		str += "  Name: " + i.Name + "\n"
	}
	if i.hasEmail() {
		str += "  Email: " + i.Email + "\n"
	}
	if i.hasSigningKey() {
		str += "  Signing Key: " + i.SigningKey
	}
	return str
}
