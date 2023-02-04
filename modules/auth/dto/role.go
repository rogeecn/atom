package dto

type RoleCreateRequest struct {
	UUID          string `json:"alias,omitempty"`
	Name          string `json:"name,omitempty"`
	ParentID      uint   `json:"parent_id,omitempty"`
	DefaultRouter string `json:"default_router,omitempty"`
}
