package file

import (
	"fmt"

	"github.com/meian/gowatch/path"
)

// Pair は検知されたファイルとテストファイルのペアを保持する
type Pair struct {
	// Detected は検知されたファイル
	Detected string
	// Test はテスト対象のファイル
	Test string
}

// NewPair は検知対象のファイルからファイルペアを返す。検知ファイルがgoファイルでなければエラーを返す。
func NewPair(detected string) (*Pair, error) {
	if !IsFile(detected) || !path.IsGoFile(detected) {
		return nil, path.NoGoFileError{Path: detected}
	}
	return &Pair{
		Detected: detected,
		Test:     path.ToTestPath(detected),
	}, nil
}

// TestEnabled はファイルペアのテストが有効であるかを返す
func (p *Pair) TestEnabled() bool {
	if p == nil {
		return false
	}
	return IsFile(p.Test)
}

func (p *Pair) String() string {
	return fmt.Sprint(p.Detected, " => ", p.Test)
}
