package karn

import (
	"fmt"
	"log"
	"os"

	"github.com/prydonius/karn/config"
	"github.com/prydonius/karn/repo"
)

func Update() {
	if !repo.IsInsideWorkTree() {
		log.Fatal("Not inside Git work tree")
	}

	dirs, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}

	id, err := config.GetIdentity(cwd, dirs)
	if err != nil {
		log.Fatal(err)
	}

	// Don't bother updating if we don't have an identity for the current working dir
	if id == nil {
		return
	}

	updated, err := repo.UpdateIdentity(id)

	if err != nil {
		log.Fatal(err)
	}

	if updated {
		fmt.Println("### Updated Git Identity, you are now commiting as:")
		fmt.Println(id)
	}
}

func Init() {
	fmt.Printf(`git() {
  karn update
  command git "$@"
}`)
}

func Install() {
	fmt.Printf("To setup karn to check identity updates automatically before running any Git commands," +
		" add the following line to your shell startup file:\n" +
		"\tif which karn > /dev/null; then eval \"$(karn init)\"; fi")
}
