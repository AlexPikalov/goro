package goro

import (
	"fmt"
	"log"
)

type Task struct {
	Name       string
	in         *PipeSync
	out        *PipeSync
	downstream *PipeSync
	next       []*Task
	worker     Worker
}

func NewTask(workerFn WorkerFn, logger *log.Logger) *Task {
	pipe := NewPipeSync()

	task := &Task{
		out:        pipe,
		downstream: pipe,
		worker:     *NewWorker(workerFn),
	}

	ctx := &Context{
		reader: task.in,
		writer: task.out,
		log:    logger,
	}

	task.worker.SetContext(ctx)

	return task
}

func (t *Task) SetInput(in *PipeSync) error {
	if t.HasInput() {
		return fmt.Errorf("Unable to set new input, it already has one")
	}

	t.in = in
	t.worker.context.SetReader(in)
	return nil
}

func (t *Task) HasInput() bool {
	return t.in != nil
}

func (t *Task) Pipe(task *Task) *Task {
	err := task.SetInput(t.downstream)
	if err != nil {
		panic(err.Error())
	}
	t.next = append(t.next, task)
	return task
}

// RunSync runs worker one time
func (t *Task) RunSync() error {
	return t.worker.Run()
}

// RunPipeSync runs each task in the pipeline downstream
// calling task.RunSync recursively
func (t *Task) RunPipeSync() error {
	err := t.RunSync()
	if err != nil {
		return err
	}

	if len(t.next) != 0 {
		for _, next := range t.next {
			err := next.RunPipeSync()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
