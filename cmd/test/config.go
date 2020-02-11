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
	// Verbose は詳細を表示するかどうか
	Verbose bool
}

// GetConfig はコマンド実行情報から生成した設定値を返す
func GetConfig(c *cli.Context) *Config {
	return &Config{
		Dir:       getDir(c),
		Recursive: c.Bool("recursive"),
		Verbose:   c.Bool("verbose"),
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
