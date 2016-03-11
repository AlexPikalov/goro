package tasks

import "github.com/AlexPikalov/goro"

func HW(ctx *goro.Context) error {
	writer, err := ctx.GetWriter()
	if err != nil {
		return err
	}
	writer.Write([]byte("Hello Goro!"))
	return nil
}
