package helper

import (
	"fmt"
	"testing"
)

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
