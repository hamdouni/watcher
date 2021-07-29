/*
	Watcher is a Go live reload program for Linux.
	It monitors the actual directory and all its subfolders for write or delete events on any file except those matching any pattern in .gitignore file.
*/
package main

import (
	_ "embed"
	"log"
	"os"
	"watcher/command"
	"watcher/ignore"
	"watcher/monitor"
)

//go:embed VERSION
var version string

func main() {

	var err error

	log.Printf("Watcher version %v", version)
	if len(os.Args) > 2 {
		log.Fatalf("watcher accepts one argument, the source path, but %v arguments provided", len(os.Args)-1)
	}

	// if a source path is passed as argument, use it instead the default one
	srcpath := "."
	if len(os.Args) == 2 {
		srcpath = os.Args[1]
	}

	// Use .gitignore if it exists or use an empty pattern
	err = ignore.InitFromFile(".gitignore", ".watcherignore")
	if err != nil {
		ignore.Init([]string{""})
	}

	ch, err := monitor.Watch(srcpath)
	if err != nil {
		log.Fatal(err)
	}
	defer monitor.Stop(ch)

	err = command.Launch(srcpath)
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
			err = command.Launch(srcpath)
			if err != nil {
				log.Printf("%v\n", err)
			}
		}
	}
}
