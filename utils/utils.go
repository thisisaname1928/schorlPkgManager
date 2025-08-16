package utils

import "os"

func GetRealPath(path string) string {
	realPath, e := os.Readlink(path)
	if e != nil {
		return path
	}

	return realPath
}
