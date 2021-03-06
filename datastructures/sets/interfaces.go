package sets

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Set[T comparable] interface {
	Size() int

	Add(value T)
	Remove(value T)
	Contains(value T) bool

	Clear()
	misc.Cloner[Set[T]]

	collections.Iterable[T]
	streams.Streamer[T]
}
