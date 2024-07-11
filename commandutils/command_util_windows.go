package commandutils

import (
	"github.com/Fcscanf/go-utils/charsetutils"
	"os/exec"
)

func Run(command ...string) (string, error) {
	out, err := exec.Command(command[0], command[1:]...).CombinedOutput()
	if err != nil {
		return "", err
	}
	output, _ := charsetutils.GbkToUtf8(out)
	return string(output), nil
}
