package mock

import (
	"fmt"
	"reflect"
)

// mountResponseKey mounts the string key that should be used to access the mock responses map,
// based on the method name and the args
func mountResponseKey(name string, args ...any) (res string) {
	res = name
	for _, p := range args {
		res = fmt.Sprintf("%s-%v:%v", res, reflect.TypeOf(p), p)
	}

	return
}
