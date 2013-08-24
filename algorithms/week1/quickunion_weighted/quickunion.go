package quickfind

type UF struct {
	Components []int
	Sizes []int
}

func NewUF(size int) *UF {
	uf := UF{
	Components: make([]int, size),
	Sizes: make([]int, size),
	}
	for i := 0; i < size; i++ {
		uf.Components[i] = i
		uf.Sizes[i] = 1
	}
	return &uf
}

func (uf UF) root(i int) int {
	for i != uf.Components[i] {
		// path compression
		uf.Components[i] = uf.Components[uf.Components[i]]
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
	// always connect smaller tree to larger tree
	if uf.Sizes[i] < uf.Sizes[j] {
		uf.Components[i] = j
		uf.Sizes[i] += uf.Sizes[j]
	} else {
		uf.Components[j] = i
		uf.Sizes[j] += uf.Sizes[i]
	}
}
