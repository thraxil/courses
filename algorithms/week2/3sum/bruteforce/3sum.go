package threesum

func countsums(a []int) int {
	N := len(a)
	count := 0
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			for k := j+1; k < N; k++ {
				if (a[i] + a[j] + a[k]) == 0 {
					count = count + 1
				}
			}
		}
	}
	return count
}

