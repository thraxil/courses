package quickfind

type UF struct {
	Components []int
}

func NewUF(size int) *UF {
	uf := UF{Components: make([]int, size)}
	for i := 0; i < size; i++ {
		uf.Components[i] = i
	}
	return &uf
}

func (uf UF) Connected(p, q int) bool {
	return uf.Components[p] == uf.Components[q]
}

func (uf *UF) Union(p, q int) {
	pid := uf.Components[p]
	qid := uf.Components[q]
	for i := range uf.Components {
		if uf.Components[i] == pid {
			uf.Components[i] = qid
		}
	}
}
