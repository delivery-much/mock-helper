package mock

// MockCall represents a mock call, with the call arguments
type MockCall struct {
	MethodName string
	Args       []any
}

// HasArgument returns if a mock call arguments contains a specific argument
func (mc *MockCall) HasArgument(arg any) bool {
	for _, a := range mc.Args {
		if argsAreEqual(arg, a) {
			return true
		}
	}

	return false
}
