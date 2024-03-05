package mock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mr := methodResponse{"value1"}
	t.Run("Should return nil if the index has no value", func(t *testing.T) {
		assert.Nil(t, mr.Get(1))
	})
	t.Run("Should return the value if the index has a value", func(t *testing.T) {
		assert.Equal(t, "value1", mr.Get(0))
	})
}

func TestGetBool(t *testing.T) {
	mr := methodResponse{true, "value2"}
	t.Run("Should panic with correct message if the index value is not an bool", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a bool value on the index 1 of the mock method response, but the index value was not an bool",
			func() {
				_ = mr.GetBool(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a bool value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetBool(2)
			},
		)
	})
	t.Run("Should return the value if the index value is a valid boolean", func(t *testing.T) {
		assert.True(t, mr.GetBool(0))
	})
}

func TestGetString(t *testing.T) {
	mr := methodResponse{"myValue", 10}

	t.Run("Should panic with correct message if the index value is not an string", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a string value on the index 1 of the mock method response, but the index value was not an string",
			func() {
				_ = mr.GetString(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a string value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetString(2)
			},
		)
	})
	t.Run("Should return the correct value if the method is a valid string", func(t *testing.T) {
		assert.Equal(t, "myValue", mr.GetString(0))
	})
}

func TestGetInt(t *testing.T) {
	mr := methodResponse{42, "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int value on the index 1 of the mock method response, but the index value was not an int",
			func() {
				_ = mr.GetInt(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetInt(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, 42, mr.GetInt(0))
	})
}

func TestGetInt8(t *testing.T) {
	mr := methodResponse{int8(42), "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int8 value on the index 1 of the mock method response, but the index value was not an int8",
			func() {
				_ = mr.GetInt8(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int8 value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetInt8(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, int8(42), mr.GetInt8(0))
	})
}

func TestGetInt16(t *testing.T) {
	mr := methodResponse{int16(42), "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int16 value on the index 1 of the mock method response, but the index value was not an int16",
			func() {
				_ = mr.GetInt16(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int16 value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetInt16(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, int16(42), mr.GetInt16(0))
	})
}

func TestGetInt32(t *testing.T) {
	mr := methodResponse{int32(42), "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int32 value on the index 1 of the mock method response, but the index value was not an int32",
			func() {
				_ = mr.GetInt32(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int32 value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetInt32(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, int32(42), mr.GetInt32(0))
	})
}

func TestGetInt64(t *testing.T) {
	mr := methodResponse{int64(42), "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int64 value on the index 1 of the mock method response, but the index value was not an int64",
			func() {
				_ = mr.GetInt64(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a int64 value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetInt64(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, int64(42), mr.GetInt64(0))
	})
}

func TestGetFloat32(t *testing.T) {
	mr := methodResponse{float32(42), "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a float32 value on the index 1 of the mock method response, but the index value was not an float32",
			func() {
				_ = mr.GetFloat32(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a float32 value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetFloat32(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, float32(42), mr.GetFloat32(0))
	})
}

func TestGetFloat64(t *testing.T) {
	mr := methodResponse{float64(42), "value2"}

	t.Run("Should panic with correct message if the index value is not an int", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a float64 value on the index 1 of the mock method response, but the index value was not an float64",
			func() {
				_ = mr.GetFloat64(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a float64 value on the index 2 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetFloat64(2)
			},
		)
	})
	t.Run("Should return the correct value if the index value is an int", func(t *testing.T) {
		assert.Equal(t, float64(42), mr.GetFloat64(0))
	})
}

func TestMethodResponseGetError(t *testing.T) {
	err := fmt.Errorf("error")
	mr := methodResponse{err, "value2", nil}

	t.Run("Should panic with correct message if the index value is not an error", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a error value on the index 1 of the mock method response, but the index value was not an error",
			func() {
				_ = mr.GetError(1)
			},
		)
	})
	t.Run("Should panic with correct message if the index has no value", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"Tried to find a error value on the index 3 of the mock method response, but the index had no value",
			func() {
				_ = mr.GetError(3)
			},
		)
	})
	t.Run("Should return the correct value if the index value is a valid error", func(t *testing.T) {
		assert.Equal(t, err, mr.GetError(0))
		assert.Nil(t, mr.GetError(2))
	})
}
