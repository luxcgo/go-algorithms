package basic

// https://golang.org/doc/effective_go#commentary

// findFirstPostition finds first position x exists in an
// array q sorted in ascending order.
func findFirstPostition(q []int, x int) int {
	if len(q) == 0 {
		return -1
	}
	l := 0
	r := len(q) - 1
	for l+1 < r {
		mid := (l + r) >> 1
		if q[mid] >= x {
			r = mid
		} else {
			l = mid
		}
	}
	if q[l] == x {
		return l
	}
	if q[r] == x {
		return r
	}
	return -1
}

// findLastPostition finds last position x exists in an
// array q sorted in ascending order.
func findLastPostition(q []int, x int) int {
	if len(q) == 0 {
		return -1
	}
	l := 0
	r := len(q) - 1
	for l+1 < r {
		mid := (l + r) >> 1
		if q[mid] > x {
			r = mid
		} else {
			l = mid
		}
	}
	if q[r] == x {
		return r
	}
	if q[l] == x {
		return l
	}
	return -1
}
