//+build windows

package path

import "regexp"

func buildPattern() *regexp.Regexp {
	return regexp.MustCompile(`^(\w:|\.{0,2}/)`)
}
