package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func MountArgsAssertionErrMsg(title string, calls []MockCall, expectedArgs ...any) (msg string) {
	msg = title
	for _, expectedArg := range expectedArgs {
		t := reflect.TypeOf(expectedArg).Name()
		msg = fmt.Sprintf("%s  ++ (%s) %v\n", msg, t, expectedArg)
	}

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

func MountMethodAssertionErrMsg(m *method, expectedArgs ...any) string {
	return MountArgsAssertionErrMsg(
		fmt.Sprintf("Failed to assert method calls.\nExpected method %s to be called with: \n", m.name),
		m.GetCalls(),
		expectedArgs...,
	)
}

func MountMockAssertionErrMsg(m *Mock, expectedArgs ...any) string {
	return MountArgsAssertionErrMsg(
		"Failed to assert mock calls.\nExpected mock to be called with: \n",
		m.GetCalls(),
		expectedArgs...,
	)
}

func (m *method) AssertCalledWith(t *testing.T, args ...any) *method {
	if !m.CalledWith(args) {
		t.Fatal(MountMethodAssertionErrMsg(m, args...))
	}

	return m
}

func (m *method) AssertCalledWithExactly(t *testing.T, args ...any) *method {
	if !m.CalledWithExactly(args) {
		t.Fatal(MountMethodAssertionErrMsg(m, args...))
	}

	return m
}

func (m *method) AssertCalled(t *testing.T) *method {
	if !m.Called() {
		t.Fatalf("Failed to assert mock calls.\nMethod %s was not called", m.name)
	}

	return m
}

func (m *method) AssertCalledOnce(t *testing.T) *method {
	if !m.Called() {
		msg := fmt.Sprintf("Failed to assert mock calls.\nExpected method %s to be called once, ", m.name)

		if callsLen := len(m.GetCalls()); callsLen == 0 {
			msg += "but it was not called"
		} else {
			msg += fmt.Sprintf("but it was called %d times", callsLen)
		}
		t.Fatalf(msg)
	}

	return m
}

func (m *method) AssertCalledTimes(t *testing.T, n int) *method {
	if !m.CalledTimes(n) {
		msg := fmt.Sprintf("Failed to assert mock calls.\nExpected method %s to be called %d times, ", m.name, n)

		if callsLen := len(m.GetCalls()); callsLen == 0 {
			msg += "but it was not called"
		} else {
			msg += fmt.Sprintf("but it was called %d times", callsLen)
		}
		t.Fatalf(msg)
	}

	return m
}

func (m *Mock) AssertCalledWith(t *testing.T, args ...any) *Mock {
	if !m.CalledWith(args) {
		t.Fatal(MountMockAssertionErrMsg(m, args...))
	}

	return m
}

func (m *Mock) AssertCalledWithExactly(t *testing.T, args ...any) *Mock {
	if !m.CalledWithExactly(args) {
		t.Fatal(MountMockAssertionErrMsg(m, args...))
	}

	return m
}

func (m *Mock) AssertCalled(t *testing.T) *Mock {
	if !m.Called() {
		t.Fatal("Failed to assert mock calls.\nMock was not called")
	}

	return m
}

func (m *Mock) AssertCalledOnce(t *testing.T) *Mock {
	if !m.Called() {
		msg := "Failed to assert mock calls.\nExpected mock to be called once, "

		if callsLen := len(m.GetCalls()); callsLen == 0 {
			msg += "but it was not called"
		} else {
			msg += fmt.Sprintf("but it was called %d times", callsLen)
		}
		t.Fatalf(msg)
	}

	return m
}

func (m *Mock) AssertCalledTimes(t *testing.T, n int) *Mock {
	if !m.CalledTimes(n) {
		msg := fmt.Sprintf("Failed to assert mock calls.\nExpected mock to be called %d times, ", n)

		if callsLen := len(m.GetCalls()); n == 0 {
			msg += "but it was not called"
		} else {
			msg += fmt.Sprintf("but it was called %d times", callsLen)
		}
		t.Fatalf(msg)
	}

	return m
}
