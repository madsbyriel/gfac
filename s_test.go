package gfac

import (
	"testing"
)

func TestMain(t *testing.T) {
	c := NewLifetime()

    _ = c.Resolve((*ILifetime)(nil)).(ILifetime)

	c.Register(&some_struct{}, (*doOfSomething)(nil))
	c.Register(&a_struct{}, (*a_something)(nil))
	c.Register(&b_struct{}, (*b_something)(nil))
	c.Register(&c_struct{}, (*c_something)(nil))

	doer := c.Resolve((*doOfSomething)(nil)).(doOfSomething)
	if doer.DoSomething() != 42069 {
		t.Fail()
	}
}

type some_struct struct {
	a a_something
	b b_something
	c c_something
}

func (*some_struct) Init(a a_something, b b_something, c c_something) *some_struct {
	return &some_struct{
		a,
		b,
		c,
	}
}

func (a *some_struct) DoSomething() int {
	return a.a.ASomething() + a.b.BSomething()*10 + a.c.CSomething()
}

type doOfSomething interface {
	DoSomething() int
}

type a_struct struct{}
type b_struct struct{}
type c_struct struct{}

func (a *a_struct) Init(b b_something, c c_something) *a_struct {
	return &a_struct{}
}

func (a *b_struct) Init(c c_something) *b_struct {
	return &b_struct{}
}

func (a *c_struct) Init() *c_struct {
	return &c_struct{}
}

func (a *a_struct) ASomething() int {
	return 8*8 + 5
}

func (a *b_struct) BSomething() int {
	return 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 10
}

func (a *c_struct) CSomething() int {
	return 41 * 4 * 100
}

type a_something interface{ ASomething() int }
type b_something interface{ BSomething() int }
type c_something interface{ CSomething() int }
