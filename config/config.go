package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"regexp"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/prydonius/karn/repo"
	yaml "gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE   = "~/.karn.yml"
	FILE_READ_ERR = "Unable to open karn configuration file. Did you create a ~/.karn.yml in your" +
		" home directory?"
)

type Dirs map[string]*repo.Identity

func GetConfig() (Dirs, error) {
	file, err := homedir.Expand(CONFIG_FILE)
	if err != nil {
		return nil, err
	}

	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New(FILE_READ_ERR)
	}

	dirs := make(Dirs)
	err = yaml.Unmarshal(source, dirs)
	if err != nil {
		return dirs, err
	}

	return dirs, nil
}

func GetIdentity(path string, dirs Dirs) (*repo.Identity, error) {
	for {
		for dir, _ := range dirs {
			verdict, err := regexp.MatchString(path+"/?", dir)
			if err != nil {
				return nil, err
			}

			if verdict {
				return dirs[dir], nil
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
