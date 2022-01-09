package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// CheckConfigurationFileExist 检查配置文件是否存在 并 返回配置文件路径
func CheckConfigurationFileExist(path string) (string, os.FileInfo, error) {
	var finallyErr error
	file, err := os.Stat(path)
	finallyErr = CatchErr(err, func(msg string) string {
		return configurationFileDoesNotExist
	}, false)
	return filepath.Join(GetPwd(), path), file, finallyErr
}
func GetPwd() string {
	dir, dirErr := os.Getwd()
	_ = CatchErr(dirErr, func(msg string) string {
		return msg
	}, true)
	return dir
}
func OpenFile(URL string) (*os.File, func()) {
	file, readErr := os.Open(URL)
	_ = CatchErr(readErr, func(msg string) string {
		return msg
	}, true)
	return file, func() {
		_ = file.Close()
	}
}
func ReadFile(URL string) []byte {
	byteFile, err := ioutil.ReadFile(URL)
	_ = CatchErr(err, func(msg string) string {
		return msg
	}, true)
	return byteFile
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func MkDir(path string) {
	err := os.Mkdir(path, os.ModePerm)
	_ = CatchErr(err, func(msg string) string {
		return msg
	}, true)
}

func CreateFile(fileName string, content string) {
	file, err := os.Create(fileName)
	_ = CatchErr(err, func(msg string) string {
		return msg
	}, true)
	_, err = file.Write([]byte(content))
	_ = CatchErr(err, func(msg string) string {
		return msg
	}, true)
}

func RemoveFile(path string) {
	removeErr := os.Remove(path)
	_ = CatchErr(removeErr, func(msg string) string {
		return msg
	}, true)
}
