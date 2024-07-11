package commandutils

import (
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	ipConfigO, err := Run("echo", "Hello")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ipConfigO)
}
