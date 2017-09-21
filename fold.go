package functional

import (
	"reflect"
)

// MakeFoldr creates a right fold function
func MakeFoldr(fun interface{}) {
	funT := reflect.TypeOf(fun).Elem()
	gunT := funT.In(0)
	sunT := reflect.FuncOf(
		[]reflect.Type{
			reflect.FuncOf(
				[]reflect.Type{gunT.In(0), gunT.In(1)},
				[]reflect.Type{gunT.In(1)},
				false),
			gunT.In(1),
			reflect.SliceOf(gunT.In(0))},
		[]reflect.Type{gunT.In(1)},
		false)

	sunV := reflect.MakeFunc(
		sunT,
		func(args []reflect.Value) []reflect.Value {
			res := args[1]
			for i := args[2].Len() - 1; i >= 0; i-- {
				res = args[0].Call([]reflect.Value{args[2].Index(i), res})[0]
			}
			return []reflect.Value{res}
		})
	reflect.ValueOf(fun).Elem().Set(sunV)
}

// MakeFoldl creates a left fold function
func MakeFoldl(fun interface{}) {
	funT := reflect.TypeOf(fun).Elem()
	gunT := funT.In(0)
	sunT := reflect.FuncOf(
		[]reflect.Type{
			reflect.FuncOf(
				[]reflect.Type{gunT.In(0), gunT.In(1)},
				[]reflect.Type{gunT.In(0)},
				false),
			gunT.In(0),
			reflect.SliceOf(gunT.In(1))},
		[]reflect.Type{gunT.In(0)},
		false)

	sunV := reflect.MakeFunc(
		sunT,
		func(args []reflect.Value) []reflect.Value {
			res := args[1]
			for i := 0; i < args[2].Len(); i++ {
				res = args[0].Call([]reflect.Value{res, args[2].Index(i)})[0]
			}
			return []reflect.Value{res}
		})
	reflect.ValueOf(fun).Elem().Set(sunV)
}
