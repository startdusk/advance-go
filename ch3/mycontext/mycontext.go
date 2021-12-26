package mycontext

import (
	"context"
	"reflect"
	"sync"
)

// 我们知道，在子节点中 WithValue 的数据，父节点是查不到的。
// 请实现一个 MyContext，其 WithValue 方法赋值的 k，v 在父节点中也可以查得到

type MyContext struct {
	context.Context
	m sync.Map
}

func (c *MyContext) WithValue(parent context.Context, key, val interface{}) context.Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	if key == nil {
		panic("nil key")
	}

	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	c.m.Store(key, val)
	return &MyContext{
		Context: parent,
		m:       c.m,
	}
}

func (c *MyContext) Value(key interface{}) interface{} {
	val, _ := c.m.Load(key)
	return val
}

func NewMyContext() *MyContext {
	return &MyContext{}
}
