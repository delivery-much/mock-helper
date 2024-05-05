package mock

import (
	"fmt"
	"testing"
)

func mountMethodArgAssertionErrMsg(ma *methodAssertion, expectedArgs ...any) string {
	verb := "to be"
	if ma.negation {
		verb = "not to be"
	}
	return mountArgsAssertionErrMsg(
		fmt.Sprintf("Failed to assert method call arguments.\nExpected method %s %s called with: \n", ma.m.name, verb),
		ma.m.GetCalls(),
		expectedArgs...,
	)
}

func mountMethodCallAssertionErrMsg(ma *methodAssertion, expectedCallN int) (msg string) {
	verb := "to be"
	if ma.negation {
		verb = "not to be"
	}

	times := "once"
	if expectedCallN > 1 {
		times = fmt.Sprintf("%d times", expectedCallN)
	}

	msg = fmt.Sprintf("Failed to assert method calls.\nExpected method %s %s called %s, ", ma.m.name, verb, times)
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

type methodAssertion struct {
	t        *testing.T
	m        *method
	negation bool
}

func (ma *methodAssertion) verify(cond bool) bool {
	if ma.negation {
		return !cond
	}

	return cond
}

// Not sets the method assertion as a negation.
//
// When this method is called, the NEGATION of the subsequent assertion will be validated.
func (ma *methodAssertion) Not() *methodAssertion {
	ma.negation = true
	return ma
}

// CalledWith asserts that the method was called at least once with the specified arguments
func (ma *methodAssertion) CalledWith(args ...any) *finishedMethodAssertion {
	failureCond := !ma.m.CalledWith(args...)
	if ma.verify(failureCond) {
		ma.t.Error(mountMethodArgAssertionErrMsg(ma, args...))
	}

	return &finishedMethodAssertion{ma}
}

// CalledWithExactly asserts that the method was called at least once with exactly the specified arguments,
// with the same values and in the same order
func (ma *methodAssertion) CalledWithExactly(args ...any) *finishedMethodAssertion {
	failureCond := !ma.m.CalledWithExactly(args...)
	if ma.verify(failureCond) {
		ma.t.Error(mountMethodArgAssertionErrMsg(ma, args...))
	}

	return &finishedMethodAssertion{ma}
}

// Called asserts that the method was called at least once
func (ma *methodAssertion) Called() *finishedMethodAssertion {
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
		ma.t.Errorf("Failed to assert method calls.\nExpected method %s %s called, %s", ma.m.name, verb, sufix)
	}

	return &finishedMethodAssertion{ma}
}

// Called asserts that the method was called exaclty once
func (ma *methodAssertion) CalledOnce() *finishedMethodAssertion {
	failureCond := !ma.m.CalledOnce()
	if ma.verify(failureCond) {
		msg := mountMethodCallAssertionErrMsg(ma, 1)
		ma.t.Errorf(msg)
	}

	return &finishedMethodAssertion{ma}
}

// Called asserts that the method was called 'n' times
func (ma *methodAssertion) CalledTimes(n int) *finishedMethodAssertion {
	failureCond := !ma.m.CalledTimes(n)
	if ma.verify(failureCond) {
		msg := mountMethodCallAssertionErrMsg(ma, n)
		ma.t.Errorf(msg)
	}

	return &finishedMethodAssertion{ma}
}

type finishedMethodAssertion struct {
	ma *methodAssertion
}

// And is used to chain method assertions
func (fma *finishedMethodAssertion) And() *methodAssertion {
	return &methodAssertion{
		t: fma.ma.t,
		m: fma.ma.m,
	}
}
