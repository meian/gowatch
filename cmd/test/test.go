package test

import (
	"log"
	"strings"
	"time"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/path"
)

// LoopTest はテストを実行するループ
func LoopTest(c *Context) {
	tick := time.Tick(500 * time.Millisecond)
	for {
		<-tick
		if !c.Triggered || c.State != None {
			continue
		}
		if c.Changed.Count() == 0 {
			<-tick
			continue
		}
		c.Triggered = false
		c.State = Waiting
		<-tick
		fs := c.Changed.PopAll()
		log.Println("run test target:", fs)
		c.State = Executing
		failed := []string{}
		for _, f := range fs {
			args, err := cmdArgs(f, c)
			if err != nil {
				log.Println("error at command args", err)
				failed = append(failed, f)
				continue
			}
			cmd := newCommand("go", args...)
			log.Println("run test:", cmd.viewMsg())
			err = cmd.Run()
			if err != nil {
				failed = append(failed, f)
			}
		}
		if len(failed) > 0 {
			log.Println("return changed to failed tests:", failed)
			c.Changed.AddSlice(failed)
		}
		c.State = None
	}
}

func cmdArgs(src string, c *Context) ([]string, error) {
	pattern, err := testPattern(src)
	if err != nil {
		return nil, err
	}
	args := []string{"test", path.DirPath(src), "-run", pattern}
	if c.Config.Verbose {
		args = append(args, "-v")
	}
	return args, nil
}

func testPattern(filePath string) (string, error) {
	tests, err := file.ScanTests(filePath)
	if err != nil {
		return "", err
	}
	nopfx := make([]string, 0, len(tests))
	for _, t := range tests {
		nopfx = append(nopfx, strings.TrimPrefix(t, "Test"))
	}
	return "^Test(" + strings.Join(nopfx, "|") + ")$", nil
}
