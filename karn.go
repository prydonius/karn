package main

import (
	"fmt"
	"log"
)

type karn struct{}

func (k *karn) Update() {
	repo := new(repo)
	if !repo.IsInsideWorkTree() {
		log.Fatal("Not inside Git work tree")
	}

	conf := new(config)
	id, err := conf.ConfiguredIdentity()
	if err != nil {
		log.Fatal(err)
	}
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

func (k *karn) Init() {
	fmt.Printf(`git() {
  karn update
  command git $@
}`)
}

func (k *karn) Install() {
	fmt.Printf("To setup karn to check identity updates automatically before running any Git commands," +
		" add the following line to your shell startup file:\n" +
		"\tif which karn > /dev/null; then eval \"$(karn init)\"; fi")
}

func main() {
	log.SetFlags(0)
	karn := new(karn)
	cli := &karnCli{karn}
	cli.init()
}
