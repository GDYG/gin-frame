package file

import (
	"fmt"
	"os"
	"time"
)

var (
	TimeFormat = "20060102"
	SavePath   = "runtime/files/"
	FileExt    = "txt"
)

func GetSavePath() string {
	return fmt.Sprintf("%s", SavePath)
}

func GetSaveFullPath() string {
	prefixPath := GetSavePath()
	suffixPath := fmt.Sprintf("%s.%s", time.Now().Format(TimeFormat), FileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func OpenFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		MkDir()
	case os.IsPermission(err):
		fmt.Printf("%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return handle
}

func MkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+GetSavePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
