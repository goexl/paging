package paging

import (
	"github.com/goexl/paging/internal/core"
)

type Stack[T any] = core.Stack[T]

func NewStack[T any](total int64, pagination *Pagination, items *[]*T) *Stack[T] {
	return &core.Stack[T]{
		Total: total,
		Items: items,

		Page: pagination.Page,
		Size: pagination.Size,
	}
}
