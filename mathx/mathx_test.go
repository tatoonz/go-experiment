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
