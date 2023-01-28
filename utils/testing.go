package utils

import (
	"os"
	"strings"
)

func IsInTesting() bool {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.v=") {
			return true
		}
	}
	return false
}
