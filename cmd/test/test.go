package test

import (
	"log"
	"strings"
	"time"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/path"
	"github.com/meian/gowatch/runtime"
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
		testMap, noTests := c.Changed.PopAll()
		c.State = Executing
		failures := []*file.Pair{}
		success, skipped := 0, len(noTests)
		for _, nt := range noTests {
			log.Println("skip test for no test file:", nt)
		}
		for testFile, pairs := range testMap {
			log.Println("run test target:", testFile)
			for _, p := range pairs {
				log.Println(p)
			}
			args, err := cmdArgs(c, testFile)
			if err != nil {
				log.Println("error at command args", err)
				failures = append(failures, pairs...)
				continue
			}
			cmd := runtime.NewCommand("go", args...)
			log.Println("run test:", cmd)
			err = cmd.Run()
			if err != nil {
				failures = append(failures, pairs...)
			} else {
				success++
			}
		}
		failed := len(failures)
		c.Changed.Add(failures...)
		log.Printf("all tests are executed: success = %d, failure = %d, skipped = %d", success, failed, skipped)
		c.State = None
	}
}

func cmdArgs(c *Context, src string) ([]string, error) {
	pattern, err := testPattern(src)
	if err != nil {
		return nil, err
	}
	args := append([]string{"test", path.DirPath(src), "-run", pattern}, c.Args...)
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
