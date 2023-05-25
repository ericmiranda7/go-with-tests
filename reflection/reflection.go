package reflection

import "reflect"

func Walk(x interface{}, f func(inp string)) {
	v := reflect.ValueOf(x)
	str := v.Field(0)

	f(str.String())
}
