package main

import (
	"fmt"
	"reflect"
)

func main() {
    s := Some{}
    t := reflect.TypeOf(s)
    m := t.Method(0)
    fmt.Printf("m.Type: %v\n", m.Type)
    fmt.Printf("m.Type.In(0): %v\n", m.Type.In(1))
    fmt.Printf("m.Func: %v\n", m.Func)
    fmt.Printf("m.Name: %v\n", m.Name)
}

type Some struct {}
func (s Some) HelloWorld(c int) string {
    return "Hell World!"
}
