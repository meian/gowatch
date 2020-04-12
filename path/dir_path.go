package path

import (
	"path/filepath"
	"regexp"
)

var (
	pattern *regexp.Regexp
)

func init() {
	pattern = buildPattern()
}

// DirPath はソースコードのディレクトリパスを返す
func DirPath(src string) string {
	pkg := filepath.Dir(src)
	switch pkg {
	case "", ".":
		return "."
	case "..":
		return ".."
	}
	pkg = filepath.ToSlash(pkg)
	if pattern.MatchString(pkg) {
		return pkg
	}
	return "./" + pkg
}
