package tap

func T[T any](v T, ds ...func(T)) T {
	for _, d := range ds {
		d(v)
	}
	return v
}
