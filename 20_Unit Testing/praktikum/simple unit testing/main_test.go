package main

import "testing"

func TestAdd(t *testing.T) {
	var expected int = 6

	var actual int = Add(2, 4)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)

	}

}

func TestSubtract(t *testing.T) {
	var expected int = 2

	var actual int = Subtract(4, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)

	}

}

func TestMultiple(t *testing.T) {
	var expected int = 8

	var actual int = Multiple(4, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)

	}

}

func TestDivision(t *testing.T) {
	var expected int = 2

	var actual int = Division(4, 2)

	if expected != actual {
		t.Errorf("expected: %v got: %v", expected, actual)

	}

}
