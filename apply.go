package functional

import (
	"reflect"
)

var (
	// AssignApplied assignes to the function pointer in the first argument
	// the application of the third argument to the function in the second
	// argument.
	AssignApplied func(interface{}, interface{}, interface{})
)

func init() {
	MakeAssignment(&AssignApplied, Apply)
}

// Apply calls an unary function with value as the argument and
// returns the first return value as interface{}.
//
// Most useful in combination with Curry().
func Apply(fun interface{}, value interface{}) interface{} {
	return reflect.
		ValueOf(fun).
		Call([]reflect.Value{reflect.ValueOf(value)})[0].
		Interface()
}
