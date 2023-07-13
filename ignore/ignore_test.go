package ignore_test

import (
	"testing"

	"github.com/sipkg/watcher/ignore"
)

func TestIgnore(t *testing.T) {
	testCases := []struct {
		desc    string
		patt    []string
		file    string
		ignored bool
	}{
		{
			desc:    "Unique pattern and matching file in root",
			patt:    []string{"*.go"},
			file:    "main.go",
			ignored: true,
		},
		{
			desc:    "Unique pattern and not matching file in root",
			patt:    []string{"*.go"},
			file:    "main.c",
			ignored: false,
		},
		{
			desc:    "Unique pattern and matching file deep in path",
			patt:    []string{"*.go"},
			file:    "/home/test/main.go",
			ignored: true,
		},
		{
			desc:    "Pattern with path and file in root",
			patt:    []string{"test/*.go"},
			file:    "main.c",
			ignored: false,
		},
		{
			desc:    "Precise pattern file and matching file deep in path",
			patt:    []string{"main.c"},
			file:    "/home/test/pass/main.c",
			ignored: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			ignore.New(tC.patt)
			result := ignore.Ignored(tC.file)
			if result != tC.ignored {
				t.Errorf("wait %v got %v", tC.ignored, result)
			}
		})
	}
}
