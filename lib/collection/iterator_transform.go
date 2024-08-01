package collection

func TransformIterator[I, O any](iterator Iterator[I], transform func(I) O) *TransformedIterator[I, O] {
	return &TransformedIterator[I, O]{
		iterator:  iterator,
		transform: transform,
	}
}

type TransformedIterator[I, O any] struct {
	iterator  Iterator[I]
	transform func(I) O
}

func (i *TransformedIterator[I, O]) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransformedIterator[I, O]) Next() O {
	return i.transform(i.iterator.Next())
}
