package cmd

import (
	"fmt"
	"log"

	"github.com/meian/gowatch/cmd/test"
	"github.com/meian/gowatch/notify"
	"github.com/meian/gowatch/terminal"
	"github.com/urfave/cli/v2"
)

// Test は変更監視とテストを実行するコマンド
var Test = &cli.Command{
	Name:  "test",
	Usage: "watch directory change and trigger test",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "recursive",
			Aliases: []string{"r"},
			Usage:   "include subdirectories",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "show detail",
		},
	},
	ArgsUsage: "[PATH]",
	Action:    testAction,
}

func testAction(c *cli.Context) error {
	config := test.GetConfig(c)
	nc, err := test.NewContext(config)
	if err != nil {
		log.Println(err)
		return cli.Exit("cannot create context", 1)
	}
	nc.Watcher, err = newWatcher(nc)
	if err != nil {
		return cli.Exit("cannot create watcher", 2)
	}

	terminal.Clear()
	config.Show()
	fmt.Println("watch directories:", nc.Directories)

	go test.LoopFSEvent(nc)
	go test.LoopTest(nc)
	err = <-nc.Done
	if err != nil {
		return cli.Exit(err, 3)
	}
	return nil
}

func newWatcher(c *test.Context) (*notify.Watcher, error) {
	watcher, err := notify.NewWatcher()
	if err != nil {
		return nil, err
	}
	for _, dir := range c.Directories {
		watcher.Add(dir)
	}
	return watcher, nil
}
