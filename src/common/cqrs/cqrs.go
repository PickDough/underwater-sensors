package cqrs

import (
	"fmt"
	"reflect"
)

type CQRS struct {
	queries  map[reflect.Type]any
	commands map[reflect.Type]any
}

func New() *CQRS {
	return &CQRS{queries: make(map[reflect.Type]any), commands: make(map[reflect.Type]any)}
}

type Query[T any] interface{}

func RegisterQuery[T any, Q Query[T]](cq *CQRS, q Query[T], qh func(Q) (*T, error)) {
	cq.queries[reflect.TypeOf(q)] = qh
}

func ExecuteQuery[T any, Q Query[T]](cq *CQRS, query Query[T]) (*T, error) {
	handler, ok := cq.queries[reflect.TypeOf(query)]
	if !ok {
		return nil, fmt.Errorf("couldn't find handler for query")
	}

	var qh func(query Q) (*T, error)
	if qh, ok = handler.(func(Q) (*T, error)); ok {
		q := query.(Q)
		return qh(q)
	}

	return nil, fmt.Errorf("query handler is not compatible")
}

type Command interface{}

func RegisterCommand[C Command](cq *CQRS, c Command, ch func(C) error) {
	cq.commands[reflect.TypeOf(c)] = ch
}

func ExecuteCommand[Q Command](cq *CQRS, command Command) error {
	handler, ok := cq.commands[reflect.TypeOf(command)]
	if !ok {
		return fmt.Errorf("couldn't find handler for query")
	}

	var qh func(Q) error
	if qh, ok = handler.(func(Q) error); ok {
		q := command.(Q)
		return qh(q)
	}

	return fmt.Errorf("query handler is not compatible")
}
