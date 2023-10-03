package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMethodResponse(t *testing.T) {
	t.Run("Should set the method response for the mock correctly", func(t *testing.T) {
		m := NewMock()

		fnName := "MyFunc"
		res := methodResponse{"res1", "res2"}
		m.SetMethodResponse(fnName, res...)

		actual := m.responses[fnName]
		assert.Equal(t, res, actual)
	})
}

func TestGetMethodResponse(t *testing.T) {
	t.Run("Should return an empty response if the method has no response specified", func(t *testing.T) {
		m := NewMock()

		actual := m.GetMethodResponse("MyFunc")
		assert.Empty(t, actual)
	})
	t.Run("Should get the method response correctly", func(t *testing.T) {
		m := NewMock()

		fnName := "MyFunc"
		res := methodResponse{"res1", "res2"}
		m.SetMethodResponse(fnName, res...)

		actual := m.GetMethodResponse("MyFunc")
		assert.Equal(t, res, actual)
	})
}

func TestRegisterMethodCall(t *testing.T) {
	t.Run("Should register a method call with arguments correctly", func(t *testing.T) {
		m := NewMock()

		fnName := "MyFunc"
		arg1 := "MYSUPERARG"
		arg2 := 10
		m.RegisterMethodCall(fnName, arg1, arg2)

		assert.NotEmpty(t, m.calls)
		assert.Equal(t, 1, len(m.calls))
		assert.Equal(t, fnName, m.calls[0].MethodName)
		assert.Equal(t, arg1, m.calls[0].Args[0])
		assert.Equal(t, arg2, m.calls[0].Args[1])
	})
	t.Run("Should register a method call without arguments correctly", func(t *testing.T) {
		m := NewMock()

		fnName := "MyFunc"
		m.RegisterMethodCall(fnName)

		assert.NotEmpty(t, m.calls)
		assert.Equal(t, 1, len(m.calls))
		assert.Equal(t, fnName, m.calls[0].MethodName)
		assert.Empty(t, m.calls[0].Args)
	})
}

func TestGetCalls(t *testing.T) {
	t.Run("Should get the mock calls correctly", func(t *testing.T) {
		m := NewMock()

		res := m.GetCalls()
		assert.Empty(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
			{
				MethodName: "MyFunc2",
			},
		}

		res = m.GetCalls()
		assert.NotEmpty(t, res)
		assert.Equal(t, 2, len(res))
		assert.Equal(t, m.calls[0], res[0])
		assert.Equal(t, m.calls[1], res[1])
	})
}

func TestCalled(t *testing.T) {
	t.Run("Should return if the method was called", func(t *testing.T) {
		m := NewMock()

		res := m.Called()
		assert.False(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
		}

		res = m.Called()
		assert.True(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
			{
				MethodName: "MyFunc2",
			},
		}

		res = m.Called()
		assert.True(t, res)
	})
}

func TestCalledOnce(t *testing.T) {
	t.Run("Should return if the method was exactly once", func(t *testing.T) {
		m := NewMock()

		res := m.CalledOnce()
		assert.False(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
		}

		res = m.CalledOnce()
		assert.True(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
			{
				MethodName: "MyFunc2",
			},
		}

		res = m.CalledOnce()
		assert.False(t, res)
	})
}

func TestCalledTimes(t *testing.T) {
	t.Run("Should return if the method was exactly n times", func(t *testing.T) {
		m := NewMock()

		res := m.CalledTimes(2)
		assert.False(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
		}

		res = m.CalledTimes(2)
		assert.False(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
			{
				MethodName: "MyFunc2",
			},
		}

		res = m.CalledTimes(2)
		assert.True(t, res)

		m.calls = []MockCall{
			{
				MethodName: "MyFunc1",
			},
			{
				MethodName: "MyFunc2",
			},
			{
				MethodName: "MyFunc3",
			},
		}

		res = m.CalledTimes(2)
		assert.False(t, res)
	})
}

