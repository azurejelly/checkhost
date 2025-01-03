package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// from https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
func Ask(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s (y/N): ", s)

		r, err := reader.ReadString('\n')
		if err != nil {
			// simply return 'false' if stdin is unavailable
			return false
		}

		r = strings.ToLower(strings.TrimSpace(r))

		if r == "y" || r == "yes" {
			return true
		} else {
			return false
		}
	}
}
