package mock

// method represents a mock use information, but filtered for a specific method
type method struct {
	name string
	mock *Mock
}

// SetResponse sets the response that the mock method should return when called
//
// Its imperative that the response values specified are
// of the same type and are in the same order as the method
// response specified in the method signature
func (m *method) SetResponse(response ...any) {
	if m.mock != nil {
		m.mock.responses[m.name] = response
	}
}

// GetResponse gets the specified response for the method
func (m *method) GetResponse() (res methodResponse) {
	if m.mock != nil {
		res = m.mock.responses[m.name]
	}

	return
}

// GetCalls returns the mock method calls
func (m *method) GetCalls() []MockCall {
	calls := []MockCall{}

	if m.mock != nil {
		for _, mockCall := range m.mock.calls {
			if mockCall.MethodName == m.name {
				calls = append(calls, mockCall)
			}
		}
	}

	return calls
}

// Called checks if the mock method was called
func (m *method) Called() bool {
	return len(m.GetCalls()) > 0
}

// CalledOnce checks if a mock method was called exactly once
func (m *method) CalledOnce() bool {
	return len(m.GetCalls()) == 1
}

// CalledTimes returns if a mock method was called 'n' times
func (m *method) CalledTimes(n int) bool {
	return len(m.GetCalls()) == n
}

// CalledWith checks if the mock method was called at least once with the specified arguments
func (m *method) CalledWith(args ...any) bool {
	for _, call := range m.GetCalls() {
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

// CalledWithExactly checks if the mock method was called at least once with exactly the specified arguments,
// with the same values and in the same order
func (m *method) CalledWithExactly(args ...any) bool {
	for _, call := range m.GetCalls() {
		if len(args) != len(call.Args) {
			continue
		}

		hasExactArgs := true
		for i, callArg := range call.Args {
			if args[i] != callArg {
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
