package vm

import (
	"Interpreter/code"
	"Interpreter/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          cl,
		basePointer: basePointer,
		ip:          -1,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
