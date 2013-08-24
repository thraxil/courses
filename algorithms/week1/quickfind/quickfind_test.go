package quickfind

import "testing"

func Test_Quickfind(t *testing.T) {
	uf := NewUF(10)
	if uf.Connected(0, 1) {
		t.Error("no connections yet")
	}
	uf.Union(0, 1)
	if !uf.Connected(0, 1) {
		t.Error("now it should be connected")
	}

}
