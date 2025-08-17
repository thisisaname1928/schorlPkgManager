package utils

import (
	"io"
	"os"
	"syscall"
)

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

func CopyFile(src string, dest string) error {
	fsrc, e := os.Open(src)
	if e != nil {
		return e
	}
	defer fsrc.Close()

	dsrc, e := os.Create(dest)
	if e != nil {
		return e
	}

	_, e = io.Copy(dsrc, fsrc)
	if e != nil {
		return e
	}

	s, e := os.Stat(src)
	if e != nil {
		return e
	}

	rawStat := s.Sys().(*syscall.Stat_t)

	e = dsrc.Chmod(os.FileMode(rawStat.Mode))
	if e != nil {
		return e
	}
	e = dsrc.Chown(int(rawStat.Uid), int(rawStat.Gid))
	if e != nil {
		return e
	}

	dsrc.Close()
	return nil
}
