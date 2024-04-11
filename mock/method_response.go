package mock

import (
	"fmt"
)

// methodResponse represents a response that a mock method should return
type methodResponse []any

// IsEmpty checks if the method response is empty
func (mr methodResponse) IsEmpty() bool {
	return len(mr) <= 0
}

// Get returns the response value specified in the method response on the 'i' index.
//
// A nil value will be returned if no response value is found on the specified index
func (mr methodResponse) Get(i int) any {
	if len(mr) < i+1 {
		return nil
	}

	return mr[i]
}

// GetBool returns a bool value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetBool(i int) bool {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a bool value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(bool)
	if !ok {
		msg := fmt.Sprintf("Tried to find a bool value on the index %d of the mock method response, but the index value was not an bool", i)
		panic(msg)
	}

	return val
}

// GetString returns a string value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetString(i int) string {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a string value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(string)
	if !ok {
		msg := fmt.Sprintf("Tried to find a string value on the index %d of the mock method response, but the index value was not an string", i)
		panic(msg)
	}

	return val
}

// GetInt returns a int value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetInt(i int) int {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a int value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(int)
	if !ok {
		msg := fmt.Sprintf("Tried to find a int value on the index %d of the mock method response, but the index value was not an int", i)
		panic(msg)
	}

	return val
}

// GetInt8 returns a int8 value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetInt8(i int) int8 {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a int8 value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(int8)
	if !ok {
		msg := fmt.Sprintf("Tried to find a int8 value on the index %d of the mock method response, but the index value was not an int8", i)
		panic(msg)
	}

	return val
}

// GetInt16 returns a int16 value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetInt16(i int) int16 {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a int16 value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(int16)
	if !ok {
		msg := fmt.Sprintf("Tried to find a int16 value on the index %d of the mock method response, but the index value was not an int16", i)
		panic(msg)
	}

	return val
}

// GetInt32 returns a int32 value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetInt32(i int) int32 {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a int32 value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(int32)
	if !ok {
		msg := fmt.Sprintf("Tried to find a int32 value on the index %d of the mock method response, but the index value was not an int32", i)
		panic(msg)
	}

	return val
}

// GetInt64 returns a int64 value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetInt64(i int) int64 {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a int64 value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(int64)
	if !ok {
		msg := fmt.Sprintf("Tried to find a int64 value on the index %d of the mock method response, but the index value was not an int64", i)
		panic(msg)
	}

	return val
}

// GetFloat32 returns a float32 value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetFloat32(i int) float32 {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a float32 value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(float32)
	if !ok {
		msg := fmt.Sprintf("Tried to find a float32 value on the index %d of the mock method response, but the index value was not an float32", i)
		panic(msg)
	}

	return val
}

// GetFloat64 returns a float64 value that should be specified in the method response on the 'i' index.
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetFloat64(i int) float64 {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a float64 value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	val, ok := mr[i].(float64)
	if !ok {
		msg := fmt.Sprintf("Tried to find a float64 value on the index %d of the mock method response, but the index value was not an float64", i)
		panic(msg)
	}

	return val
}

// GetError returns a error value that should be specified in the method response on the 'i' index.
//
// (Nil is also considered a valid error)
//
// This method panics if no response value is found on the specified index, or the value type is wrong
func (mr methodResponse) GetError(i int) error {
	if len(mr) < i+1 {
		msg := fmt.Sprintf("Tried to find a error value on the index %d of the mock method response, but the index had no value", i)
		panic(msg)
	}

	if mr[i] == nil {
		return nil
	}

	val, ok := mr[i].(error)
	if !ok {
		msg := fmt.Sprintf("Tried to find a error value on the index %d of the mock method response, but the index value was not an error", i)
		panic(msg)
	}

	return val
}
