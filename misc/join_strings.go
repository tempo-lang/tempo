package misc

import "fmt"

func JoinStrings[T any](strings []T, delimeter string) string {
	if len(strings) == 0 {
		return ""
	}

	result := ""

	for _, s := range strings {
		result += fmt.Sprintf("%v%s", s, delimeter)
	}
	result = result[:len(result)-len(delimeter)]

	return result
}
