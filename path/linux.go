//+build !windows

package path

import "regexp"

func buildPattern() *regexp.Regexp {
	return regexp.MustCompile(`^\.{0,2}/`)
}
