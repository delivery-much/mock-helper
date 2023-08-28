package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetResponse(t *testing.T) {
	t.Run("Should set the method response for the mock correctly", func(t *testing.T) {
		m := NewMock()

		method := m.Method("MyFunc")
		res := methodResponse{"res1", "res2"}

		method.SetResponse(res[0], res[1])

		assert.Equal(t, res, m.responses[method.name])
	})
	t.Run("Should not break if method mock is nil", func(t *testing.T) {
		m := method{}

		assert.Nil(t, m.mock)
		m.SetResponse(42)
	})
}

func TestGetResponse(t *testing.T) {
	t.Run("Should return an empty response if the method has no response specified", func(t *testing.T) {
		m := NewMock()
		method := m.Method("MyFunc")

		actual := method.GetResponse()
		assert.Empty(t, actual)
	})
	t.Run("Should get the method response correctly", func(t *testing.T) {
		m := NewMock()
		method := m.Method("MyFunc")

		res := methodResponse{"res1", "res2"}
		method.SetResponse(res[0], res[1])

		actual := method.GetResponse()
		assert.Equal(t, res, actual)
	})
	t.Run("Should not break if method mock value is nil", func(t *testing.T) {
		m := method{}

		assert.Nil(t, m.mock)
		m.GetResponse()
	})
}

func TestGetMethodCalls(t *testing.T) {
	t.Run("Should return empty and dont break if the method mock is nil", func(t *testing.T) {
		m := method{
			name: "MyMethod",
		}

		res := m.GetCalls()

		assert.Empty(t, res)
	})
	t.Run("Should return the calls correctly", func(t *testing.T) {
		mock := NewMock()

		mock.RegisterMethodCall("MyFunc1", 10)
		mock.RegisterMethodCall("MyFunc2", "param", 50)
		mock.RegisterMethodCall("MyFunc2", true)

		func1Calls := mock.Method("MyFunc1").GetCalls()
		func2Calls := mock.Method("MyFunc2").GetCalls()

		assert.Equal(t, 1, len(func1Calls))
		assert.Equal(t, 2, len(func2Calls))
		assert.Equal(t, 0, len(mock.Method("SomeOtherFunc").GetCalls()))

		expectedFunc1Calls := []MockCall{
			{
				MethodName: "MyFunc1",
				Args: []any{
					10,
				},
			},
		}
		expectedFunc2Calls := []MockCall{
			{
				MethodName: "MyFunc2",
				Args: []any{
					"param",
					50,
				},
			},
			{
				MethodName: "MyFunc2",
				Args: []any{
					true,
				},
			},
		}

		assert.Equal(t, expectedFunc1Calls, func1Calls)
		assert.Equal(t, expectedFunc2Calls, func2Calls)
	})
}

func TestMethodCalled(t *testing.T) {
	t.Run("Should return if the method was called correctly", func(t *testing.T) {
		mock := NewMock()

		mock.RegisterMethodCall("MyFunc1")
		mock.RegisterMethodCall("MyFunc2")
		mock.RegisterMethodCall("MyFunc2")

		assert.True(t, mock.Method("MyFunc1").Called())
		assert.True(t, mock.Method("MyFunc2").Called())
		assert.False(t, mock.Method("SomeOtherFunc").Called())
	})
}

func TestMethodCalledOnce(t *testing.T) {
	t.Run("Should return if the method was called exactly once", func(t *testing.T) {
		mock := NewMock()

		mock.RegisterMethodCall("MyFunc1")
		mock.RegisterMethodCall("MyFunc2")
		mock.RegisterMethodCall("MyFunc2")

		assert.False(t, mock.Method("MyFunc2").CalledOnce())
		assert.True(t, mock.Method("MyFunc1").CalledOnce())
		assert.False(t, mock.Method("SomeOtherFunc").CalledOnce())
	})
}

