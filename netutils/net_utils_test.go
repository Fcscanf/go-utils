package netutils

import (
	"fmt"
	"testing"
)

func TestLocalIPv4s(t *testing.T) {

	if localIPs, err := LocalIPv4s(); err == nil {
		for _, iP := range localIPs {
			fmt.Println(iP)
		}
	}
}
