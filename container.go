package gfac

import (
	"fmt"
	"reflect"
)

// Exported
func NewLifetime() ILifetime {
    return (&container{}).Init().Resolve((*ILifetime)(nil)).(ILifetime)
}

type container struct {
    ctors map[reflect.Type]reflect.Value
    objects map[reflect.Type]reflect.Value
}

func (*container) Init() *container {
    cont := &container{
        ctors: make(map[reflect.Type]reflect.Value),
        objects: make(map[reflect.Type]reflect.Value),
    }
    cont.Register(&container{}, (*ILifetime)(nil))

    return cont
}

type ILifetime interface {
    Register(source any, target any)
    Resolve(target any) any
}

func (c *container) Register(source any, target any) {
    sourceValue := reflect.ValueOf(source)
    targetType := reflect.TypeOf(target).Elem()

    _, found := c.ctors[targetType]

    if found {
        panic(fmt.Sprintf("The type (%v) is already registered!", targetType))
    }

    if !sourceValue.Type().Implements(targetType) {
        panic(fmt.Sprintf("%v does not implement %v", sourceValue.Type(), targetType))

    }

    m := sourceValue.MethodByName("Init")
    if !m.IsValid() {
        panic(fmt.Sprintf("The type (%v) doesn't have an Init constructor!", sourceValue))
    }

    c.ctors[targetType] = m;
}

func (c *container) Resolve(target any) any {
    return c.getOrCreate(reflect.TypeOf(target).Elem()).Interface()
}

func (c *container) getOrCreate(t reflect.Type) reflect.Value {
    ctor, found := c.ctors[t];

    if !found {
        panic(fmt.Sprintf("The type (%v) is not registered!", t))
    }

    obj, found := c.objects[t];
    if found {
        return obj
    }

    obj = c.constructObject(ctor)
    c.objects[t] = obj

    return obj
}

func (c *container) constructObject(ctor reflect.Value) reflect.Value {
    ctorType := ctor.Type()
    numParams := ctorType.NumIn()
    params := make([]reflect.Value, numParams)

    for i := 0; i < numParams; i++ {
        pType := ctorType.In(i)
        params[i] = c.getOrCreate(pType)
    }

    return ctor.Call(params)[0]
}
