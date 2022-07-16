package list

import (
	"unsafe"

	list_head "github.com/kazu/lista_encabezado"
)

type PtrList[T any] interface {
	*T
}

type List[T any] struct {
	Element T
	list_head.ListHead
}

func (l *List[T]) Offset() uintptr {
	return unsafe.Offsetof(l.ListHead)
}

func (l *List[T]) PtrListHead() *list_head.ListHead {
	return &l.ListHead
}

func (l *List[T]) FromListHead(head *list_head.ListHead) list_head.List {
	return l.FromHead(head)
}

func (l *List[T]) FromHead(head *list_head.ListHead) *List[T] {
	return (*List[T])(list_head.ElementOf(l, head))
}

func New[T any]() *List[T] {

	l := new(List[T])
	l.Init()

	return l
}
