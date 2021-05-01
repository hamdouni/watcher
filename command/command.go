/*
	Build, launch or stop a go program.
*/
package command

import (
	"os"
	"os/exec"
)

var cmd *exec.Cmd

// Launch invoke the go compiler to build the program in the tmp folder and launch it.
func Launch() error {

	// we build the program in the tmp folder
	c := exec.Command("go", "build", "-o", "/tmp/out.bin", ".")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		return err
	}

	// now we launch the program
	cmd = exec.Command("/tmp/out.bin")
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
