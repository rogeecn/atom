package rbac

type IRbac interface {
	Can(role, method, path string) bool
	Reload() error
}
