package paging_test

import (
	"testing"

	"github.com/goexl/paging"
	"github.com/goexl/paging/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	assert.NotNil(t, paging.NewStack(1, &core.Pagination{
		Page:    1,
		Size:    1,
		Keyword: "",
		Order:   "",
		Sort:    "",
	}, 1))
}
