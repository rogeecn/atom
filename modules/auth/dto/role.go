package dto

type RoleRequestFilter struct {
	Name          *string `form:"name"`
	ParentID      *uint   `form:"parent_id"`
	DefaultRouter *string `form:"default_router"`
}

type RoleRequestForm struct {
	Name          string `json:"name,omitempty"`
	ParentID      uint   `json:"parent_id,omitempty"`
	DefaultRouter string `json:"default_router,omitempty"`
}

type RoleTree struct {
	ID            uint64      `json:"id,omitempty"`
	UUID          string      `json:"uuid,omitempty"`
	Name          string      `json:"name,omitempty"`
	ParentID      uint64      `json:"parent_id,omitempty"`
	DefaultRouter string      `json:"default_router,omitempty"`
	Children      []*RoleTree `json:"children,omitempty"`
}
