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
