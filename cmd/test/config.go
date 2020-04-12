package test

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

// Config はWatcherの設定を保持する
type Config struct {
	// Dirs は監視対象のディレクトリパス一覧
	Dirs []string
	// Args は go test に渡す引数
	Args []string
	// Recursive はサブディレクトリを監視するかどうか
	Recursive bool
}

// NewConfig は設定情報を生成する
func NewConfig(c *cli.Context) (*Config, error) {
	dirs, args := parseArgs(c.Args().Slice())
	correctDirs(dirs)
	for _, arg := range args {
		if arg == "-run" {
			return nil, errors.New("cannot use -run")
		}
	}
	return &Config{
		Dirs:      dirs,
		Args:      args,
		Recursive: c.Bool("recursive"),
	}, nil
}

func parseArgs(slice []string) (dirs []string, args []string) {
	dirs = []string{"."}
	if len(slice) == 0 {
		return dirs, []string{}
	}
	for i, arg := range slice {
		if strings.HasPrefix(arg, "-") {
			dirs, slice = slice[0:i], slice[i:]
			break
		}
	}
	if len(slice) > 0 && slice[0] == "--" {
		slice = slice[1:]
	}
	return dirs, slice
}

func correctDirs(dirs []string) {
	for i, dir := range dirs {
		dirs[i] = strings.TrimRight(filepath.ToSlash(dir), "/")
	}
}

// Show はコンフィグの情報を表示する
func (c *Config) Show() {
	fmt.Printf("directory: %v (recursive: %v)\n", c.Dirs, c.Recursive)
	fmt.Println("arguments:", c.Args)
}
