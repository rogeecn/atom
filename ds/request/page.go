package request

type PageFilter struct {
	Page    uint64 `form:"page"`
	PerPage uint64 `form:"per_page"`
}
