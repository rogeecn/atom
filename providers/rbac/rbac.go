package rbac

type IRbac interface {
	Can(role, method, path string) bool
	JsonPermissionsForUser(string) (string, error)
	Reload() error
}
