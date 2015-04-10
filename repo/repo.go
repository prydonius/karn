package repo

import (
	"errors"
	"os/exec"
	"strings"
)

func IsInsideWorkTree() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}

func CurrentIdentity() (*Identity, error) {
	name, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		return nil, err
	}

	email, err := exec.Command("git", "config", "user.email").Output()
	if err != nil {
		return nil, err
	}

	current := &Identity{}
	current.Name = strings.TrimSpace(string(name))
	current.Email = strings.TrimSpace(string(email))

	return current, nil
}

func SetIdentity(id *Identity) error {
	if id.hasName() {
		err := exec.Command("git", "config", "--local", "user.name", id.Name).Run()
		if err != nil {
			return err
		}
	}

	if id.hasEmail() {
		err := exec.Command("git", "config", "--local", "user.email", id.Email).Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateIdentity(id *Identity) (bool, error) {
	current, err := CurrentIdentity()
	if err != nil {
		return false, errors.New("Couldn't retrieve identity from Git repository.")
	}

	if (id.hasName() && id.Name != current.Name) ||
		(id.hasEmail() && id.Email != current.Email) {
		err = SetIdentity(id)
		if err != nil {
			return false, errors.New("Failed setting new identity.")
		} else {
			return true, nil
		}
	}
	return false, nil
}
