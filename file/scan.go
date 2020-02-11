package file

import (
	"bufio"
	"os"
	"regexp"
)

var pattern *regexp.Regexp

// ScanTests はテストソース内からテスト関数名を収集する
func ScanTests(testPath string) ([]string, error) {
	file, err := os.Open(testPath)
	if err != nil {
		return nil, &ReadError{Name: testPath}
	}
	defer file.Close()
	names := make([]string, 0)
	sc := bufio.NewScanner(file)
	p := compilePattern()
	for sc.Scan() {
		line := sc.Text()
		m := p.FindStringSubmatch(line)
		if len(m) == 2 {
			names = append(names, m[1])
		}
	}
	return names, nil
}

// テスト関数を解析するパターン
func compilePattern() *regexp.Regexp {
	if pattern == nil {
		pattern = regexp.MustCompile(`^func\s+(Test\w+)\(\w+\s+\*testing\.T\)`)
	}
	return pattern
}
