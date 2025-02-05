package container

import (
	"fmt"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
    c := &Container{}

    if err := c.Register(c, reflect.TypeOf(c)); err != nil {
        t.Errorf("err: %v\n", err)
        return;
    }

    obj, err := c.Resolve(reflect.TypeOf(c))
    if err != nil {
        t.Errorf("err: %v\n", err)
        return 
    }
    if s, ok := obj.(*Container); ok {
        fmt.Printf("s: %v\n", s)
        return
    }
    t.Errorf("Did not resolve container!")
}
