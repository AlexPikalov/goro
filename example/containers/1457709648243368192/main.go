package main

import (
	"fmt"
	"os"
)
import "github.com/AlexPikalov/goro"
import "github.com/AlexPikalov/goro/utils"

var GORO *goro.TaskManager

func init() {
	log, err := utils.NewGoroLog("./gorodebug.log")
	if err != nil {
		panic(err.Error())
	}
	GORO = goro.NewTaskManager(log)
}

func main() {
	Config(GORO)
	err := Run(GORO)
	if err != nil {
		fmt.Println("goro container error: ", err.Error())
		os.Exit(1)
	}
}
