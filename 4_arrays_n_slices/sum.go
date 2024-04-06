package arraysnslices

func Sum(arr []int) (sum int) {
	// like enumerate in python
	// _ means that youre ignoring the index value
	for _, num := range arr {
		sum += num
	}

	return

}

// variadic funcitons: functions that take in a variable
// number of arguments
func SumAll(numbersToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersToSum)

	// this is intializing a array of slices, that is lengthOfNumbers long ?
	// sums := make([]int, lengthOfNumbers)

	var sums []int

	for _, numbers := range numbersToSum {
		// indexing slices
		// sums[i] = Sum(numbers)

		//appending to slice
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numsToSum {
		// slicing a slice, slice[low:high] "Take from one to the end"
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))

		}

	}

	return sums
}
