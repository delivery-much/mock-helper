package mock

// MockCall represents a mock call, with the call arguments
type MockCall struct {
	MethodName string
	Args       []any
}

// HasArgument checks if a mock call arguments contains a specific argument
func (mc *MockCall) HasArgument(arg any) bool {
	for _, a := range mc.Args {
		if a == arg {
			return true
		}
	}

	return false
}
