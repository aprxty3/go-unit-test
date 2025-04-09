package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}
	result := HelloWorld("Asset")
	assert.Equal(t, "Hello Asset", result, "Result must be Hello Asset")

}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Asset")
	assert.Equal(t, "Hello Asset", result, "Result must be Hello Asset")
	fmt.Println("TestHelloWorldAsset Done")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Asset")
	require.Equal(t, "Hello Asset", result, "Result must be Hello Asset")
	fmt.Println("TestHelloWorldAsset Done")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Enr")
	if result != "Hello Enr" {
		t.Error("Result must be Hello Enr")
	}

	fmt.Println("TestHelloWorld Done")

}

func TestHelloWorldAji(t *testing.T) {
	result := HelloWorld("Aji")
	if result != "Hello Aji" {
		t.Fatal("Result must be Hello Aji")
	}

	fmt.Println("TestHelloWorldAji Done")
}

func TestMain(m *testing.M) {
	// setup code
	fmt.Println("Before Unit Test")
	m.Run()
	// teardown code
	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Aji", func(t *testing.T) {
		result := HelloWorld("Aji")
		if result != "Hello Aji" {
			t.Fatal("Result must be Hello Aji")
		}
	})

	t.Run("Enr", func(t *testing.T) {
		result := HelloWorld("Enr")
		if result != "Hello Enr" {
			t.Fatal("Result must be Hello Enr")
		}
	})
}

func TestTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected string
	}{
		{name: "Aji", age: 20, expected: "Hello Aji"},
		{name: "Enr", age: 30, expected: "Hello Enr"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.name)
			assert.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}

}
