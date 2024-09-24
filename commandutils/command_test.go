package commandutils

import (
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	ipConfigO, err := Run("ipconfig")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ipConfigO)
}
