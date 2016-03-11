package tasks

import (
	"fmt"
	"io/ioutil"
)

import "github.com/AlexPikalov/goro"

func Logger(ctx *goro.Context) error {
	reader, err := ctx.GetReader()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	log := ctx.GetLog()
	log.Println("Regular log")
	fmt.Println("READ", string(data))
	return nil
}
