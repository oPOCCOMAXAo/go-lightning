package lightning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Push(t *testing.T) {
	l := &list{}
	l.Push(Vector{X: 1}, true)
	l.Push(Vector{X: 2}, false)
	l.Push(Vector{X: 3}, true)
	l.Push(Vector{X: 4}, false)

	assert.Equal(
		t,
		[]Vector{{X: 3}, {X: 1}, {X: 2}, {X: 4}},
		l.data,
	)
}

func TestList_Take(t *testing.T) {
	l := &list{
		data: []Vector{{X: 3}, {X: 1}, {X: 2}, {X: 4}},
	}
	takeOrder := []bool{false, true, false, true}

	var res []Vector
	for _, b := range takeOrder {
		v, _ := l.Take(b)
		res = append(res, v)
	}

	assert.Equal(
		t,
		[]Vector{{X: 4}, {X: 3}, {X: 2}, {X: 1}},
		res,
	)
}
