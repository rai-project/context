package context

import "context"

type Context struct {
	context.Context
}

func New() Context {
	return Background()
}

func Background() Context {
	return Context{context.Background()}
}

func TODO() Context {
	return Context{context.TODO()}
}

func (c *Context) WithValue(key, val interface{}) Context {
	return Context{context.WithValue(c, key, val)}
}
