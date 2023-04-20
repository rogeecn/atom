package response

type PageResponse[T any] struct {
	Items []T    `json:"items,omitempty"`
	Total uint64 `json:"total,omitempty"`
}
