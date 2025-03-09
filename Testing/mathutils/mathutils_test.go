package mathutils

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3,5)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d but got %d",expected,result)
	}
}

Table driven testing
package mathutils

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{3, 5, 8},
		{10, -2, 8},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

//Benchmarking (go test -bench)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10,20)
	}
}
