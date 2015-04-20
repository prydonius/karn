package repo

import (
	"errors"
	"os/exec"
	"strings"
)

const (
	FAILED_UPDATE_ERR = "Failed setting new identity."
)

func IsInsideWorkTree() bool {
	if err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Run(); err != nil {
		return false
	}

	return true
}

func CurrentIdentity() (*Identity, error) {
	name, err := GetConfig("user.name")
	if err != nil {
		return nil, err
	}

	email, err := GetConfig("user.email")
	if err != nil {
		return nil, err
	}

	current := &Identity{}
	current.Name = strings.TrimSpace(name)
	current.Email = strings.TrimSpace(email)

	return current, nil
}

func GetConfig(key string) (string, error) {
	value, err := exec.Command("git", "config", key).Output()
	return string(value), err
}

func SetConfig(key string, value string) error {
	return exec.Command("git", "config", "--local", key, value).Run()
}

func UpdateIdentity(id *Identity) (bool, error) {
	current, err := CurrentIdentity()
	if err != nil {
		return false, errors.New("Couldn't retrieve identity from Git repository.")
	}

	updated := false
	if id.NameNotEqual(current.Name) {
		if err := SetConfig("user.name", id.Name); err != nil {
			return false, errors.New(FAILED_UPDATE_ERR)
		}
		updated = true
	}

	if id.EmailNotEqual(current.Email) {
		if err := SetConfig("user.email", id.Email); err != nil {
			return false, errors.New(FAILED_UPDATE_ERR)
		}
		updated = true
	}

	return updated, nil
}
