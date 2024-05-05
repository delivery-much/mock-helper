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

// utility function to mount the error message when asserting mock or method calls
func mountArgsAssertionErrMsg(title string, calls []MockCall, expectedArgs ...any) (msg string) {
	msg = title

	expectedArgsStr := ""
	for _, expectedArg := range expectedArgs {
		t := reflect.TypeOf(expectedArg).Name()
		expectedArgsStr = fmt.Sprintf("%s  ++ (%s) %v\n", expectedArgsStr, t, expectedArg)
	}
	if len(expectedArgsStr) == 0 {
		expectedArgsStr = "  ++ (no arguments)\n"
	}

	msg = fmt.Sprintf("%s%s", msg, expectedArgsStr)

	if len(calls) == 0 {
		msg = fmt.Sprintf("%s\nBut it was not called\n", msg)
		return
	}

	msg = fmt.Sprintf("%s\nActual calls:\n", msg)
	for i, call := range calls {
		msg = fmt.Sprintf("%s[%d]:\n", msg, i+1)
		if len(call.Args) == 0 {
			msg = fmt.Sprintf("%s  -- (no arguments)\n", msg)
			continue
		}

		for _, callArg := range call.Args {
			t := reflect.TypeOf(callArg).String()
			msg = fmt.Sprintf("%s  -- (%s) %v\n", msg, t, callArg)
		}
	}

	return
}

// checkCalledWith it's a common implementation between the mock and method structs.
// it checks if any of the mock or method calls have the specified arguments
func checkCalledWith(calls []MockCall, args ...any) bool {
	if len(args) == 0 {
		for _, call := range calls {
			if len(call.Args) == 0 {
				return true
			}
		}

		return false
	}

	for _, call := range calls {
		hasArgs := true
		for _, arg := range args {
			if !call.HasArgument(arg) {
				hasArgs = false
				break
			}
		}

		if hasArgs {
			return true
		}
	}

	return false
}

// checkCalledWithExactly it's a common implementation between the mock and method structs.
// it checks if any of the mock or method calls have exactly the specified arguments
func checkCalledWithExactly(calls []MockCall, args ...any) bool {
	if len(args) == 0 {
		for _, call := range calls {
			if len(call.Args) == 0 {
				return true
			}
		}

		return false
	}

	for _, call := range calls {
		if len(args) != len(call.Args) {
			continue
		}

		hasExactArgs := true
		for i, callArg := range call.Args {
			if !argsAreEqual(args[i], callArg) {
				hasExactArgs = false
				break
			}
		}

		if hasExactArgs {
			return true
		}
	}

	return false
}
