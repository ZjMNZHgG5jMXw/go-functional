package functional

import (
	"reflect"
)

var (
	// AssignComposed assignes to the function pointer in the first argument
	// the composition of the functions in the remaining arguments.
	AssignComposed func(sun, fun, gun interface{}, runs ...interface{})
)

func init() {
	MakeAssignment(&AssignComposed, Compose)
}

// Compose returns the composition of the functions.
//
// If one of the function returns an error, the error is passed.
// If the error in the return values of one function is non-nil,
// the remaining functions are not called and the error is returned.
func Compose(fun, gun interface{}, runs ...interface{}) (sun interface{}) {
	n := len(runs)
	switch n {
	case 0:
		return compose(fun, gun)
	case 1:
		return compose(fun, compose(gun, runs[0]))
	default:
		return Compose(fun, gun, append(runs[0:n-2], compose(runs[n-2], runs[n-1]))...)
	}
}

// compose :: (b -> c) -> (a -> b) -> a -> c
func compose(fun, gun interface{}) (sun interface{}) {
	var (
		funT = reflect.TypeOf(fun)
		gunT = reflect.TypeOf(gun)
		errT = reflect.TypeOf((*error)(nil)).Elem()
	)

	// create function composition
	var (
		ins  []reflect.Type
		outs []reflect.Type
	)
	for i := 0; i < gunT.NumIn(); i++ {
		ins = append(ins, gunT.In(i))
	}
	for i := 0; i < funT.NumOut(); i++ {
		outs = append(outs, funT.Out(i))
	}

	// pass on error, if exists
	if !funT.Out(funT.NumOut()-1).Implements(errT) && gunT.NumOut() > 1 {
		outs = append(outs, errT)
	}

	sunT := reflect.FuncOf(ins, outs, gunT.IsVariadic())
	sun = reflect.MakeFunc(sunT,
		func(args []reflect.Value) (suns []reflect.Value) {
			os := reflect.ValueOf(gun).Call(args)
			// don't compose in case of error, return it immediately
			if len(os) > 1 && !os[1].IsNil() {
				for _, t := range outs[0 : len(outs)-1] {
					suns = append(suns, reflect.Zero(t))
				}
				suns = append(suns, os[1])
				return
			}
			suns = reflect.ValueOf(fun).Call(os[0:1])
			// append nil error value, if necessary
			if len(suns) < len(outs) {
				suns = append(suns, reflect.Zero(errT))
			}
			return
		}).Interface()
	return
}
