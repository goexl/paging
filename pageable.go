package paging

// Pageable 可被分页
type Pageable interface {
	// OrderBy 排序字符串
	OrderBy(sorter Sorter) string

	// Mysql 分页参数
	Mysql() (start int, offset int)

	// Start 开始下标
	Start() int

	// Limit 限制个数
	Limit() int
}
