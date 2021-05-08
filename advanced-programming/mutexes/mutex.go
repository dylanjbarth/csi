package mutex

type footex struct {
	locked bool
}

func (f *footex) Lock() {
	if !f.locked {
		f.locked = true
	}
}

func (f *footex) Unlock() {
	if f.locked {
		f.locked = false
	}
}
