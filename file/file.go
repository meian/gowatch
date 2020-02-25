package file

import "os"

// IsFile はパスがファイルであるかを返す
func IsFile(name string) bool {
	file, err := os.Stat(name)
	return err == nil && !file.IsDir()
}
