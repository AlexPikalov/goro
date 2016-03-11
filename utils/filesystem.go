package utils

import (
	"io/ioutil"
	"os"
)

var FS *Filesystem

func init() {
	FS = &Filesystem{}
}

type Filesystem struct{}

func (fs *Filesystem) CopyFile(src, dest string) error {
	srcFile, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	_, err = destFile.Write(srcFile)
	return err
}
