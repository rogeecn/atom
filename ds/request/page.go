package request

type PageFilter struct {
	Page  uint64 `form:"page"`
	Limit uint64 `form:"limit"`
}
