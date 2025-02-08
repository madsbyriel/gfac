package main

import (
	"fmt"
	"reflect"
)

type Container struct {
    ctors map[reflect.Type]reflect.Value
    objects map[reflect.Type]reflect.Value
}

func (*Container) Init() *Container {
    return &Container{
        ctors: make(map[reflect.Type]reflect.Value),
        objects: make(map[reflect.Type]reflect.Value),
    }
}

func (c *Container) Register(source any, target any) {
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

func (c *Container) Resolve(target any) any {
    return c.getOrCreate(reflect.TypeOf(target).Elem()).Interface()
}

func (c *Container) getOrCreate(t reflect.Type) reflect.Value {
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

func (c *Container) constructObject(ctor reflect.Value) reflect.Value {
    ctorType := ctor.Type()
    numParams := ctorType.NumIn()
    params := make([]reflect.Value, numParams)

    for i := 0; i < numParams; i++ {
        pType := ctorType.In(i)
        params[i] = c.getOrCreate(pType)
    }

    return ctor.Call(params)[0]
}
