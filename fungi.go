package fungi

// Process executes a series of functions and stops at the first error encountered.
// Returns nil if all functions succeed.
func Process(funcs ...func() error) error {
	for _, fn := range funcs {
		if err := fn(); err != nil {
			return err // Return the first error encountered
		}
	}
	return nil // No errors found, return nil
}

// Iter iterates over a slice, calling a callback function for each element.
// Stops and returns the first error encountered, or nil if all succeed.
func Iter[T any](someSlice []T, someCallback func(i int, item T) error) error {
	for i, item := range someSlice {
		err := someCallback(i, item)
		if err != nil {
			return err
		}
	}
	return nil
}

// Map transforms a slice into a new slice by applying a transformation function
// to each element, passing its index and value.
func Map[T any, U any](someSlice []T, transform func(i int, item T) U) []U {
	result := make([]U, len(someSlice))
	for i, item := range someSlice {
		result[i] = transform(i, item)
	}
	return result
}

// Filter returns a new slice containing only the elements of the original slice
// that satisfy the predicate function, which receives the index and value.
func Filter[T any](someSlice []T, predicate func(i int, item T) bool) []T {
	var result []T
	for i, item := range someSlice {
		if predicate(i, item) {
			result = append(result, item)
		}
	}
	return result
}

// Reduce aggregates a slice into a single value by applying a reducer function
// to each element along with an accumulator, starting with the initial value.
func Reduce[T any, U any](someSlice []T, initial U, reducer func(i int, acc U, item T) U) U {
	acc := initial
	for i, item := range someSlice {
		acc = reducer(i, acc, item)
	}
	return acc
}

// Find searches a slice for the first element that satisfies the predicate function,
// returning the element and true if found, or the zero value and false if not.
func Find[T any](someSlice []T, predicate func(i int, item T) bool) (T, bool) {
	var zero T
	for i, item := range someSlice {
		if predicate(i, item) {
			return item, true
		}
	}
	return zero, false
}

// Every checks whether all elements in a slice satisfy the predicate function.
// Returns true if all elements satisfy the predicate, or false otherwise.
func Every[T any](someSlice []T, predicate func(i int, item T) bool) bool {
	for i, item := range someSlice {
		if !predicate(i, item) {
			return false
		}
	}
	return true
}

// Some checks whether at least one element in a slice satisfies the predicate function.
// Returns true if any element satisfies the predicate, or false otherwise.
func Some[T any](someSlice []T, predicate func(i int, item T) bool) bool {
	for i, item := range someSlice {
		if predicate(i, item) {
			return true
		}
	}
	return false
}
