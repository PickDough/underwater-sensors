package cqrs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestResult struct {
	result string
}

type TestQuery struct {
	data string
	Query[TestResult]
}

type TestQueryHandler struct{}

func (qh TestQueryHandler) Handle(q TestQuery) (*TestResult, error) {

	return &TestResult{result: q.data}, nil
}

func TestCQRS(t *testing.T) {
	cq := New()
	qh := TestQueryHandler{}
	q := TestQuery{data: "test"}

	RegisterQuery[TestResult, TestQuery](
		cq,
		TestQuery{data: "test"},
		qh.Handle,
	)

	r, err := ExecuteQuery[TestResult, TestQuery](cq, q)

	assert.Nil(t, err)
	assert.Equal(t, q.data, r.result)
}
