package test

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

// Config はWatcherの設定を保持する
type Config struct {
	// Dir は監視対象のディレクトリパス
	Dir string
	// Recursive はサブディレクトリを監視するかどうか
	Recursive bool
}

// NewConfig は設定情報を生成する
func NewConfig(c *cli.Context) *Config {
	return &Config{
		Dir:       getDir(c),
		Recursive: c.Bool("recursive"),
	}
}

func getDir(c *cli.Context) string {
	if path := c.Args().Get(0); path != "" {
		return path
	}
	return "."
}

// Show はコンフィグの情報を表示する
func (c *Config) Show() {
	println("Config:")
	json, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(json))
}
