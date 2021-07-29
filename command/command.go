/*
	Build, launch or stop a go program.
*/
package command

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var cmd *exec.Cmd

// Launch invoke the go compiler to build the program in the tmp folder and launch it.
func Launch(srcpath string) error {

	// we build the program in a temp dir
	dir, err := ioutil.TempDir("", "watcherdir")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	tmpbin := filepath.Join(dir, "out.bin")
	c := exec.Command("go", "build", "-o", tmpbin, srcpath)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()
	if err != nil {
		return err
	}

	// now we launch the program
	cmd = exec.Command(tmpbin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func Kill() error {
	err := cmd.Process.Kill()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
