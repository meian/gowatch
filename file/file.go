package file

import (
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/meian/gowatch/path"
)

const (
	// GoExt はgoソースの拡張子
	GoExt = ".go"
)

// FInfo is file info
type FInfo struct {
	// Name is file path
	Name string
	// Exists is file exists
	Exists bool
	// ModTime is file last modified time
	ModTime time.Time
	// Hash is crc32 in file binary
	Hash uint32
}

// NewFInfo create new FInfo
func NewFInfo(name string) (*FInfo, error) {
	file, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return &FInfo{Name: name, Exists: false}, nil
		}
		return nil, err
	}
	if file.IsDir() {
		return nil, path.NoGoFileError{Path: name}
	}
	hash, err := computeHash(file.Name())
	if err != nil {
		log.Print("cannot get file hash", err)
		return nil, err
	}
	f := &FInfo{
		Name:    name,
		Exists:  true,
		ModTime: file.ModTime(),
		Hash:    hash,
	}
	return f, nil
}

func computeHash(name string) (uint32, error) {
	d, err := ioutil.ReadFile(name)
	if err != nil {
		return 0, err
	}
	return crc32.ChecksumIEEE(d), nil
}
