package core

type Stack[T any] struct {
	Total int64 `json:"total,omitempty"`
	Items *[]*T `json:"items,omitempty"`

	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
}
