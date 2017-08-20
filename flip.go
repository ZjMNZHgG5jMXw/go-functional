package functional

import (
	"reflect"
)

var (
	// AssignFlipped assignes to the function pointer in the first argument
	// the flipped version of the function in the second argument.
	AssignFlipped func(interface{}, interface{})
)

func init() {
	MakeAssignment(&AssignFlipped, Flip)
}

// Flip expects a function as argument that returns exactly one function
// and returns a function in which the argument lists of the inputi
// functions are flipped.
func Flip(fun interface{}) (sun interface{}) {

	var (
		funT      = reflect.TypeOf(fun)
		gunT      = funT.Out(0)
		firstIns  []reflect.Type
		firstOuts []reflect.Type
		lastIns   []reflect.Type
		lastOuts  []reflect.Type
	)

	for i := 0; i < funT.NumIn(); i++ {
		lastIns = append(lastIns, funT.In(i))
	}

	for i := 0; i < gunT.NumOut(); i++ {
		lastOuts = append(lastOuts, gunT.Out(i))
	}

	for i := 0; i < gunT.NumIn(); i++ {
		firstIns = append(firstIns, gunT.In(i))
	}

	firstOuts = []reflect.Type{
		reflect.FuncOf(
			lastIns,
			lastOuts,
			funT.IsVariadic())}

	sunT := reflect.FuncOf(
		firstIns,
		firstOuts,
		gunT.IsVariadic())

	sun = reflect.
		MakeFunc(sunT,
			func(firstArgs []reflect.Value) []reflect.Value {
				return []reflect.Value{
					reflect.MakeFunc(firstOuts[0],
						func(lastArgs []reflect.Value) []reflect.Value {
							return reflect.
								ValueOf(fun).
								Call(lastArgs)[0].
								Call(firstArgs)
						})}
			}).
		Interface()

	return
}
