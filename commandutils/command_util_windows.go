package commandutils

import (
	"github.com/Fcscanf/go-utils/charsetutils"
	"os/exec"
)

// Run
//
//	The first parameter is the command interpreter or a standalone executable program.
//	Some programs require the help of a command interpreter, such as ls,
//	while others can be executed independently using the program name, such as ipconfig
func Run(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		return "", err
	}
	output, _ := charsetutils.GbkToUtf8(out)
	return string(output), nil
}
