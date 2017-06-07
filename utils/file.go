package utils

import (
	"strings"
	"os"
	"fmt"
	"path/filepath"
	"io"
)

func AbsolutePath(p string) (string,error) {

	if strings.HasPrefix(p, "~") {
		home := os.Getenv("HOME")
		if home == "" {
			panic(fmt.Sprintf("can not found HOME in envs, '%s' AbsPh Failed!", p))
		}
		p = fmt.Sprint(home, string(p[1:]))
	}
	s, err := filepath.Abs(p)

	if nil != err {
		return  "",err
	}
	return s,nil
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	defer dst.Close()
	return io.Copy(dst, src)
}

