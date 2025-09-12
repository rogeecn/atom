package ptr

// Of returns a pointer to v.
func Of[T any](v T) *T { return &v }

// Deref returns the value pointed to by p, or def if p is nil.
func Deref[T any](p *T, def T) T {
    if p == nil {
        return def
    }
    return *p
}

// Zero returns the zero value for type T.
func Zero[T any]() (z T) { return }

