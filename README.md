<p align="center"><img src="assets/gopher.png" width="500"></p>

<h1 align="center">
  mock-helper
</h2>

Mock helper is a tool to help create readable and easy tests, creating a simple and idiomatic way to mock interfaces and structures.

Mock helper is based on the [Mock package from stretchr](https://pkg.go.dev/github.com/stretchr/testify/mock) and the [Jest test framework from NodeJS](https://jestjs.io/).

It aims to combine stretchr's strategy to create mocks, with the readability from the tests created with the Jest framework.

- [Setup](#setup)
- [How to Mock](#how-to-mock)
- [Features](#features)
  - [Mock](#mock)
    - [func NewMock](#func-newmock)
    - [func SetMethodResponse](#func-setmethodresponse)
    - [func GetMethodResponse](#func-getmethodresponse)
    - [func RegisterMethodCall](#func-registermethodcall)
    - [func GetCalls](#func-getcalls)
    - [func Called](#func-called)
    - [func CalledOnce](#func-calledonce)
    - [func CalledTimes](#func-calledtimes)
    - [func CalledWith](#func-calledwith)
    - [func CalledWithExactly](#func-calledwithexactly)
    - [func Reset](#func-reset)
    - [func Method](#func-method)
  - [MockCall](#mockcall)
    - [func HasArgument](#func-hasargument)
  - [Method](#method)
    - [func SetResponse](#func-setresponse)
    - [func GetResponse](#func-getresponse)
    - [func GetCalls](#func-getcalls-1)
    - [func Called](#func-called-1)
    - [func CalledOnce](#func-calledonce-1)
    - [func CalledTimes](#func-calledtimes-1)
    - [func CalledWith](#func-calledwith-1)
    - [func CalledWithExactly](#func-calledwithexactly-1)
  - [MethodResponse](#methodresponse)
    - [func Get](#func-get-1)
    - [func GetString](#func-getstring)
    - [func GetInt](#func-getint)
    - [func GetError](#func-geterror)

## Setup
To download mock-helper and add it to your project, just run:

```shell
$ go get github.com/delivery-much/mock-helper
```

## How to mock
As said before, mock helper gives you tools to help you create and test mocks in a easy and readable way.
The implementation of the mock itself is up to you.

For instance, lets say you have a database interface:
```go
type MyDBInterface struct {
  GetUserCount(userID string) (int, error)
}
```
This interface has one method, that receives a string, and returns a user count and an error.

And you want to test this method:
```go
func GetCount(db MyDBInterface, userID string) (int, error) {
  c, err := db.GetUserCount(userID)
  if err != nil {
    return 0, err
  }

  return c, err
}
```
That uses your interface.

You could easily do that using mock helper!

Just create a mock implementation for your interface, like this:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type dbMock struct {
  mock.Mock
}

// create a function to instantiate your mock
func NewDBMock() dbMock {
  return dbMock{
    mock.NewMock(),
  }
}

func (m *InterfaceMock) GetUserCount(userID string) (c int, err error) {
  // Register the method call with the correct parameters
  m.RegisterMethodCall("GetUserCount", userID)

  // Gets the method response that was specified on the tests, if any
  res := m.GetMethodResponse("GetUserCount")

  if len(res) == 0 {
    // if no response was specified, returns the default values
    return
  }

  // if a response was specified, gets the values from it
  return res.GetInt(0), res.GetError(1)
}
```

And then test your function, like this:
```go
dbMock := NewDBMock()
// set the mock response
dbMock.Method("GetUserCount").SetResponse(5, nil)

userID := "mockUserID"

// call your function
c, err := GetCount(dbMock, userID)

// make your assertions
assert.Equal(t, 5, c)
assert.Nil(t, err)
assert.True(t, dbMock.CalledOnce())
assert.True(t, dbMock.CalledWith(userID))

// reset the mock
dbMock.Reset()

// set the mock response (this time with an error)
mockErr := errors.New("Mock error")
dbMock.Method("GetUserCount").SetResponse(5, mockErr)

// call your function
c, err = GetCount(dbMock, userID)

// make your assertions
assert.Equal(t, 0, c)
assert.Equal(t, mockErr, err)
assert.True(t, dbMock.CalledOnce())
assert.True(t, dbMock.CalledWith(userID))
```

This way you can easily mock your interfaces and assert that they where called the correct way, and their return value where used correctly.


## Features

### Mock

In its root, the Mock library has a series of tools to help you mock your interfaces and create your assertions.

#### func NewMock

The NewMock function returns a new and empty Mock struct.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func NewMyMock() MyMock {
  return MyMock{
    mock.NewMock(),
  }
}
```

#### func SetMethodResponse

The SetMethodResponse function sets the response that the mock will return when calling the method with the specified name.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }

  m.SetMethodResponse("MyMethod", 10, "something")

  m.GetMethodResponse("MyMethod") // Should return [10, "something"]
}
```
**Note:** Calling `m.SetMethodResponse("MyMethod", ...)` is directly equivalent of calling `m.Method("MyMethod").SetResponse(...)`.

#### func GetMethodResponse

The GetMethodResponse function returns the response that was specified for the method with the given name.

This method will return a [MethodResponse](#methodresponse). This response will be empty if no response was specified for the method, or will contain the response values for that method.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }

  m.SetMethodResponse("MyMethod1", 10, "something")

  m.GetMethodResponse("MyMethod1") // Should return [10, "something"]
  m.GetMethodResponse("MyMethod2") // Should return []
}
```
**Note:** Calling `m.GetMethodResponse("MyMethod")` is directly equivalent of calling `m.Method("MyMethod").GetResponse()`.

#### func RegisterMethodCall

The RegisterMethodCall function registers a mock method call.

It receives the method name and the params that where passed on that method call, and registers that information inside the mock to be used later.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func (m *MyMock) MyMockFunc(param1 string, param2 int) {
  m.RegisterMethodCall("MyMockFunc", param1, param2)

  // rest of the implementation goes here
}
```
**Note:** It's imperative that the RegisterMethodCall its used correctly, with the correct function name and the correct parameters, so that later, in the tests, the assertions can be made safely and avoid false negatives or positives.

#### func GetCalls

The GetCalls function returns the calls that where made on that mock.
It returns a slice of [MockCall](#mockcall), this slice will be empty if the mock had no calls, or it will contain the calls that where made on that mock.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m1 := MyMock{
    mock.NewMock(),
  }

  m2 := MyMock{
    mock.NewMock(),
  }

  m1.RegisterMethodCall("MyMockFunc", 10, "42")

  m1.GetCalls() // Should return [{ "MyMockFunc", [10, "42"] }]
  m2.GetCalls() // Should return []
}
```

#### func Called

The Called function checks whether the mock has been called at least once. 
It returns a boolean value indicating whether any method calls have been registered on the mock instance.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  m.Called() // Returns false, as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.Called() // Returns true, since at least one method call has been registered
}
```

#### func CalledOnce

The CalledOnce function checks whether the mock has been called exactly once. It returns a boolean value indicating whether there is exactly one method call registered on the mock instance.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  m.CalledOnce() // Returns false, as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledOnce() // Returns true, since there is exactly one method call registered

  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledOnce() // Returns false, since there whas more than one method call registered
}
```

#### func CalledTimes
The CalledTimes function checks whether the mock has been called a specific number of times. 
It takes an integer argument `n` and returns a boolean value indicating whether the method calls count matches the provided number `n`.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  m.CalledTimes(0) // Returns true, as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledTimes(1) // Returns true, since there is exactly one method call registered
  m.CalledTimes(2) // Returns false, since there's only one method call
}
```

#### func CalledWith
The CalledWith function checks whether the mock has been called at least once with specific arguments.
It takes variadic arguments representing the expected arguments and returns a boolean value indicating whether there is a method call with (at least) those arguments.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledWith("param1", 42) // Returns true, since there's a matching method call
  m.CalledWith(42, "param1") // Returns true, since there's a matching method call
  m.CalledWith("param1", 43) // Returns false, since there's no matching call
  m.CalledWith("param1") // Returns true, since one of the method calls has the param "param1"
}
```

#### func CalledWithExactly
The CalledWithExactly function checks whether the mock has been called at least once with exactly matching arguments, in the same order. 
It takes variadic arguments representing the expected arguments and returns a boolean value indicating whether there is a method call with (exactly) those arguments.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledWith("param1", 42) // Returns true, since there's a matching method call
  m.CalledWith(42, "param1") // Returns false, since the method order is incorrect
  m.CalledWith("param1", 43) // Returns false, since there's no matching call
  m.CalledWith("param1") // Returns false, since there's no matching call
}
```

#### func Reset
The Reset function clears all registered method calls and responses from the mock, effectively resetting it to an empty state.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.Called() // Returns true, since there's a registered method call
  
  m.Reset()
  m.Called() // Returns false, as all method calls have been cleared
}
```

#### func Method
The Method function returns a [Method](#method) instance that can be used to filter the mock's use information for a specific method.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.SetResponse("mock response")
  
  // Perform tests using 'myMethod'
}
```

### MockCall
The `MockCall` structure represents a mock method call, including the method name and the arguments passed during the call.

It has the properties:
- `MethodName` (string): The name of the method that was called.
- `Args` ([]any): A slice containing the arguments passed to the method during the call.


#### func HasArgument

The HasArgument method checks whether a specific argument exists within the list of arguments for a mock call.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

func main() {
  call := mock.MockCall{
    MethodName: "MyMethod",
    Args:       []interface{}{"param1", 42},
  }

  result1 := call.HasArgument("param1") // Returns true
  result2 := call.HasArgument(42)       // Returns true
  result3 := call.HasArgument("foo")    // Returns false
}
```

### Method

The Method type represents mock usage information that is filtered for a specific method.

#### func SetResponse

The SetResponse function allows you to specify the response values that the mock method should return when it's called.
The provided response values must match the method's response signature in type and order.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.SetResponse("mock response")
  
  myMethod.GetResponse() // Should return "mock response"
}
```

#### func GetResponse
The GetResponse function gets the specified response for the method.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.SetResponse("mock response")
  
  response := myMethod.GetResponse() // Should return "mock response"
}
```

#### func GetCalls
The GetCalls function returns the mock method calls that were made specifically for the filtered method.
It returns a slice of [MockCall](#mockcall)

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.GetCalls() // Returns [], as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  calls := myMethod.GetCalls() // Returns a slice containing calls for "MyMethod"
}
```

#### func Called
The Called function checks if the mock method was called at least once.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.Called() // Returns false, as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  myMethod.Called() // Returns true, since "MyMethod" was called
}
```

#### func CalledOnce
The CalledOnce function checks if the mock method was called exactly once.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.CalledOnce() // Returns false, as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  myMethod.CalledOnce() // Returns true, since "MyMethod" was called exactly once

  m.RegisterMethodCall("MyMethod", "param1", 42)
  myMethod.CalledOnce() // Returns false, since "MyMethod" was called more than once
}
```

#### func CalledTimes
The CalledTimes function checks whether the mock method has been called a specific number of times. 
It takes an integer argument `n` and returns a boolean value indicating whether the method calls count matches the provided number `n`.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  m := MyMock{
    mock.NewMock(),
  }
  
  myMethod := m.Method("MyMethod")
  myMethod.CalledTimes(0) // Returns true, as no method calls have been registered
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  
  myMethod.CalledTimes(1) // Returns true, since "MyMethod" was called once
  myMethod.CalledTimes(2) // Returns false, since "MyMethod" was called only once
}
```

#### func CalledWith
The CalledWith function checks whether the mock method has been called at least once with specific arguments.
It takes variadic arguments representing the expected arguments and returns a boolean value indicating whether there is a method call with (at least) those arguments.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  mock := MyMock{
    mock.NewMock(),
  }

  m := mock.Method("MyMethod")
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledWith("param1", 42) // Returns true, since there's a matching method call
  m.CalledWith(42, "param1") // Returns true, since there's a matching method call
  m.CalledWith("param1", 43) // Returns false, since there's no matching call
  m.CalledWith("param1") // Returns true, since one of the method calls has the param "param1"
}
```

#### func CalledWithExactly
The CalledWithExactly function checks whether the mock method has been called at least once with exactly matching arguments, in the same order. 
It takes variadic arguments representing the expected arguments and returns a boolean value indicating whether there is a method call with (exactly) those arguments.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

type MyMock struct {
  mock.Mock
}

func MyTest() {
  mock := MyMock{
    mock.NewMock(),
  }

  m := mock.Method("MyMethod")
  
  m.RegisterMethodCall("MyMethod", "param1", 42)
  m.CalledWith("param1", 42) // Returns true, since there's a matching method call
  m.CalledWith(42, "param1") // Returns false, since the method order is incorrect
  m.CalledWith("param1", 43) // Returns false, since there's no matching call
  m.CalledWith("param1") // Returns false, since there's no matching call
}
```

### MethodResponse

The `MethodResponse` type represents a response that a mock method should return. 
It provides methods for retrieving specific types of response values from the method response.

#### func Get

The Get function returns the response value specified in the method response at the given index.
This function panics if no value is found on the specified index.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

func main() {
  response := mock.MethodResponse{"value1", 42}
  
  value := response.Get(0) // Returns "value1"
  value2 := response.Get(1) // Returns 42
  value2 := response.Get(3) // Panics!!
}
```

#### func GetString
The GetString function returns a string value specified in the method response at the given index. 
It panics if no response value is found on the specified index or if the value type is not a string.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

func main() {
  response := mock.MethodResponse{"hello", "world", 42}
  
  value := response.GetString(0) // Returns "hello"
  value2 := response.GetString(1) // Returns "world"
  value2 := response.GetString(2) // Panics!!! (since 42 is not a string)
  value2 := response.GetString(3) // Panics!!! (since there is no value on the index 3)
}
```

#### func GetInt
The GetInt function returns an integer value specified in the method response at the given index. 
It panics if no response value is found on the specified index or if the value type is not a integer.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
)

func main() {
  response := mock.MethodResponse{42, "test"}
  
  value := response.GetInt(0) // Returns 42
  value2 := response.GetInt(1) // Panics!!! (since "test" is not a integer)
  value2 := response.GetInt(2) // Panics!!! (since there is no value on the index 2)
}
```

#### func GetError
The GetError function returns an error value specified in the method response at the given index.
It panics if no response value is found on the specified index or if the value type is not an error. 
A nil value is considered a valid error response.

Example usage:
```go
import (
  "github.com/delivery-much/mock-helper/mock"
  "errors"
)

func main() {
  response := mock.MethodResponse{errors.New("error 1"), nil, 42}
  
  value := response.GetError(0) // Returns the error instance
  value2 := response.GetError(1) // Returns nil (valid error response)
  value2 := response.GetError(2) // Panics!!! (since 42 is not an error)
  value2 := response.GetError(3) // Panics!!! (since there is no value on the index 3)
}
```
