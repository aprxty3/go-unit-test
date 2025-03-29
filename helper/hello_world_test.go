package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Aji")
	if result != "HelloAji" {
		panic("Result is not Hello Aji")
	}

}
