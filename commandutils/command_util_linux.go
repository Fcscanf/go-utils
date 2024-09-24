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
//	Run the script using the sh command interpreter
func RunSh(args ...string) (string, error) {
	return Run("/bin/sh", args...)
}

// RunCommand
//
//	Use the sh command interpreter to execute commands
func RunCommand(args string) (string, error) {
	return Run("/bin/sh", "-c", args)
}