func TestMethodCalledNTimes(t *testing.T) {
	t.Run("Should return if the method was called exactly once", func(t *testing.T) {
		mock := NewMock()

		mock.RegisterMethodCall("MyFunc1")
		mock.RegisterMethodCall("MyFunc2")
		mock.RegisterMethodCall("MyFunc2")
		mock.RegisterMethodCall("MyFunc3")
		mock.RegisterMethodCall("MyFunc3")
		mock.RegisterMethodCall("MyFunc3")

		assert.False(t, mock.Method("MyFunc2").CalledTimes(3))
		assert.False(t, mock.Method("MyFunc1").CalledTimes(3))
		assert.True(t, mock.Method("MyFunc3").CalledTimes(3))
		assert.False(t, mock.Method("SomeOtherFunc").CalledTimes(3))

		assert.True(t, mock.Method("SomeOtherFunc").CalledTimes(0))
	})
}

func TestMethodCalledWith(t *testing.T) {
	t.Run("Should return false if the method was not called with the arguments", func(t *testing.T) {
		mock := NewMock()
		method1 := mock.Method("MyFunc1")
		method2 := mock.Method("MyFunc2")
		arg1 := "MyArg"
		arg2 := 10

		assert.False(t, method1.CalledWith())
		assert.False(t, method2.CalledWith())

		mock.RegisterMethodCall("MyFunc1", arg1)
		mock.RegisterMethodCall("MyFunc2", arg2)

		assert.False(t, method2.CalledWith(arg1))
		assert.False(t, method1.CalledWith(arg2))

		mock.Reset()
		mock.RegisterMethodCall("MyFunc1", arg1)
		mock.RegisterMethodCall("MyFunc1", arg2)

		assert.False(t, method2.CalledWith(arg1, arg2))
	})
	t.Run("Should return true if the method was called with the arguments", func(t *testing.T) {
		mock := NewMock()
		method1 := mock.Method("MyFunc1")
		method2 := mock.Method("MyFunc2")
		arg1 := "MyArg"
		arg2 := 10

		mock.RegisterMethodCall("MyFunc1", arg1)
		mock.RegisterMethodCall("MyFunc2", arg2)

		assert.True(t, method1.CalledWith(arg1))
		assert.True(t, method2.CalledWith(arg2))

		mock.Reset()

		mock.RegisterMethodCall("MyFunc1", arg1, arg2)
		assert.True(t, method1.CalledWith(arg2, arg1))

		mock.Reset()

		mock.RegisterMethodCall("MyFunc1", arg1, arg2, "anotherArg")
		assert.True(t, method1.CalledWith(arg2, arg1))
	})
}

func TestMethodCalledWithExactly(t *testing.T) {
	t.Run("Should return false if the method was not called with the arguments", func(t *testing.T) {
		mock := NewMock()
		method1 := mock.Method("MyFunc1")
		method2 := mock.Method("MyFunc2")
		arg1 := "MyArg"
		arg2 := 10

		assert.False(t, method1.CalledWithExactly())
		assert.False(t, method2.CalledWithExactly())

		mock.RegisterMethodCall("MyFunc1", arg1)
		mock.RegisterMethodCall("MyFunc2", arg2)

		assert.False(t, method2.CalledWithExactly(arg1))
		assert.False(t, method1.CalledWithExactly(arg2))

		mock.Reset()

		mock.RegisterMethodCall("MyFunc1", arg1, arg2)
		assert.False(t, method1.CalledWithExactly(arg2, arg1))

		mock.Reset()

		mock.RegisterMethodCall("MyFunc1", arg1, arg2, "anotherArg")
		assert.False(t, method1.CalledWithExactly(arg1, arg2))
	})
	t.Run("Should return true if the method was called with the exact same arguments, on the same order", func(t *testing.T) {
		mock := NewMock()
		method1 := mock.Method("MyFunc1")
		method2 := mock.Method("MyFunc2")
		arg1 := "MyArg"
		arg2 := 10

		mock.RegisterMethodCall("MyFunc1", arg1)
		mock.RegisterMethodCall("MyFunc2", arg2)

		assert.True(t, method1.CalledWithExactly(arg1))
		assert.True(t, method2.CalledWithExactly(arg2))

		mock.Reset()

		mock.RegisterMethodCall("MyFunc1", arg1, arg2)
		mock.RegisterMethodCall("MyFunc2", arg2, "someOtherArg")

		assert.True(t, method1.CalledWithExactly(arg1, arg2))
		assert.True(t, method2.CalledWithExactly(arg2, "someOtherArg"))
	})
}
