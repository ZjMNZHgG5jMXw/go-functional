package functional

import (
	"reflect"
)

type pointerMode int

const (
	// MaybeMode will copy the value of returned pointers
	MaybeMode pointerMode = iota
	// KeepPointerMode will copy the address of returned pointers
	KeepPointerMode
)

// MakeAssignment transforms fun in an assignment
// and stores it to sun.
func MakeAssignment(sun, fun interface{}, mode ...pointerMode) {
	var (
		funT = reflect.TypeOf(fun)
		ins  []reflect.Type
	)

	// sun's first argument needs to be an interface containing or
	// a pointer to fun's output type
	switch funT.Out(0).Kind() {
	case reflect.Interface:
		ins = append(ins, funT.Out(0))
	case reflect.Ptr:
		if isMaybe(mode) {
			ins = append(ins, funT.Out(0))
		} else {
			ins = append(ins, reflect.PtrTo(funT.Out(0)))
		}
	default:
		ins = append(ins, reflect.PtrTo(funT.Out(0)))
	}

	// followed by the rest of fun's arguments
	for i := 0; i < funT.NumIn(); i++ {
		ins = append(ins, funT.In(i))
	}

	// sun's type signature
	bunT := reflect.FuncOf(ins, nil, funT.IsVariadic())

	// sun's implementation
	bun := reflect.MakeFunc(
		bunT,
		func(args []reflect.Value) (outs []reflect.Value) {
			var res []reflect.Value
			// unpack slice of variadic argument
			if funT.IsVariadic() {
				var (
					// all but the first and the last argument
					inVs = args[1 : len(args)-1]
					// last argument
					variadic = args[len(args)-1]
				)
				// append elements from variadic argument
				for i := 0; i < variadic.Len(); i++ {
					inVs = append(inVs, variadic.Index(i))
				}
				res = reflect.ValueOf(fun).Call(inVs)
			} else {
				// pass all but the first argument
				res = reflect.ValueOf(fun).Call(args[1:])
			}
			switch funT.Out(0).Kind() {
			case reflect.Interface:
				// set interface{} content
				args[0].Elem().Elem().Set(res[0].Elem())
			case reflect.Ptr:
				if isMaybe(mode) {
					if !res[0].IsNil() {
						// set content behind pointer
						args[0].Elem().Set(res[0].Elem())
					}
				} else {
					// set pointer content
					args[0].Elem().Set(res[0])
				}
			default:
				// set pointer content
				args[0].Elem().Set(res[0])
			}
			return
		})

	// assign the assignment function to sun
	reflect.ValueOf(sun).Elem().Set(bun)
}

func isMaybe(mode []pointerMode) bool {
	return len(mode) == 0 || mode[0] != KeepPointerMode
}
