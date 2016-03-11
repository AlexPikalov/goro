package utils

import (
	"log"
	"os"
	"path/filepath"
)

func NewGoroLog(file string) (*log.Logger, error) {
	path, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	out, err := os.Create(path)

	l := log.New(out, "[Goro Debug]", log.Llongfile|log.LstdFlags)
	return l, nil
}
