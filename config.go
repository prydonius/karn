package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE   = "~/.karn.yml"
	FILE_READ_ERR = "Unable to open karn configuration file. Did you create a ~/.karn.yml in your" +
		" home directory?"
)

type config struct {
	Dirs map[string]*identity
}

func (c *config) LoadConfiguration() error {
	file, err := homedir.Expand(CONFIG_FILE)
	if err != nil {
		return err
	}

	source, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New(FILE_READ_ERR)
	}

	err = yaml.Unmarshal(source, &c.Dirs)
	if err != nil {
		return err
	}

	return nil
}

func (c *config) ConfiguredIdentity() (*identity, error) {
	err := c.LoadConfiguration()
	if err != nil {
		return nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	id, err := c.IdentityInDir(cwd)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (c *config) IdentityInDir(path string) (*identity, error) {
	for {
		for dir, _ := range c.Dirs {
			verdict, err := regexp.MatchString(path+"/?", dir)
			if err != nil {
				return nil, err
			}

			if verdict {
				return c.Dirs[dir], nil
			}
		}

		path, _ = filepath.Split(path)
		length := len(path)

		if length == 1 {
			break
		}

		// Remove trailing slash
		path = path[:length-1]
	}

	return nil, nil
}
