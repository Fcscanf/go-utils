package commandutils

import (
	"os/exec"
)

func Run(command ...string) (string, error) {
	var out []byte
	var err error
	out, err = exec.Command(command[0], command[1:]...).CombinedOutput()
	return string(out), err
}
