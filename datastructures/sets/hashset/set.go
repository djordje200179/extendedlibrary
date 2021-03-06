package hashset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Set[T comparable] hashmap.Map[T, bool]

func New[T comparable]() Set[T] { return Set[T](hashmap.New[T, bool]()) }

func Collector[T comparable]() streams.Collector[T, sets.Set[T]] { return sets.Collector[T](New[T]()) }

func (set Set[T]) m() hashmap.Map[T, bool] { return hashmap.Map[T, bool](set) }

func (set Set[T]) Size() int { return set.m().Size() }

func (set Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.m().Set(value, true)
	}
}

func (set Set[T]) Remove(value T)        { set.m().Remove(value) }
func (set Set[T]) Contains(value T) bool { return set.m().Contains(value) }

func (set Set[T]) Clear() { set.m().Clear() }

func (set Set[T]) Clone() sets.Set[T] {
	cloned := set.m().Clone().(hashmap.Map[T, bool])
	return Set[T](cloned)
}

func (set Set[T]) Iterator() collections.Iterator[T] { return iterator[T]{set.m().Iterator()} }

func (set Set[T]) Stream() streams.Stream[T] {
	return streams.Map(set.m().Stream(), func(pair misc.Pair[T, bool]) T { return pair.First })
}
