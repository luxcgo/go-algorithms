package basic

var tmp []int

func mergeSort(q []int, l int, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1

	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)

	k := 0
	i := l
	j := mid + 1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			i++
		} else {
			tmp[k] = q[j]
			j++
		}
		k++
	}
	copy(tmp[k:], q[i:mid+1])
	copy(tmp[k:], q[j:r+1])

	copy(q[l:r+1], tmp)
}