func TestCalledWith(t *testing.T) {
	t.Run("Should return false if the mock was not called with the arguments", func(t *testing.T) {
		m := NewMock()
		arg1 := "MyArg"
		arg2 := 10

		res := m.CalledWith()
		assert.False(t, res)

		res = m.CalledWith(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg1}},
		}
		res = m.CalledWith(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg2}},
		}
		res = m.CalledWith(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg1}},
			{Args: []any{arg2}},
		}
		res = m.CalledWith(arg1, arg2)
		assert.False(t, res)
	})
	t.Run("Should return true if the mock was called with the arguments", func(t *testing.T) {
		m := NewMock()
		arg1 := "MyArg"
		arg2 := 10

		m.calls = []MockCall{
			{Args: []any{arg1, arg2}},
		}
		res := m.CalledWith(arg1, arg2)
		assert.True(t, res)

		m.calls = []MockCall{
			{Args: []any{arg2, arg1}},
		}
		res = m.CalledWith(arg1, arg2)
		assert.True(t, res)

		m.calls = []MockCall{
			{Args: []any{arg2, "some other argument", arg1, 42}},
		}
		res = m.CalledWith(arg1, arg2)
		assert.True(t, res)

		m.calls = []MockCall{
			{Args: []any{42}},
			{Args: []any{arg1, arg2}},
			{Args: []any{"some other argument"}},
		}
		res = m.CalledWith(arg1, arg2)
		assert.True(t, res)
	})
	t.Run("Should be able to compare slices", func(t *testing.T) {
		m := NewMock()
		sliceArg := []string{"1", "2", "3"}

		m.calls = []MockCall{
			{Args: []any{sliceArg}},
		}

		res := m.CalledWith(sliceArg)
		assert.True(t, res)

		res = m.CalledWith([]string{"3", "2", "1"})
		assert.False(t, res)

		res = m.CalledWith(20)
		assert.False(t, res)
	})
	t.Run("Should be able to compare maps", func(t *testing.T) {
		m := NewMock()
		mapArg := map[string]int{"1": 3, "2": 4, "3": 5}

		m.calls = []MockCall{
			{Args: []any{mapArg}},
		}

		res := m.CalledWith(mapArg)
		assert.True(t, res)

		res = m.CalledWith(map[string]int{"3": 5, "2": 4, "1": 3})
		assert.True(t, res)

		res = m.CalledWith(map[string]int{"3": 5, "2": 4})
		assert.False(t, res)

		res = m.CalledWith(20)
		assert.False(t, res)
	})
}

func TestCalledWithExactly(t *testing.T) {
	t.Run("Should return false if the mock was not called with the arguments", func(t *testing.T) {
		m := NewMock()
		arg1 := "MyArg"
		arg2 := 10

		res := m.CalledWithExactly()
		assert.False(t, res)

		res = m.CalledWithExactly(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg1}},
		}
		res = m.CalledWithExactly(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg2}},
		}
		res = m.CalledWithExactly(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg1}},
			{Args: []any{arg2}},
		}
		res = m.CalledWithExactly(arg1, arg2)
		assert.False(t, res)
	})
	t.Run("Should return true if the mock was called with the exact same arguments, on the same order", func(t *testing.T) {
		m := NewMock()
		arg1 := "MyArg"
		arg2 := 10

		m.calls = []MockCall{
			{Args: []any{arg1, arg2}},
		}
		res := m.CalledWithExactly(arg1, arg2)
		assert.True(t, res)

		m.calls = []MockCall{
			{Args: []any{arg2, arg1}},
		}
		res = m.CalledWithExactly(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{arg2, "some other argument", arg1, 42}},
		}
		res = m.CalledWithExactly(arg1, arg2)
		assert.False(t, res)

		m.calls = []MockCall{
			{Args: []any{42}},
			{Args: []any{arg1, arg2}},
			{Args: []any{"some other argument"}},
		}
		res = m.CalledWithExactly(arg1, arg2)
		assert.True(t, res)
	})
	t.Run("Should be able to compare slices", func(t *testing.T) {
		m := NewMock()
		sliceArg := []string{"1", "2", "3"}

		m.calls = []MockCall{
			{Args: []any{sliceArg}},
		}

		res := m.CalledWithExactly(sliceArg)
		assert.True(t, res)

		res = m.CalledWithExactly([]string{"3", "2", "1"})
		assert.False(t, res)

		res = m.CalledWithExactly(20)
		assert.False(t, res)
	})
	t.Run("Should be able to compare maps", func(t *testing.T) {
		m := NewMock()
		mapArg := map[string]int{"1": 3, "2": 4, "3": 5}

		m.calls = []MockCall{
			{Args: []any{mapArg}},
		}

		res := m.CalledWithExactly(mapArg)
		assert.True(t, res)

		res = m.CalledWithExactly(map[string]int{"3": 5, "2": 4, "1": 3})
		assert.True(t, res)

		res = m.CalledWithExactly(map[string]int{"3": 5, "2": 4})
		assert.False(t, res)

		res = m.CalledWithExactly(20)
		assert.False(t, res)
	})
}

func TestReset(t *testing.T) {
	t.Run("Should reset the mock", func(t *testing.T) {
		m := NewMock()

		m.RegisterMethodCall("MyMethod", 42)
		m.SetMethodResponse("MyMethod", "response")

		m.Reset()

		assert.NotNil(t, m.responses)
		assert.Empty(t, m.responses)
		assert.Empty(t, m.calls)
	})
}

func TestMethod(t *testing.T) {
	t.Run("Should return a valid method struct with the specified method name", func(t *testing.T) {
		m := NewMock()

		fnName := "MyFunc"
		method := m.Method(fnName)

		assert.Equal(t, fnName, method.name)
		assert.Equal(t, m, *method.mock)
	})
}
