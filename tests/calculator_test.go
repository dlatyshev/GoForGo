package main

import "testing"
import "github.com/dlatyshev/GoForGo/calc"

func TestAdd(t *testing.T) {
	result := calc.Add(2.0, 3.1)
	if result != 5.1 {
		t.Errorf("Expected 5.1, got %.2f", result)
	}
}

func TestSubtract(t *testing.T) {
	result := calc.Subtract(5.0, 2.0)
	if result != 3.0 {
		t.Errorf("Expected 3.0, got %.2f", result)
	}
}

func TestMultiply(t *testing.T) {
	result := calc.Multiply(2.0, 3.0)
	if result != 6.0 {
		t.Errorf("Expected 6.0, got %.2f", result)
	}
}

func TestDivide(t *testing.T) {
	result := calc.Divide(6.0, 2.0)
	if result != 3.0 {
		t.Errorf("Expected 3.0, got %.2f", result)
	}

	result = calc.Divide(6.0, 0.0)
	if result != 0.0 {
		t.Errorf("Expected 0.0, got %.2f", result)
	}
}
