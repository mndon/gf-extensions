package slicex

func partition[T ordered](s []T, low, high int) int {
	for j := low; j < high; j++ {
		if s[j] < s[high] {
			s[low], s[j] = s[j], s[low]
			low++
		}
	}
	s[low], s[high] = s[high], s[low]
	return low
}

func quickSort[T ordered](s []T, low, high int) {
	if low < high {
		if high-low < 12 {
			insertionSort(s, low, high)
		} else {
			p := partition(s, low, high)
			quickSort(s, low, p-1)
			quickSort(s, p+1, high)
		}
	}
}

func insertionSort[T ordered](s []T, a, b int) {
	for i := 1; i < b-a+1; i++ {
		j := i
		for j > 0 {
			if s[a+j] < s[a+j-1] {
				s[a+j-1], s[a+j] = s[a+j], s[a+j-1]
			}
			j--
		}
	}
}

func partitionFunc[T any](s []T, low, high int, less func(T, T) bool) int {
	for j := low; j < high; j++ {
		if less(s[j], s[high]) {
			s[low], s[j] = s[j], s[low]
			low++
		}
	}
	s[low], s[high] = s[high], s[low]
	return low
}

func quickSortFunc[T any](s []T, low, high int, less func(T, T) bool) {
	if low < high {
		if high-low < 12 {
			insertionSortFunc(s, low, high, less)
		} else {
			p := partitionFunc(s, low, high, less)
			quickSortFunc(s, low, p-1, less)
			quickSortFunc(s, p+1, high, less)
		}
	}
}

func insertionSortFunc[T any](s []T, a, b int, less func(T, T) bool) {
	for i := 1; i < b-a+1; i++ {
		j := i
		for j > 0 {
			if less(s[a+j], s[a+j-1]) {
				s[a+j-1], s[a+j] = s[a+j], s[a+j-1]
			}
			j--
		}
	}
}
