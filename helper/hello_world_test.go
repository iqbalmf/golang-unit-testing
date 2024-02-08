package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// jalankan dengan "go test -v -r 'name_func'" di terminal
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hi Iqbal" {
		// panic("Result is not Hi Iqbal")
		t.Fail() //unit test running, melanjutkan ke line selanjutnya
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldFauzan(t *testing.T) {
	result := HelloWorld("Fauzan")
	if result != "Hi Fauzan" {
		// panic("Result is not Hi Fauzan")
		t.FailNow() //otomatis unit test error, tidak dilanjutkan
	}
	fmt.Println("TestHelloWorldFauzan Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Iqbal")
	assert.Equal(t, "Hello Iqbal", result, "RESULT MUST 'Hello Iqbal'")
	fmt.Println("TestHelloWorldAssert with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Iqbal")
	require.Equal(t, "Hello Iqbals", result, "RESULT MUST 'Hello Iqbal'")
	fmt.Println("TestHelloWorldRequire with Require Done")
}
