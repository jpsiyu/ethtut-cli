package main

import "github.com/jpsiyu/ethtut-cli/cmd"

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	cmd.Execute()
}
