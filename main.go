/*
	Watcher is a Go live reload program for Linux.
	It monitors the actual directory and all its subfolders for write or delete events on any file except those matching any pattern in .gitignore file.
*/
package main

import (
	_ "embed"
	"log"
	"watcher/command"
	"watcher/ignore"
	"watcher/monitor"
)

//go:embed VERSION
var version string

func main() {

	var err error

	log.Printf("Watcher version %v", version)

	// Use .gitignore if it exists or use an empty pattern
	err = ignore.InitFromFile(".gitignore", ".watcherignore")
	if err != nil {
		ignore.Init([]string{""})
	}

	ch, err := monitor.Watch()
	if err != nil {
		log.Fatal(err)
	}
	defer monitor.Stop(ch)

	err = command.Launch()
	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := <-ch
		if !ignore.IgnoredFile(ev.Path()) {
			log.Printf("Live reload on event %v\n", ev)
			err = command.Kill()
			if err != nil {
				log.Printf("%v\n", err)
			}
			err = command.Launch()
			if err != nil {
				log.Printf("%v\n", err)
			}
		}
	}
}
