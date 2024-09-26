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

// RunSh
//
//	Execute script files, such as test.bat,test.cmd
func RunSh(args ...string) (string, error) {
	return Run(args[0], args[1:]...)
}

// RunCommand
//
//	Use the cmd command interpreter to execute single-line commands, such as dir
func RunCommand(args ...string) (string, error) {
	return Run("cmd", append([]string{"/c"}, args...)...)
}
