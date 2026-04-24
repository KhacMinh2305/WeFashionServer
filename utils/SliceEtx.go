package utils

func Map[T, R *any](input *T, transform func(*T) (*R, error)) (*R, error) {
	return transform(input)
}
