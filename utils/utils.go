package utils

import "os"

func GetRealPath(path string) string {
	realPath, e := os.Readlink(path)
	if e != nil {
		return path
	}

	return realPath
}

func IsFileExist(path string) bool {
	_, e := os.Stat(path)
	if os.IsNotExist(e) {
		return false
	}

	return true
}
