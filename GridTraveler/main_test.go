package main

import "testing"

func TestGridTravelerBruteForce(context *testing.T) {
	context.Run("Give an empty array should return 0", func(t *testing.T) {
		result := GridTravelerBruteForce(0, 0)
		if result != 0 {
			t.Errorf("Expected 0 got %d ", result)
		}
	})

	context.Run("Give an one by one matrix should return 1", func(t *testing.T) {
		expected := 1
		result := GridTravelerBruteForce(1, 1)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][]
	context.Run("Give an 3x1 matrix should return 1", func(t *testing.T) {
		expected := 1
		result := GridTravelerBruteForce(3, 1)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// []
	// []
	// []
	context.Run("Give an 1x3 matrix should return 1", func(t *testing.T) {
		expected := 1
		result := GridTravelerBruteForce(1, 3)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][]
	// [][]
	context.Run("Give an 2x2 matrix should return 2", func(t *testing.T) {
		expected := 2
		result := GridTravelerBruteForce(2, 2)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][]
	// [][]
	// [][]
	context.Run("Give an 2x3 matrix should return 3", func(t *testing.T) {
		expected := 3
		result := GridTravelerBruteForce(2, 3)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][]
	// [][][]
	context.Run("Give an 3x2 matrix should return 3", func(t *testing.T) {
		expected := 3
		result := GridTravelerBruteForce(3, 2)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][]
	// [][][]
	// [][][]
	context.Run("Give an 3x3 matrix should return 6", func(t *testing.T) {
		expected := 6
		result := GridTravelerBruteForce(3, 3)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][][]
	// [][][][]
	context.Run("Give an 4x2 matrix should return 4", func(t *testing.T) {
		expected := 4
		result := GridTravelerBruteForce(4, 2)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][]
	// [][]
	// [][]
	// [][]
	context.Run("Give an 2x4 matrix should return 4", func(t *testing.T) {
		expected := 4
		result := GridTravelerBruteForce(2, 4)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][][]
	// [][][][]
	// [][][][]
	context.Run("Give an 4x3 matrix should return 10", func(t *testing.T) {
		expected := 10
		result := GridTravelerBruteForce(4, 3)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][]
	// [][][]
	// [][][]
	// [][][]
	context.Run("Give an 3x4 matrix should return 10", func(t *testing.T) {
		expected := 10
		result := GridTravelerBruteForce(3, 4)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

	// [][][][]
	// [][][][]
	// [][][][]
	// [][][][]
	context.Run("Give an 4x4 matrix should return 20", func(t *testing.T) {
		expected := 20
		result := GridTravelerBruteForce(4, 4)
		if result != expected {
			t.Errorf("Expected %d got %d ", expected, result)
		}
	})

}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		m        int
		n        int
		expected int
	}{
		{
			desc:     "Should return zero when provided zeros",
			m:        0,
			n:        0,
			expected: 0,
		},
		{
			desc:     "Should return 20 when provided 4x4",
			m:        4,
			n:        4,
			expected: 20,
		},
		{
			desc:     "Should return 10 when provided 3x4",
			m:        3,
			n:        4,
			expected: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			memo := make(map[string]int)
			res := GridTraveler(tC.m, tC.n, &memo)
			if res != tC.expected {
				t.Errorf("Expected: %d, for %dx%d , but got %d", tC.expected, tC.m, tC.n, res)
			}
		})
	}
}

func BenchmarkGridTraveler(b *testing.B) {
	b.Run("Dynamic version", func(b *testing.B) {
		memo := make(map[string]int)
		GridTraveler(18, 18, &memo)
	})

	b.Run("Brute force version", func(b *testing.B) {
		GridTravelerBruteForce(18, 18)
	})
}
