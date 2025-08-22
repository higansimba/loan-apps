package calculator_test

import (
	"testing"

	"github.com/higansama/loan-apps/calculator"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(5, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Add(5, 3) = %v; want %v", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := calculator.Subtract(10, 4)
	expected := 6.0
	if result != expected {
		t.Errorf("Subtract(10, 4) = %v; want %v", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := calculator.Multiply(2, 4)
	expected := 8.0
	if result != expected {
		t.Errorf("Multiply(2, 4) = %v; want %v", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result, err := calculator.Divide(10, 2)
	expected := 5.0
	if err != nil || result != expected {
		t.Errorf("Divide(10, 2) = %v, %v; want %v, nil", result, err, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := calculator.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error but got nil")
	}
}
