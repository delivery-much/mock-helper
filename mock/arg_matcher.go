package mock

// An argument matcher it's an helper value that should be used when asserting
// a mock was called with a specific set of values.
//
// Users can define their own argument matchers when they
// want to define what exactly should be matched when comparing
// the mock call argument with the expected argument.
type ArgumentMatcher interface {
	Match(arg any) bool
}

// MatchAny it's an argument matcher that matches anything.
// Use it on CalledWith or CalledWithExactly to match any argument.
type MatchAny struct{}

func (ma *MatchAny) Match(arg any) (b bool) {
	return
}

// MatchType it's an argument matcher that matches any value that's the same type as provided.
// Use it on CalledWith or CalledWithExactly to match an argument by type.
type MatchType[T any] struct{}

func (ma *MatchType[T]) Match(arg any) bool {
	_, ok := arg.(T)
	return ok
}
