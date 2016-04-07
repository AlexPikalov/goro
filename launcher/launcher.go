package main

import (
	"fmt"
	"os"
)
import "github.com/AlexPikalov/goro"
import "github.com/AlexPikalov/goro/utils/gorolog"

var GORO *goro.TaskManager

func init() {
	glog, err := gorolog.NewGoroLog("./gorodebug.log")
	if err != nil {
		panic(err.Error())
	}
	GORO = goro.NewTaskManager(glog)
}

func main() {
	Config(GORO)
	err := Run(GORO)
	if err != nil {
		fmt.Println("goro container error: ", err.Error())
		os.Exit(1)
	}
}
