package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Version はバージョン番号を出力するコマンド
var Version *cli.Command

func init() {
	Version = &cli.Command{
		Name:    "version",
		Usage:   "print app version",
		Aliases: []string{"v"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "appname",
				Aliases: []string{"a"},
				Usage:   "print with app name",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("a") {
				fmt.Println(c.App.Name, c.App.Version)
			} else {
				fmt.Println(c.App.Version)
			}
			return nil
		},
	}
}
