package quicksort

func QuickSort(a []int) {
	l := 0
	h := len(a) - 1

	if l < h {
		p := partition(a)
		QuickSort(a[:p])
		QuickSort(a[p:])
	}
}

func partition(a []int) int {
	l := 0
	h := len(a) - 1

	i := l - 1
	piv := a[h]

	for j := l; j < h; j++ {
		if a[j] <= piv {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[h] = a[h], a[i+1]
	return i + 1
}
