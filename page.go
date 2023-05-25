package paging

import (
	"fmt"
	"strings"
)

// Page 分页对象
type Page struct {
	// 当前页
	Page int `default:"1" json:"page" yaml:"page" xml:"page" toml:"page" validate:"min=1"`
	// 每页个数
	Size int `default:"20" json:"size" yaml:"size" xml:"size" toml:"size" validate:"min=1"`
	// 查询关键字
	Keyword string `json:"keyword" yaml:"keyword" xml:"keyword" toml:"keyword"`
	// 排序顺序
	// nolint: lll
	Order string `default:"DESC" json:"order" yaml:"order" xml:"order" toml:"order" validate:"oneof=asc ASC ascending ASCENDING desc DESC descending DESCENDING"`
	// 排序字段
	Sort string `json:"sort" yaml:"sort" xml:"sort" toml:"sort"`
}

func (p *Page) OrderBy() string {
	order := asc
	if strings.HasPrefix(strings.ToUpper(p.Order), desc) {
		order = desc
	}

	return fmt.Sprintf("`%s` %s", p.Sort, order)
}

func (p *Page) Mysql() (start int, offset int) {
	return p.Size, (p.Page - 1) * p.Size
}

func (p *Page) Start() int {
	return (p.Page - 1) * p.Size
}

func (p *Page) Limit() int {
	return p.Size
}
