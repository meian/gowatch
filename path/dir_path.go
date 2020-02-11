package path

import (
	"path/filepath"
	"regexp"

	"github.com/meian/gowatch/util"
)

// DirPath はソースコードのディレクトリパスを返す
func DirPath(src string) string {
	pkg := UnixPath(filepath.Dir(src))
	if pkg == "" {
		return pkg
	}
	if m, _ := regexp.MatchString(pattern(), pkg); m {
		return pkg
	}
	return "./" + pkg
}

func pattern() string {
	if util.IsWindows() {
		return `^(\w:|\.{0,2}/)`
	}
	return `^\.{0,2}/`
}
