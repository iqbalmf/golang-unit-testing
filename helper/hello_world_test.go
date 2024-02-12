package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Iqbal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Iqbal")
		}
	})
	b.Run("Fauzan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fauzan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Iqbal")
	}
}

func BenchmarkHelloWorldTable(b *testing.B) {
	testBench := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Iqbal",
			request:  "Iqbal",
			expected: "Hello Iqbal",
		},
		{
			name:     "Fauzan",
			request:  "Fauzan",
			expected: "Hello Fauzan",
		},
	}
	for _, benchmark := range testBench {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before UNIT TEST")
	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

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
	require.Equal(t, "Hello Iqbal", result, "RESULT MUST 'Hello Iqbal'")
	fmt.Println("TestHelloWorldRequire with Require Done")
}

func TestHelloWorldSkipTest(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := HelloWorld("Iqbal")
	require.Equal(t, "Hello Iqbal", result, "Result must be 'Hello Iqbal'")
}

func TestSubTest(t *testing.T) {
	t.Run("Iqbal", func(t *testing.T) {
		result := HelloWorld("Iqbal")
		require.Equal(t, "Hello Iqbal", result, "Result Must be 'Hello Iqbal")
	})
	t.Run("Fauzan", func(t *testing.T) {
		result := HelloWorld("Fauzan")
		require.Equal(t, "Hello Fauzan", result, "Result Must be 'Hello Fauzan")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Iqbal",
			request:  "Iqbal",
			expected: "Hello Iqbal",
		},
		{
			name:     "Fauzan",
			request:  "Fauzan",
			expected: "Hello Fauzan",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
