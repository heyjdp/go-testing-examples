package math

import "testing"

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {
	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, -4, 3},
	}

	for _, datum := range testData {
		result := Add(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Add(1, 3) FAILED. Expected %d, got %d", datum.result, result)
		} else {
			t.Logf("Add(1, 3) PASSED. Expected %d, got %d", datum.result, result)
		}
	}
}

type DivideData struct {
	x, y   int
	result float64
}

func TestDivide(t *testing.T) {
	testData := []DivideData{
		{1, 2, 0.5},
		{10, 5, 2.0},
		{5, 0, 0.0},
	}

	for _, datum := range testData {
		result := Divide(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Add(1, 3) FAILED. Expected %f, got %f", datum.result, result)
		} else {
			t.Logf("Add(1, 3) PASSED. Expected %f, got %f", datum.result, result)
		}
	}
}
