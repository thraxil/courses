package binarysearch

const notfound = -1

func search(a []int, v int) int {
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := ((hi - lo) / 2) + lo
		if v < a[mid] {
			// left
			hi = mid - 1
		} else if v > a[mid] {
			// right
			lo = mid + 1
		} else {
			return mid
		}
	}
	// if we reach the end, we know we are done
	return notfound

}
