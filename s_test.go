package gfac

import (
	"testing"
)

func TestMain(t *testing.T) {
    c := (&Container{}).Init()

    c.Register(&SomeStruct{}, (*DoOfSomething)(nil))
    c.Register(&AStruct{}, (*ASomething)(nil))
    c.Register(&BStruct{}, (*BSomething)(nil))
    c.Register(&CStruct{}, (*CSomething)(nil))

    doer := c.Resolve((*DoOfSomething)(nil)).(DoOfSomething)

    if doer.DoSomething() != 42069 {
        t.Fail()
    }
}

type SomeStruct struct {
    a ASomething
    b BSomething
    c CSomething
}

func (*SomeStruct) Init(a ASomething, b BSomething, c CSomething) *SomeStruct {
    return &SomeStruct{
        a,
        b,
        c,
    }
}

func (a *SomeStruct) DoSomething() int {
    return a.a.ASomething() + a.b.BSomething() * 10 + a.c.CSomething()
}

type DoOfSomething interface {
    DoSomething() int
}

type AStruct struct {}
type BStruct struct {}
type CStruct struct {}

func (a *AStruct) Init(b BSomething, c CSomething) *AStruct {
    return &AStruct{}
}

func (a *BStruct) Init(c CSomething) *BStruct {
    return &BStruct{}
}

func (a *CStruct) Init() *CStruct {
    return &CStruct{}
}

func (a *AStruct) ASomething() int {
    return 8 * 8 + 5
}

func (a *BStruct) BSomething() int {
    return 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 10
}

func (a *CStruct) CSomething() int {
    return 41 * 4 * 100
}

type ASomething interface { ASomething() int }
type BSomething interface { BSomething() int }
type CSomething interface { CSomething() int }
