package commandutils

import (
	"os/exec"
)

// Run
//
//	The first parameter is the command interpreter or a standalone executable program.
//	Some programs require the help of a command interpreter, such as ls,
//	while others can be executed independently using the program name, such as ipconfig
func Run(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).CombinedOutput()
	return string(out), err
}

// RunSh
//
//	Execute a Shell script file, such as test.sh
func RunSh(args ...string) (string, error) {
	return Run("/bin/sh", args...)
}

// RunCommand
//
//	Use the sh command interpreter to execute a single-line command, such as ls
func RunCommand(args string) (string, error) {
	return Run("/bin/sh", "-c", args)
}
