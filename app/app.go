package app

import (
	"fmt"
	"time"

	"github.com/meian/gowatch/cmd"
	"github.com/urfave/cli/v2"
)

// NewApp アプリインスタンスを返す
func NewApp() *cli.App {
	app := &cli.App{
		Name:      "gowatch",
		Version:   "0.0.1",
		Compiled:  time.Now(),
		Copyright: "(c) 2020 kitamin",
		Usage:     "Watch file change and run test",
		Authors: []*cli.Author{
			&cli.Author{
				Name: "kitamin",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "run with show verbose",
			},
		},
		HideVersion: true,
		OnUsageError: func(c *cli.Context, err error, isSub bool) error {
			fmt.Println(err)
			fmt.Println("show usage with:", c.App.Name, "-h")
			return err
		},
	}
	app.HelpName = app.Name

	app.Commands = []*cli.Command{
		cmd.Test,
		cmd.Version,
	}

	return app
}
