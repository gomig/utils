package utils

// If generate quick if
func If[T any](cond bool, yes T, no T) T {
	if cond {
		return yes
	}
	return no
}

// Try simulate try catch block
func Try(e error, fail func(e error), pass func(), finally ...func(e error)) {
	if e != nil {
		fail(e)
	} else {
		pass()
	}
	if len(finally) > 0 {
		finally[0](e)
	}
}

// Contains check if slice contains  item
func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
