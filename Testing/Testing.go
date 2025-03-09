// Go has a built-in testing framework using the testing package.
// Test files must end with _test.go (e.g., mathutils_test.go).

//File: mathutils.go//
package mathutils

// Function to add two numbers
func Add(a, b int) int {
	return a + b
}

//File: mathutils_test.go//
package mathutils

import "testing"

// Unit test for Add function
func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
//

Run test: go test

output: ok mathutils 0.001s

✅ Key Takeaways:

t.Errorf() reports test failures.
go test runs all _test.go files.

// 3. Table-Driven Tests(Multiple cases in one Test)

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

4. Benchmarking (go test-bench)

Go allows benchmarking functions using testing.B.

package mathutils

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

Run: go test -bench .

✅ Key Takeaways:

b.N is the number of iterations for performance testing.
go test -bench . runs all benchmarks.

5. Test Coverage (go test -cover)

go test -cover

✅ Key Takeaways:

Helps find untested parts of your code.

6. Using t.Run() for subsets
Subtests allow organizing tests logically.

package mathutils

import "testing"

func TestOperations(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		if Add(3, 5) != 8 {
			t.Error("Addition failed")
		}
	})

	t.Run("Zero Case", func(t *testing.T) {
		if Add(0, 0) != 0 {
			t.Error("Zero case failed")
		}
	})
}

7. using setup and teardown (TestMain)
TestMain runs before and after all tests.

package mathutils

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup before tests")
	exitVal := m.Run() // Run all tests
	fmt.Println("Teardown after tests")
	os.Exit(exitVal)
}

✅ Key Takeaways:

TestMain is useful for database connections, mocks, and cleanup.

Exercises
Write tests for a Multiply(a, b int) int function.
Modify TestAdd() to use subtests with t.Run().
Add a benchmark for Multiply().