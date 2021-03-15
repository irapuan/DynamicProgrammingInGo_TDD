package main

import "testing"

func Test_algo(context *testing.T) {

	context.Run("returns true in case the sum of the values in the array returns the interested kvalue", func(t *testing.T) {
		a := []int{10, 15, 3, 7}
		k := 17

		res_to_test := Func_for_test(a, k)

		if res_to_test != true {
			t.Errorf("Expected true but got false")
		}
	})

	context.Run("returns false in case the sum of the values in the array returns the interested kvalue", func(t *testing.T) {
		a := []int{15, 3, 7}
		k := 17

		res_to_test := Func_for_test(a, k)

		if res_to_test == true {
			t.Errorf("Expected false, but got true")
		}
	})
}
