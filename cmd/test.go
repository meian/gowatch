package cmd

import (
	"fmt"
	"log"

	"github.com/meian/gowatch/cmd/test"
	"github.com/meian/gowatch/notify"
	"github.com/meian/gowatch/terminal"
	"github.com/urfave/cli/v2"
)

const (
	testErrorConfig int = iota + 1
	testErrorContext
	testErrorWatcher
	testErrorOnTest
)

// Test は変更監視とテストを実行するコマンド
var Test *cli.Command

func init() {
	Test = &cli.Command{
		Name:  "test",
		Usage: "watch file change and trigger test, default PATH is current directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "include subdirectories",
			},
		},
		ArgsUsage: "[PATH]",
		Action:    testAction,
	}
}

func testAction(c *cli.Context) error {
	config, err := test.NewConfig(c)
	if err != nil {
		return cli.Exit(fmt.Errorf("cannot create config: %s", err), testErrorConfig)
	}
	nc, err := test.NewContext(config)
	if err != nil {
		log.Println(err)
		return cli.Exit("cannot create context", testErrorContext)
	}
	nc.Watcher, err = newWatcher(nc)
	if err != nil {
		return cli.Exit("cannot create watcher", testErrorWatcher)
	}

	terminal.Clear()
	config.Show()
	fmt.Println("watch directories:", nc.Directories)

	go test.LoopFSEvent(nc)
	go test.LoopTest(nc)
	err = <-nc.Done
	if err != nil {
		return cli.Exit(err, testErrorOnTest)
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
