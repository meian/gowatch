package path

import (
	"path/filepath"
	"regexp"

	"github.com/meian/gowatch/runtime"
)

// DirPath はソースコードのディレクトリパスを返す
func DirPath(src string) string {
	pkg := filepath.Dir(filepath.ToSlash(src))
	switch pkg {
	case "", ".":
		return "."
	case "..":
		return ".."
	}
	pkg = filepath.ToSlash(pkg)
	if m, _ := regexp.MatchString(pattern(), pkg); m {
		return pkg
	}
	return "./" + pkg
}

func pattern() string {
	if runtime.IsWindows {
		return `^(\w:|\.{0,2}/)`
	}
	return `^\.{0,2}/`
}
