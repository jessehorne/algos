package algos

func iterative_binarysearch_int(sorted []int, value int) int {
	low := 0
	high := len(sorted) - 1

	for {

		mid := (low + high) / 2

		if value == sorted[mid] {
			return mid
		} else if low == high {
			// if low and high are equal and the above value hasn't been found, then there's nothing else we can do
			return -1
		} else if value > sorted[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}
