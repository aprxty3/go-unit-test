package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
