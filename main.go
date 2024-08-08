/*
Watcher is a Go live reload program for Linux.
It monitors the actual directory and all its subfolders for write or delete events
on any file except those matching any pattern in .gitignore or .watcherignore file.
*/
package main

import (
	_ "embed"
	"flag"
	"log"

	"github.com/sipkg/watcher/command"
	"github.com/sipkg/watcher/ignore"
	"github.com/sipkg/watcher/monitor"
)

//go:embed VERSION
var version string

func main() {
	var (
		srcpath = flag.String("dir", ".", "Directory to watch")
		program = flag.String("run", "", "Program to run")
		test    = flag.Bool("test", false, "Run tests")
		quiet   = flag.Bool("quiet", false, "Log only errors")
		help    = flag.Bool("help", false, "Command line usage")
		args    = flag.String("args", "", "Args to pass surrounded with quotes")
	)
	flag.Parse()

	if !*quiet {
		log.Printf("Watcher version %v", version)
	}

	if *help {
		flag.Usage()
		return
	}

	// Use .gitignore or .watcherignore if they exist.
	// Otherwise use an empty pattern.
	if err := ignore.Read(".gitignore", ".watcherignore"); err != nil {
		log.Printf("reading ignore files: %s\nusing empty pattern", err)
		ignore.New([]string{""})
	}

	// Launch file monitoring
	ch, err := monitor.Watch(*srcpath)
	if err != nil {
		log.Fatalf("launching monitor: %s", err)
	}
	defer monitor.Stop(ch)

	// Launch command
	launch := command.Launch
	if *test {
		launch = command.Test
	}
	err = launch(*program, *args)
	if err != nil {
		log.Fatalf("launching program: %s", err)
	}

	// Main loop reacts to file events if not ignored
	for {
		ev := <-ch
		if !ignore.Ignored(ev.Path()) {
			if !*quiet {
				log.Printf("Live reload on event %v\n", ev)
			}
			if !*test {
				err = command.Kill()
				if err != nil {
					log.Printf("%v\n", err)
				}
			}
			if err := launch(*program, *args); err != nil {
				log.Printf("%v\n", err)
			}
		}
	}
}
