# Architecture

Architecture, components, workflow.

### Usage

`goro gorofile.go`

### Components

**goro**
Has following methods:

`Task(name string, func(input []byte)([]byte, error))` is a method that
registers new task. It takes two arguments a name of new task and a reductor
function. Reductor function takes one argument, it is a sequence of bytes that
could be transformed or utilized in some way.

`LoadTask(name string)(*Worker, error)` is a method that creates a worker that contains
a task which name is provided as first argument. In case of any failures it
returns an error
 