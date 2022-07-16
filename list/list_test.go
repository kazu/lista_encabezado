package list_test

import (
	"testing"

	list "github.com/kazu/lista_encabezado/list"
	"github.com/stretchr/testify/assert"
)

func Test_List(t *testing.T) {

	type A struct {
		I int
	}

	a := list.New[A]()
	a.Element.I = 123
	b := list.New[A]().FromHead(&a.ListHead)

	assert.NotNil(t, a)
	assert.Equal(t, a.Element.I, b.Element.I)
}
