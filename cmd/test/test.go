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
		testMap, _ := c.Changed.PopAll()
		c.State = Executing
		failed := []*file.Pair{}
		s, f := 0, 0
		for testFile, pairs := range testMap {
			log.Println("run test target:", testFile)
			args, err := cmdArgs(c, testFile)
			if err != nil {
				log.Println("error at command args", err)
				failed = append(failed, pairs...)
				f++
				continue
			}
			cmd := newCommand("go", args...)
			log.Println("run test:", cmd.viewMsg())
			err = cmd.Run()
			if err != nil {
				failed = append(failed, pairs...)
				f++
			} else {
				s++
			}
		}
		c.Changed.Add(failed...)
		log.Printf("all tests are executed: success = %d, failure = %d", s, f)
		c.State = None
	}
}

func cmdArgs(c *Context, src string) ([]string, error) {
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
