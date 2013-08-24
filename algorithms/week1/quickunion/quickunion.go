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

func (uf UF) root(i int) int {
	for i != uf.Components[i] {
		i = uf.Components[i]
	}
	return i
}

func (uf UF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf *UF) Union(p, q int) {
	i := uf.root(p)
	j := uf.root(q)
	uf.Components[i] = j
}
