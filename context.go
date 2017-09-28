package context

import "context"

// Context ...
type Context struct {
	context.Context
}

// Background ...
func Background() Context {
	return Context{context.Background()}
}

// TODO ...
func TODO() Context {
	return Context{context.TODO()}
}

// WithValue ...
func (c Context) WithValue(key, val interface{}) Context {
	return Context{context.WithValue(c, key, val)}
}
