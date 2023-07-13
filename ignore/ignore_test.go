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
			desc:    "Passant",
			patt:    []string{"*.go"},
			file:    "main.go",
			ignored: true,
		},
		{
			desc:    "Passant",
			patt:    []string{"*.go"},
			file:    "main.c",
			ignored: false,
		},
		{
			desc:    "Passant",
			patt:    []string{"*.go"},
			file:    "/home/test/main.go",
			ignored: true,
		},
		{
			desc:    "Passant",
			patt:    []string{"test/*.go"},
			file:    "main.c",
			ignored: false,
		},
		{
			desc:    "Passant",
			patt:    []string{"main.c"},
			file:    "/home/test/pass/main.c",
			ignored: false,
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
