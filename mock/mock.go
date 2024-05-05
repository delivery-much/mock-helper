package mock

import "testing"

// Mock represents a mock and its use information
type Mock struct {
	responses map[string]methodResponse
	calls     []MockCall
}

// NewMock returns a new mock struct
func NewMock() Mock {
	return Mock{
		responses: make(map[string]methodResponse),
	}
}

// SetMethodResponse sets a response that the mock will return
// when calling the method specified in the methodName
//
// Its imperative that the response values specified are
// of the same type and are in the same order as the method
// response specified in the method signature
func (mock *Mock) SetMethodResponse(methodName string, response ...any) {
	if mock.responses != nil {
		mock.responses[methodName] = response
	}
}

// GetMethodResponse gets the specified response for a method
func (mock *Mock) GetMethodResponse(methodName string, args ...any) (res methodResponse) {
	key := mountResponseKey(methodName, args...)

	res = mock.responses[key]
	if res.IsEmpty() {
		res = mock.responses[methodName]
	}

	return
}

// RegisterMethodCall registers a method call on a mock given the method name
// and the call arguments
func (mock *Mock) RegisterMethodCall(methodName string, args ...any) {
	mock.calls = append(mock.calls, MockCall{
		MethodName: methodName,
		Args:       args,
	})
}

// GetResponseAndRegister it's equivalent of calling RegisterMethodCall and GetMethodResponse subsequently.
//
// It gets the specified response for a method, given the method name and the args,
// and also registers a method call given those args.
func (mock *Mock) GetResponseAndRegister(methodName string, args ...any) (res methodResponse) {
	mock.RegisterMethodCall(methodName, args...)

	return mock.GetMethodResponse(methodName, args...)
}

// GetCalls returns the mock calls
func (mock *Mock) GetCalls() []MockCall {
	return mock.calls
}

// Called returns if the mock was called
func (mock *Mock) Called() bool {
	return len(mock.calls) > 0
}

// CalledOnce returns if a mock was called exactly once
func (mock *Mock) CalledOnce() bool {
	return len(mock.calls) == 1
}

// CalledTimes returns if a mock was called 'n' times
func (mock *Mock) CalledTimes(n int) bool {
	return len(mock.calls) == n
}

// CalledWith returns if the mock was called at least once with the specified arguments
func (mock *Mock) CalledWith(args ...any) bool {
	return checkCalledWith(mock.calls, args...)
}

// CalledWithExactly returns if the mock was called at least once with exactly the specified arguments,
// with the same values and in the same order
func (mock *Mock) CalledWithExactly(args ...any) bool {
	return checkCalledWithExactly(mock.calls, args...)
}

// Reset resets a mock to an empty state
func (mock *Mock) Reset() {
	*(mock) = NewMock()
}

// Method filters the mock use information for a specific method
func (mock *Mock) Method(name string) *method {
	return &method{
		name,
		mock,
	}
}

func (mock *Mock) Assert(t *testing.T) *mockAssertion {
	return &mockAssertion{
		t: t,
		m: mock,
	}
}
