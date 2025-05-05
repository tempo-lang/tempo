package misc

import "fmt"

func JoinStrings[T any](strings []T, delimeter string) string {
	return JoinStringsFunc(strings, delimeter, func(elem T) string {
		return fmt.Sprintf("%v", elem)
	})
}

func JoinStringsFunc[T any](items []T, delimeter string, format func(T) string) string {
	if len(items) == 0 {
		return ""
	}

	result := ""

	for _, s := range items {
		result += format(s) + delimeter
	}
	result = result[:len(result)-len(delimeter)]

	return result
}
