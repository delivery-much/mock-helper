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

// argsAreEqual matches two mock arguments to see if they are equal.
// matchArg its the value to match.
// usedArg its the argument that was actually used in the mock call
func argsAreEqual(matchArg, usedArg any) bool {
	matcher, ok := matchArg.(ArgumentMatcher)
	if ok {
		return matcher.Match(usedArg)
	}

	return reflect.DeepEqual(matchArg, usedArg)
}
