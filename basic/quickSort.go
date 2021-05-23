package basic

func quickSort(q []int, l int, r int) {
	if l >= r {
		return
	}

	i, j := l-1, r+1
	x := q[l]
	for i < j {
		i++
		for q[i] < x {
			i++
		}
		j--
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	// [l,j] 和 [j+1,r] 的 pivot(line11) 不能选择 q[r]，会有边界问题，导致死循环，和数组中任意其他位置均可
	// 循环不变量:
	// [l,i] 区间若合法则其中所有值 <= pivot
	// [j,r] 区间若合法则其中所有值 >= pivot
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}
