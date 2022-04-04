package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 10) != 11 {
		t.Error("Expected 1 (+) 10 = 11")
	}
	if Addition(-1, -10) != -11 {
		t.Error("Expected -1 (+) -10 = -11")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(1, 10) != -9 {
		t.Error("Expected 1 (-) 10 = -9")
	}
	if Subtraction(-1, -10) != 9 {
		t.Error("Expected -1 (-) -10 = 9")
	}
}

func TestDivision(t *testing.T) {
	if Division(10, 2) != 5 {
		t.Error("Expected 10 (:) 2 = 5")
	}
	if Division(-10, 2) != -5 {
		t.Error("Expected -10 (:) 2 = -5")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(1, 10) != 10 {
		t.Error("Expected 1 (x) 10 = 10")
	}
	if Multiplication(-1, 10) != -10 {
		t.Error("Expected -1 (x) 10 = -10")
	}
}
