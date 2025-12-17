package core

import (
	"fmt"
	"strings"

	"github.com/goexl/paging/internal/constant"
)

// Paging 分页对象
type Paging struct {
	// 当前页
	Page int `default:"1" json:"page,omitempty" query:"page" validate:"min=1"`
	// 每页个数
	Size int `default:"20" json:"size,omitempty" query:"size" validate:"min=1"`
	// 查询关键字
	Keyword string `json:"keyword,omitempty" query:"keyword"`
	// 排序顺序
	Order string `default:"DESC" json:"order,omitempty" query:"order" validate:"oneof=asc ASC ascending ASCENDING desc DESC descending DESCENDING"` // nolint: lll
	// 排序字段
	Sort string `json:"sort,omitempty" query:"sort"`
}

func (p *Paging) OrderBy() string {
	order := constant.ASC
	if strings.HasPrefix(strings.ToUpper(p.Order), constant.DESC) {
		order = constant.DESC
	}

	return fmt.Sprintf("`%s` %s", p.Sort, order)
}

func (p *Paging) MySQL() (start int, offset int) {
	return p.Size, (p.Page - 1) * p.Size
}

func (p *Paging) Start() int {
	return (p.Page - 1) * p.Size
}

func (p *Paging) Limit() int {
	return p.Size
}
