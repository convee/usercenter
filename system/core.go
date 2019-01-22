package system

import (
	"path/filepath"
	"os"
	"crypto/md5"
	"encoding/hex"
)

//根目录
func RootPath() string  {
	return filepath.Join(os.Getenv("GOPATH"), "/src/usercenter")
}

func Md5(source string) string  {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}