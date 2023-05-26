package reflection

import "reflect"

func Walk(x interface{}, f func(inp string)) {
	v := getValue(x)

	switch v.Kind() {
	case reflect.String:
		f(v.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			Walk(v.Index(i).Interface(), f)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			Walk(v.Field(i).Interface(), f)
		}
	case reflect.Map:
		for _, iv := range v.MapKeys() {
			Walk(v.MapIndex(iv).Interface(), f)
		}
	case reflect.Chan:
		for mv, ok := v.Recv(); ok; mv, ok = v.Recv() {
			Walk(mv.Interface(), f)
		}
	case reflect.Func:
		retArgs := v.Call(nil)
		for _, mv := range retArgs {
			Walk(mv.Interface(), f)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
