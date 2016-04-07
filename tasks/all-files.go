package tasks

import (
	"encoding/json"

	"github.com/AlexPikalov/goro"
	"github.com/AlexPikalov/goro/utils/files"
)

func AllFiles(ctx *goro.Context) error {
	writer, err := ctx.GetWriter()
	if err != nil {
		return err
	}

	patterns := ctx.GetArgs()
	ps := make([]string, 0)
	for _, p := range patterns {
		ps = append(ps, p.(string))
	}

	allFiles, err := files.FindAll(ps...)
	if err != nil {
		return err
	}
	data, err := json.Marshal(allFiles)
	if err != nil {
		return err
	}
	writer.Write(data)
	return nil
}
