package context

import "context"

type Context struct {
	context.Context
}

func (c *Context) WithValue(key, val interface{}) Context {
	return Context{context.WithValue(c, key, val)}
}
