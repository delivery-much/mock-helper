package mock

import (
	"fmt"
	"testing"
)

func mountMockArgAssertionErrMsg(ma *mockAssertion, expectedArgs ...any) string {
	verb := "to be"
	if ma.negation {
		verb = "not to be"
	}
	return mountArgsAssertionErrMsg(
		fmt.Sprintf("Failed to assert mock call arguments.\nExpected mock %s called with: \n", verb),
		ma.m.GetCalls(),
		expectedArgs...,
	)
}

func mountMockCallAssertionErrMsg(ma *mockAssertion, expectedCallN int) (msg string) {
	verb := "to be"
	if ma.negation {
		verb = "not to be"
	}

	times := "once"
	if expectedCallN > 1 {
		times = fmt.Sprintf("%d times", expectedCallN)
	}

	msg = fmt.Sprintf("Failed to assert mock calls.\nExpected mock %s called %s, ", verb, times)
	if ma.negation {
		msg += "but it was"
		return
	}

	switch callsLen := len(ma.m.GetCalls()); callsLen {
	case 0:
		msg += "but it was not called"
	case 1:
		msg += "but it was called once"
	default:
		msg += fmt.Sprintf("but it was called %d times", callsLen)
	}

	return
}

type mockAssertion struct {
	t        *testing.T
	m        *Mock
	negation bool
}

// Not sets the mock assertion as a negation.
//
// When this method is called, the NEGATION of the subsequent assertion will be validated.
func (ma *mockAssertion) Not() *mockAssertion {
	ma.negation = true
	return ma
}

func (ma *mockAssertion) verify(cond bool) bool {
	if ma.negation {
		return !cond
	}

	return cond
}

// CalledWith asserts that the mock was called at least once with the specified arguments
func (ma *mockAssertion) CalledWith(args ...any) *finishedMockAssertion {
	failureCond := !ma.m.CalledWith(args...)
	if ma.verify(failureCond) {
		ma.t.Error(mountMockArgAssertionErrMsg(ma, args...))
	}

	return &finishedMockAssertion{ma}
}

// CalledWithExactly asserts that the mock was called at least once with exactly the specified arguments,
// with the same values and in the same order
func (ma *mockAssertion) CalledWithExactly(args ...any) *finishedMockAssertion {
	failureCond := !ma.m.CalledWithExactly(args...)
	if ma.verify(failureCond) {
		ma.t.Error(mountMockArgAssertionErrMsg(ma, args...))
	}

	return &finishedMockAssertion{ma}
}

// Called asserts that the mock was called at least once
func (ma *mockAssertion) Called() *finishedMockAssertion {
	wasCalled := ma.m.Called()
	failureCond := !wasCalled
	if ma.verify(failureCond) {
		verb := "to be"
		if ma.negation {
			verb = "not to be"
		}

		sufix := "but it wasn't"
		if wasCalled {
			sufix = "but it was"
		}
		ma.t.Errorf("Failed to assert mock calls.\nExpected mock %s called, %s", verb, sufix)
	}

	return &finishedMockAssertion{ma}
}

// Called asserts that the mock was called exaclty once
func (ma *mockAssertion) CalledOnce() *finishedMockAssertion {
	failureCond := !ma.m.CalledOnce()
	if ma.verify(failureCond) {
		msg := mountMockCallAssertionErrMsg(ma, 1)
		ma.t.Errorf(msg)
	}

	return &finishedMockAssertion{ma}
}

// Called asserts that the mock was called 'n' times
func (ma *mockAssertion) CalledTimes(n int) *finishedMockAssertion {
	failureCond := !ma.m.CalledTimes(n)
	if ma.verify(failureCond) {
		msg := mountMockCallAssertionErrMsg(ma, n)
		ma.t.Errorf(msg)
	}

	return &finishedMockAssertion{ma}
}

type finishedMockAssertion struct {
	ma *mockAssertion
}

// And is used to chain mock assertions
func (fma *finishedMockAssertion) And() *mockAssertion {
	return &mockAssertion{
		m: fma.ma.m,
		t: fma.ma.t,
	}
}
