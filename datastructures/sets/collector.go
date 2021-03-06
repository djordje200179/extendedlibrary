package sets

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T comparable] struct {
	set Set[T]
}

func Collector[T comparable](empty Set[T]) streams.Collector[T, Set[T]] {
	return collector[T]{empty}
}

func (collector collector[T]) Supply(value T) { collector.set.Add(value) }

func (collector collector[T]) Finish() Set[T] { return collector.set }
