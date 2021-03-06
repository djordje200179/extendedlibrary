package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

func (stream Stream[T]) ForEach(function functions.ParamCallback[T]) {
	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		function(elem.Get())
	}
}

func Reduce[T, P any](stream Stream[T], accumulator P, reducer functions.Reducer[T, P]) P {
	acc := accumulator

	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		acc = reducer(acc, elem.Get())
	}

	return acc
}

func (stream Stream[T]) Any(predictor functions.Predictor[T]) bool {
	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		if predictor(elem.Get()) {
			return true
		}
	}

	return false
}

func (stream Stream[T]) All(predictor functions.Predictor[T]) bool {
	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		if !predictor(elem.Get()) {
			return false
		}
	}

	return true
}

type Collector[T, R any] interface {
	Supply(value T)
	Finish() R
}

func Collect[T, R any](stream Stream[T], collector Collector[T, R]) R {
	stream.ForEach(func(elem T) {
		collector.Supply(elem)
	})

	return collector.Finish()
}

func (stream Stream[T]) Count() int {
	count := 0

	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		count++
	}

	return count
}

func (stream Stream[T]) Max(comparator functions.Comparator[T]) optional.Optional[T] {
	var max T
	set := false

	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		data := elem.Get()
		if !set || comparator(data, max) == comparison.FirstBigger {
			max = data
			set = true
		}
	}

	return optional.New(max, set)
}

func (stream Stream[T]) Min(comparator functions.Comparator[T]) optional.Optional[T] {
	var min T
	set := false

	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		data := elem.Get()
		if !set || comparator(data, min) == comparison.FirstSmaller {
			min = data
			set = true
		}
	}

	return optional.New(min, set)
}

func (stream Stream[T]) First() optional.Optional[T] {
	return stream.supplier.Supply()
}

func (stream Stream[T]) Find(predictor functions.Predictor[T]) optional.Optional[T] {
	for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
		value := elem.Get()
		if predictor(value) {
			return optional.FromValue(value)
		}
	}

	return optional.Empty[T]()
}
