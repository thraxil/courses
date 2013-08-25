package binarysearch

const notfound = -1

func search(a []int, v int) int {
	if len(a) == 0 {
		return notfound
	}
	lo := 0
	hi := len(a) - 1
	mid := ((hi - lo) / 2) + lo

	for lo != hi {
		if a[mid] == v {
			return mid
		}
		if v < a[mid] {
			// left
			hi = mid - 1
		} else {
			// right
			lo = mid + 1
		}
		mid = ((hi - lo) / 2) + lo
	}
	if a[mid] == v {
		return mid
	}
	// if we reach the end, we know we are done
	return notfound

}
