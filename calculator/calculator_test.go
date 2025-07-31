package calculator

import (
	"testing"
)

func TestNew(t *testing.T) {
	calc := New()

	if calc == nil {
		t.Error("New() returned nil")
	}

	expected := "Basic calculator v1.0"
	if calc.String() != expected {
		t.Errorf("String() = %q, wanted %q", calc.String(), expected)
	}
}

func TestCalculator_Add(t *testing.T) {
	calc := New()

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 2.5, 3.5, 6.0},
		{"negative numbers", -5.0, -3.0, -8.0},
		{"mixed signs", -10.0, 15.0, 5.0},
		{"zero values", 0.0, 5.0, 5.0},
		{"large numbers", 999999.99, 0.01, 1000000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := New()

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive result", 10.0, 3.0, 7.0},
		{"negative result", 3.0, 10.0, -7.0},
		{"zero result", 5.0, 5.0, 0.0},
		{"subtracting negative", 5.0, -3.0, 8.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := New()

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 4.0, 5.0, 20.0},
		{"negative numbers", -3.0, -4.0, 12.0},
		{"mixed signs", -6.0, 2.0, -12.0},
		{"multiply by zero", 100.0, 0.0, 0.0},
		{"multiply by one", 7.5, 1.0, 7.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := New()

	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"normal division", 10.0, 2.0, 5.0, false},
		{"division by zero", 10.0, 0.0, 0.0, true},
		{"negative numbers", -15.0, 3.0, -5.0, false},
		{"result less than 1", 1.0, 4.0, 0.25, false},
		{"divide zero", 0.0, 5.0, 0.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestCalculator_DivideByZeroErrorMessage(t *testing.T) {
	calc := New()

	_, err := calc.Divide(10.0, 0.0)
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
		return
	}

	expectedMsg := "division by zero is not allowed"
	if err.Error() != expectedMsg {
		t.Errorf("Error message = %q, want %q", err.Error(), expectedMsg)
	}
}
