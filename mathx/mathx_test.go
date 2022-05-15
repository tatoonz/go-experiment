package mathx_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/goexperiment/mathx"
)

func TestFibonacci(t *testing.T) {
	actual, err := mathx.Fibonacci(42)

	assert.NoError(t, err)
	assert.Equal(t, uint64(267914296), actual)
}

func ExampleFibonacci() {
	fmt.Println(mathx.Fibonacci(3))
}

func TestSum(t *testing.T) {
	assert.Equal(t, int(3), mathx.Sum(1, 2))
	assert.Equal(t, int8(3), mathx.Sum[int8](1, 2))
	assert.Equal(t, int16(3), mathx.Sum[int16](1, 2))
	assert.Equal(t, int32(3), mathx.Sum[int32](1, 2))
	assert.Equal(t, int64(3), mathx.Sum[int64](1, 2))

	assert.Equal(t, uint(3), mathx.Sum[uint](1, 2))
	assert.Equal(t, uint8(3), mathx.Sum[uint8](1, 2))
	assert.Equal(t, uint16(3), mathx.Sum[uint16](1, 2))
	assert.Equal(t, uint32(3), mathx.Sum[uint32](1, 2))
	assert.Equal(t, uint64(3), mathx.Sum[uint64](1, 2))
	assert.Equal(t, uint64(3), mathx.Sum[uint64](1, 2))

	assert.Equal(t, float32(3.0), mathx.Sum[float32](1.0, 2.0))
	assert.Equal(t, float64(3.0), mathx.Sum(1.0, 2.0))

	assert.Equal(t, complex64(3.0), mathx.Sum[complex64](1.0, 2.0))
	assert.Equal(t, complex128(3.0), mathx.Sum[complex128](1.0, 2.0))
}
