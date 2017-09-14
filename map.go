package functional

import (
	"reflect"
)

var (
	AssignMap func(sun, fun interface{})
)

func init() {
	MakeAssignment(&AssignMap, Map)
}

func MakeMap(fun interface{}) {
	reflect.ValueOf(fun).Elem().Set(makeMap(reflect.TypeOf(fun).Elem()))
}

func Map(fun interface{}) interface{} {
	funT := reflect.TypeOf(fun)
	runT := reflect.FuncOf(
		[]reflect.Type{
			reflect.FuncOf(
				[]reflect.Type{funT.In(0)},
				[]reflect.Type{funT.Out(0)},
				false),
			reflect.SliceOf(funT.In(0))},
		[]reflect.Type{reflect.SliceOf(funT.Out(0))},
		false)
	runV := makeMap(runT)

	return reflect.MakeFunc(
		reflect.FuncOf(
			[]reflect.Type{reflect.SliceOf(funT.In(0))},
			[]reflect.Type{reflect.SliceOf(funT.Out(0))},
			false),
		func(args []reflect.Value) []reflect.Value {
			return runV.Call(append([]reflect.Value{reflect.ValueOf(fun)}, args...))
		}).Interface()
}

func makeMap(funT reflect.Type) reflect.Value {
	gunT := funT.In(0)
	sunT := reflect.FuncOf(
		[]reflect.Type{
			reflect.FuncOf(
				[]reflect.Type{gunT.In(0)},
				[]reflect.Type{gunT.Out(0)},
				false),
			reflect.SliceOf(gunT.In(0))},
		[]reflect.Type{reflect.SliceOf(gunT.Out(0))},
		false)

	return reflect.MakeFunc(
		sunT,
		func(args []reflect.Value) []reflect.Value {
			res := reflect.MakeSlice(reflect.SliceOf(gunT.Out(0)), 0, 0)
			for i := 0; i < args[1].Len(); i++ {
				res = reflect.Append(
					res,
					args[0].Call([]reflect.Value{args[1].Index(i)})[0])
			}
			return []reflect.Value{res}
		})
}
