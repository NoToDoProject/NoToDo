package common

import "os"

// IsFileExists 判断文件是否存在
func IsFileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if fileInfo != nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
