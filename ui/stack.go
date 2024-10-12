package ui

import (
	"image"

	"github.com/bin16/wooden-fish/app"
	"github.com/bin16/wooden-fish/util"
)

type Stack struct {
	Box
}

func (u *Stack) Layout(ow, oh int) (bw, bh int) {
	for _, n := range u.Children() {
		var w, h = n.Layout(ow, oh)
		bw = max(bw, w)
		bh = max(bh, h)
	}

	return
}

func (u *Stack) SetBounds(r image.Rectangle) {
	u.Scene.SetBounds(r)
	for _, n := range u.Children() {
		n.SetBounds(r)
	}
}

type StackOpt func(u *Stack)
type StackOptions struct{}

var StackOpts StackOptions

func (StackOptions) Layer(n app.Scene) StackOpt {
	return func(u *Stack) {
		u.children = append(u.children, n)
	}
}

func NewStack(opts ...StackOpt) *Stack {
	var stack = &Stack{}
	for _, o := range opts {
		o(stack)
	}

	return stack
}

func Layers(items ...app.Scene) *Stack {
	var stack = NewStack(
		util.Map(items, StackOpts.Layer)...,
	)

	return stack
}
