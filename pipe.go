package goro

import "io"

type PipeSync struct {
	data []byte
}

func NewPipeSync() *PipeSync {
	ps := &PipeSync{}
	return ps
}

func (ps *PipeSync) Write(b []byte) (int, error) {
	ps.data = b
	return len(b), nil
}

func (ps *PipeSync) Read(b []byte) (int, error) {
	if len(ps.data) == 0 {
		return 0, io.EOF
	}
	n := copy(b, ps.data)
	ps.data = ps.data[n:]
	return n, nil
}
