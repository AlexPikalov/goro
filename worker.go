package goro

import (
	"errors"
	"log"
)

type Worker struct {
	//	sync.RWMutex
	context *Context
	runner  WorkerFn
}

func NewWorker(runner WorkerFn) *Worker {
	return &Worker{runner: runner}
}

func (w *Worker) GetContext() *Context {
	//	w.RLock()
	//	defer w.RUnlock()

	return w.context
}

func (w *Worker) SetContext(ctx *Context) {
	//	w.Lock()
	//	defer w.Unlock()

	w.context = ctx
}

func (w *Worker) Run() error {
	//	w.Lock()
	//	defer w.Unlock()
	ctx := w.GetContext()
	return w.runner(ctx)
}

type Context struct {
	reader *PipeSync
	writer *PipeSync
	log    *log.Logger
}

func (ctx *Context) GetReader() (*PipeSync, error) {
	var err error
	if ctx.reader == nil {
		err = errors.New("reader is undefined")
	}
	return ctx.reader, err
}

func (ctx *Context) SetReader(r *PipeSync) {
	ctx.reader = r
}

func (ctx *Context) GetWriter() (*PipeSync, error) {
	var err error
	if ctx.writer == nil {
		err = errors.New("writer is undefined")
	}
	return ctx.writer, err
}

func (ctx *Context) SetWriter(w *PipeSync) {
	ctx.writer = w
}

func (ctx *Context) GetLog() *log.Logger {
	return ctx.log
}

type WorkerFn func(*Context) error
