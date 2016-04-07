package main

import "github.com/AlexPikalov/goro"
import "github.com/AlexPikalov/goro/tasks"

// mandatory
func Run(g goro.Goro) error {
	prodTask := g.LoadTask("prod", "containers/**/*", "!containers/**/*file.go")
	logTask := g.LoadTask("log")
	prodTask.Pipe(logTask)
	return prodTask.RunPipeSync()
}

// mandatory
func Config(g goro.Goro) {
	g.AddTask("log", tasks.Logger)
	g.AddTask("prod", tasks.AllFiles)
}
