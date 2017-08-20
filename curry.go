package functional

import (
	"reflect"
)

var (
	// AssignCurried assignes to the function pointer in the first argument
	// the curried version of the function in the second argument.
	AssignCurried func(interface{}, interface{})
	// AssignUncurried assignes to the function pointer in the first argument
	// the uncurried version of the function in the second argument.
	AssignUncurried func(interface{}, interface{})
)

func init() {
	MakeAssignment(&AssignCurried, Curry)
	MakeAssignment(&AssignUncurried, Uncurry)
}

// Curry is a kind of a curry with type signature ((a, b) -> c) -> a -> b -> c,
// but the (a, b)-part of the argument can be extended to an n-tuple
// for an arbitrary number n, e. g., ((a, b, c) -> d) -> a -> (b, c) -> d.
func Curry(fun interface{}) (curried interface{}) {
	funT := reflect.TypeOf(fun)

	// create a new function type with the first argument removed
	var ins, outs []reflect.Type
	for i := 1; i < funT.NumIn(); i++ {
		ins = append(ins, funT.In(i))
	}
	for i := 0; i < funT.NumOut(); i++ {
		outs = append(outs, funT.Out(i))
	}
	innerT := reflect.FuncOf(ins, outs, funT.IsVariadic())

	// create a new function type that takes the first argument
	// and returns a function of the previous type definition
	curriedT := reflect.FuncOf([]reflect.Type{funT.In(0)}, []reflect.Type{innerT}, false)
	curried = reflect.
		MakeFunc(
			curriedT,
			func(args []reflect.Value) []reflect.Value {
				return []reflect.Value{reflect.MakeFunc(
					innerT,
					func(innerArgs []reflect.Value) []reflect.Value {
						newArgs := []reflect.Value{args[0]}
						newArgs = append(newArgs, innerArgs...)
						return reflect.ValueOf(fun).Call(newArgs)
					})}
			}).
		Interface()
	return
}

// (a -> b -> c) -> (a, b) -> c
func Uncurry(fun interface{}) (uncurried interface{}) {
	// check function type
	funT := reflect.TypeOf(fun)
	gunT := funT.Out(0)

	// create a new function type
	var ins, outs []reflect.Type
	for i := 0; i < funT.NumIn(); i++ {
		ins = append(ins, funT.In(i))
	}
	for i := 0; i < gunT.NumIn(); i++ {
		ins = append(ins, gunT.In(i))
	}
	for i := 0; i < gunT.NumOut(); i++ {
		outs = append(outs, gunT.Out(i))
	}

	uncurriedT := reflect.FuncOf(ins, outs, gunT.IsVariadic())
	uncurried = reflect.
		MakeFunc(
			uncurriedT,
			func(args []reflect.Value) []reflect.Value {
				res := reflect.ValueOf(fun).Call(args[:funT.NumIn()])
				return res[0].Call(args[funT.NumIn():])
			}).
		Interface()
	return
}
