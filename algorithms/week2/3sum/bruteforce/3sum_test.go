package threesum
import "testing"

func Test_3sum(t *testing.T) {
	a1 := []int{1,2,3}
	if countsums(a1) != 0 {
		t.Error("shouldn't be any")
	}
	a2 := []int{1,2,3,-3}
	if countsums(a2) != 1 {
		t.Error("should be exactly 1")
	}
}
