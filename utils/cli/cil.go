package cli

func Tap[T any](cmd *T, funcs ...func(*T)) *T {
	for _, f := range funcs {
		f(cmd)
	}
	return cmd
}
