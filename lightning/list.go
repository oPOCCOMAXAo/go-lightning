package lightning

import (
	"fmt"
)

type list struct {
	data []Vector
}

func (l list) String() string {
	return fmt.Sprint(l.data)
}

func (l *list) Push(value Vector, toHead bool) {
	if toHead {
		l.data = append([]Vector{value}, l.data...)
	} else {
		l.data = append(l.data, value)
	}
}

func (l *list) Take(fromHead bool) (res Vector, ok bool) {
	length := len(l.data)
	if length == 0 {
		return Vector{}, false
	}
	if fromHead {
		res = l.data[0]
		l.data = l.data[1:]
	} else {
		res = l.data[length-1]
		l.data = l.data[:length-1]
	}
	return res, true
}

func (l *list) Peek(fromHead bool) (res Vector, ok bool) {
	length := len(l.data)
	if length == 0 {
		return Vector{}, false
	}
	if fromHead {
		res = l.data[0]
	} else {
		res = l.data[length-1]
	}
	return res, true
}

type List struct {
	list
	pushHead bool
	takeHead bool
}

func (l *List) Push(v Vector) {
	l.list.Push(v, l.pushHead)
}

func (l *List) Pop() (Vector, bool) {
	return l.list.Take(l.takeHead)
}

func (l *List) Peek() (Vector, bool) {
	return l.list.Peek(l.takeHead)
}

func (l *List) Array() []Vector {
	return append([]Vector{}, l.data...)
}

func NewStack() *List {
	return &List{pushHead: false, takeHead: false}
}

func NewQueue() *List {
	return &List{pushHead: false, takeHead: true}
}
