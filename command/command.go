/* Build, launch or stop a go program. */
package command

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var cmd *exec.Cmd

// RunGo executes go command
func RunGo(args ...string) error {
	c := exec.Command("go", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// Test invoke the go test chain.
func Test(srcpath string, args string) error {
	return RunGo("test", srcpath)
}

// Launch invoke the go compiler to build the program in the tmp folder and launch it.
func Launch(srcpath string, arguments string) error {

	// we build the program in a temp dir
	dir, err := ioutil.TempDir("", "watcherdir")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	tmpbin := filepath.Join(dir, "out.bin")
	if err := RunGo("build", "-o", tmpbin, srcpath); err != nil {
		return err
	}

	var args []string
	if len(arguments) > 0 {
		args = strings.Split(arguments, " ")
	}
	// now we launch the program without waiting for it to complete
	cmd = exec.Command(tmpbin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func Kill() error {
	if err := cmd.Process.Kill(); err != nil {
		return err
	}
	return cmd.Wait()
}
