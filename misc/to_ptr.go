package misc

// toPtr converts a value to a pointer to the value
func ToPtr[T any](val T) *T {
	return &val
}
