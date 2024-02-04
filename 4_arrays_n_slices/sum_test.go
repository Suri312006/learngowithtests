package arraysnslices

import "testing"

func TestSum(t *testing.T) {
	// t.Run("Collection of 5 numbers", func(t *testing.T) {
	//
 //        // this is a slice in the go language, where the length of the array
 //        // doesnt need to be known 
	// 	numbers := []int{1, 2, 3, 4, 5}
	//
	// 	got := Sum(numbers)
	// 	want := 15
	//
	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	//
	// })
	t.Run("Collection of any size ", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}
