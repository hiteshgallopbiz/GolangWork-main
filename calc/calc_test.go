package calc_test

import (
	"testing"

	"calculater/calc"
)

func TestAdditionInt(t *testing.T) {
	input := calc.Input{Num1: 5, Num2: 6}
	expected := 11
	result := input.Addition()

	if result != expected {
		t.Errorf("AdditionInt() failed: expected %d, but got %d", expected, result)
	}
}

func TestAdditionFloat(t *testing.T) {
	input := calc.Inputfloat{Numfloat1: 5.15, Numfloat2: 6.34}
	expected := 11.49
	result := input.AdditionFloat()

	if result != expected {
		t.Errorf("Addition was incorrect, got: %f, expected: %f.", result, expected)
	}
}

func BenchmakAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := calc.Input{Num1: 7, Num2: 8}
		input.Addition()
	}

}
