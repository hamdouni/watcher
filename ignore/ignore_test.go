package ignore_test

import (
	"os"
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

// Use also non existant files to test no border effect
func TestFileIgnore(t *testing.T) {

	thePattern := "pattern/*.ext"

	file, err := os.CreateTemp("", "ignore")
	if err != nil {
		t.Fatalf("creating temp file: %s", err)
	}
	_, err = file.WriteString(thePattern)
	if err != nil {
		t.Fatalf("writing in temp file %s: %s", file.Name(), err)
	}
	err = file.Close()
	if err != nil {
		t.Fatalf("closing temp file %s: %s", file.Name(), err)
	}

	err = ignore.Read(file.Name(), "inexistant_file1", "inexistant_file2")
	if err != nil {
		t.Fatalf("reading temp file %s: %s", file.Name(), err)
	}

	cases := []struct {
		desc    string
		file    string
		ignored bool
	}{
		{"matching", "pattern/toto.ext", true},
		{"not matching", "elsewhere/toto.ext", false},
	}
	for _, uc := range cases {
		t.Run(uc.desc, func(t *testing.T) {
			got := ignore.Ignored(uc.file)
			if uc.ignored != got {
				t.Errorf("wait %v got %v", uc.ignored, got)
			}
		})
	}
}
