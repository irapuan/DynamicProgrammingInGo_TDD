package main

import "testing"

func TestFibonacci(context *testing.T) {
	context.Run("Fib(1) returns 1", func(t *testing.T) {
		input := 1
		expected := 1
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(2) returns 1", func(t *testing.T) {
		input := 2
		expected := 1
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(3) returns 2", func(t *testing.T) {
		input := 3
		expected := 2
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(4) returns 3", func(t *testing.T) {
		input := 4
		expected := 3
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(5) returns 5", func(t *testing.T) {
		input := 5
		expected := 5
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(6) returns 8", func(t *testing.T) {
		input := 6
		expected := 8
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(7) returns 13", func(t *testing.T) {
		input := 7
		expected := 13
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

}

func TestFibonacciMeomization(context *testing.T) {
	context.Run("Fib(1) returns 1", func(t *testing.T) {
		input := 1
		expected := 1
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(2) returns 1", func(t *testing.T) {
		input := 2
		expected := 1
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(3) returns 2", func(t *testing.T) {
		input := 3
		expected := 2
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(5) returns 5", func(t *testing.T) {
		input := 5
		expected := 5
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(6) returns 8", func(t *testing.T) {
		input := 6
		expected := 8
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("Fib(7) returns 13", func(t *testing.T) {
		input := 7
		expected := 13
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

}

func BenchmarkFibonacci(context *testing.B) {

	context.Run("Fib(43) returns 433494437", func(t *testing.B) {
		input := 43
		expected := 433494437
		memo := []int{}
		returned := Fib(input, &memo)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})

	context.Run("FibBruteForce(43) returns 433494437", func(t *testing.B) {
		input := 43
		expected := 433494437
		returned := FibBruteForce(input)
		if expected != returned {
			t.Errorf("Value expected %d is different from value returned %d", expected, returned)
		}

	})
}
