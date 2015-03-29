package main

import (
	"errors"
	"os/exec"
	"strings"
)

type repo struct{}

func (r *repo) IsInsideWorkTree() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}

func (r *repo) CurrentIdentity() (*identity, error) {
	name, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		return nil, err
	}

	email, err := exec.Command("git", "config", "user.email").Output()
	if err != nil {
		return nil, err
	}

	current := new(identity)
	current.Name = strings.TrimSpace(string(name))
	current.Email = strings.TrimSpace(string(email))

	return current, nil
}

func (r *repo) SetIdentity(id *identity) error {
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

func (r *repo) UpdateIdentity(id *identity) (bool, error) {
	current, err := r.CurrentIdentity()
	if err != nil {
		return false, errors.New("Couldn't retrieve identity from Git repository.")
	}

	if (id.hasName() && id.Name != current.Name) ||
		(id.hasEmail() && id.Email != current.Email) {
		err = r.SetIdentity(id)
		if err != nil {
			return false, errors.New("Failed setting new identity.")
		} else {
			return true, nil
		}
	}
	return false, nil
}
