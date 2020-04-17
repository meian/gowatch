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
	// Dir は監視対象のディレクトリパス
	Dir string
	// Args は go test に渡す引数
	Args []string
	// Recursive はサブディレクトリを監視するかどうか
	Recursive bool
}

// NewConfig は設定情報を生成する
func NewConfig(c *cli.Context) (*Config, error) {
	config := &Config{
		Args:      []string{},
		Recursive: c.Bool("recursive"),
	}
	args := c.Args().Slice()
	if len(args) == 0 {
		config.Dir = "."
		return config, nil
	}
	if !strings.HasPrefix(args[0], "-") {
		config.Dir = strings.TrimRight(filepath.ToSlash(args[0]), "/")
		args = args[1:]
	} else {
		config.Dir = "."
	}
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "-run" {
			return nil, errors.New("cannot use -run")
		}
	}
	config.Args = args
	return config, nil
}

// Show はコンフィグの情報を表示する
func (c *Config) Show() {
	fmt.Println("directory:", c.Dir)
	fmt.Println("recursive:", c.Recursive)
	fmt.Println("arguments:", c.Args)
}
