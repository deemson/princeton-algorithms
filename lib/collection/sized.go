package collection

type Sized interface {
	Size() int
}

func IsEmpty(sized Sized) bool {
	return sized.Size() == 0
}
