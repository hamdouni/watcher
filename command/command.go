/* Build, launch or stop a go program. */
package command

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var cmd *exec.Cmd

// run executes go command
func run(args ...string) error {
	c := exec.Command("go", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// Test invokes the go test chain.
func Test(srcpath string, args string) error {
	return run("test", srcpath)
}

// Launch invokes the go compiler to build the program in the tmp folder and launch it.
func Launch(srcpath string, arguments string) error {
	// we build the program in a temp dir
	dir, err := os.MkdirTemp("", "watcherdir")
	if err != nil {
		return fmt.Errorf("creates temp dir: %w", err)
	}
	defer os.RemoveAll(dir)
	tmpbin := filepath.Join(dir, "out.bin")
	if err := run("build", "-o", tmpbin, srcpath); err != nil {
		return fmt.Errorf("builds go program: %w", err)
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

// Kill kills the previously executed command
func Kill() error {
	if err := cmd.Process.Kill(); err != nil {
		return fmt.Errorf("killing command: %w", err)
	}
	return cmd.Wait()
}
