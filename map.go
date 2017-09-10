package functional

import (
	"reflect"
)

func MakeMap(fun interface{}) {
	funT := reflect.TypeOf(fun).Elem()
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

	reflect.ValueOf(fun).Elem().
		Set(
			reflect.MakeFunc(
				sunT,
				func(args []reflect.Value) (results []reflect.Value) {
					res := reflect.MakeSlice(reflect.SliceOf(gunT.Out(0)), 0, 0)
					for i := 0; i < args[1].Len(); i++ {
						res = reflect.Append(
							res,
							args[0].Call([]reflect.Value{args[1].Index(i)})[0])
					}
					return []reflect.Value{res}
				}))
}