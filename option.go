// Package opt  provides the specification and implementation of golang's optional values.
package opt

// Option defines the specification for a golang optional value.
type Option[T any] interface {

	// Get returns value of type T if present or panic if absent
	Get() T

	// GetOrElse returns value of type T if present else return the value passed
	// as a parameter.
	GetOrElse(T) T

	// GetOrZero returns value of type T if present else returns the zero value of T.
	GetOrZero() T

	// Empty returns true if value is absent else returns false.
	Empty() bool
}

// some implements Option when value is present
type some[T any] struct {
	val T
}

// Get returns the value contained in some.
func (s some[T]) Get() T {
	return s.val
}

// GetOrElse returns the value contained in some.
func (s some[T]) GetOrElse(_ T) T {
	return s.val
}

// GetOrZero returns the value contained in some.
func (s some[T]) GetOrZero() T {
	return s.val
}

// Empty returns false for all instances of some.
func (s some[T]) Empty() bool {
	return false
}

// Some returns an Option with value of type T present.
func Some[T any](val T) Option[T] {
	return some[T]{
		val: val,
	}
}

// none implements Option for absent values.
type none[T any] struct{}

// Get panics for none values.
func (n none[T]) Get() T {
	panic("None has no value")
}

// GetOrElse returns the value passed in the parameter.
func (n none[T]) GetOrElse(val T) T {
	return val
}

// GetOrZero returns the zero value of T.
func (n none[T]) GetOrZero() T {
	var t T
	return t
}

// Empty returns true for all instances of none.
func (n none[T]) Empty() bool {
	return true
}

// None returns the absent value for some type T.
func None[T any]() Option[T] {
	return none[T]{}
}

// Map convert an Option of type T to an Option of type K.
func Map[T, K any](mapper func(t T) K, opt Option[T]) Option[K] {
	if opt.Empty() {
		return None[K]()
	}
	value := opt.Get()
	return Some(mapper(value))
}
