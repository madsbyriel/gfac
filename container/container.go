package container

import "reflect"

type Container struct {
    typeToAny map[reflect.Type]any
}

func (c *Container) Resolve(t reflect.Type) (any, error) {
    if c.typeToAny == nil {
        return nil, NotFoundResolveError{}
    }

    object, ok := c.typeToAny[t]
    if ok {
        return object, nil
    } else {
        return nil, NotFoundResolveError{}
    }
}

func (c *Container) Register(object any, t reflect.Type) error {
    if c.typeToAny == nil {
        c.typeToAny = make(map[reflect.Type]any)
    }

    if _, ok := c.typeToAny[t]; ok {
        return TypeAlreadyRegistered{}
    }
    c.typeToAny[t] = object

    return nil
}

type LifetimeScope interface {
    Resolve() (any, error)
}

type TypeAlreadyRegistered struct {}
func (e TypeAlreadyRegistered) Error() string {
    return "This type is already registered"
}

type NotFoundResolveError struct {}
func (e NotFoundResolveError) Error() string {
    return "Did not find the value given this type"
}
