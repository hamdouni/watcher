/*
	Use package sabhiram/go-gitignore and add the capacity to read the ignore patterns from a file.
	The pattern must be initalized with one of the 2 functions available :
		- Init with a pattern string as parameter
		- InitFromFile with the full path to a git ignore file.
	After initializing the pattern, use IsIgnoredFile to see if a file is ignored or not.
*/
package ignore

import (
	"bufio"
	"os"

	gitignore "github.com/sabhiram/go-gitignore"
)

var ignoreObject *gitignore.GitIgnore

// IgnoredFile return true it the file is ignored
func IgnoredFile(path string) bool {
	return ignoreObject.MatchesPath(path)
}

// Init initialize the ignore pattern
func Init(pattern []string) {
	ignoreObject = gitignore.CompileIgnoreLines(pattern...)
}

// InitFromFile read one or more gitignore files and intialize the ignore pattern
func InitFromFile(paths ...string) error {
	var pattern []string
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			pattern = append(pattern, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	Init(pattern)

	return nil
}
