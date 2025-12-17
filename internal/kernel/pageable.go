package kernel

// Pageable 可被分页
type Pageable interface {
	// OrderBy 排序字符串
	OrderBy() string

	// Limit 分页参数
	Limit() (start int, offset int)

	// Start 开始下标
	Start() int
}
