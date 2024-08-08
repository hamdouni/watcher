/*
Use package sabhiram/go-gitignore and add the capacity to read the ignore
patterns from files.
The pattern must be initalized with one of the 2 functions available :
  - New with a pattern string as parameter
  - Read with any number of paths to ignore files.

After initializing the pattern, use Ignored to see if a file is ignored or not.
*/
package ignore

import (
	"bufio"
	"os"

	gitignore "github.com/sabhiram/go-gitignore"
)

var ignoreObject *gitignore.GitIgnore

// Ignored returns true it the file is ignored
func Ignored(path string) bool {
	return ignoreObject.MatchesPath(path)
}

// New initializes the ignore pattern
func New(pattern []string) {
	ignoreObject = gitignore.CompileIgnoreLines(pattern...)
}

// Read reads one or more ignore files and initialize the ignore pattern
func Read(paths ...string) error {
	var pattern []string
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil && os.IsNotExist(err) {
			continue
		} else if err != nil {
			return err
		}
		defer file.Close()
		// we read the files line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			pattern = append(pattern, scanner.Text())
		}
	}
	New(pattern)
	return nil
}
