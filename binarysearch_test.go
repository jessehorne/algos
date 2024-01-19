package algos

import "testing"

func Test_BinarySearchInt(t *testing.T) {
	evenData := []int{1, 2, 5, 8, 9, 0, 10, 14}

	// test evenData with valid search value
	evenIndex := iterative_binarysearch_int(evenData, 5)
	if evenIndex != 2 {
		t.Errorf("iterative_binarysearch_int evenData index of value 5 should be 2 but was %d", evenIndex)
	}

	// test evenData with invalid value
	evenIndex = iterative_binarysearch_int(evenData, 200)
	if evenIndex != -1 {
		t.Errorf("iterative_binarysearch_int evenData index of value 200 should not be found (-1) but was %d", evenIndex)
	}

	oddData := []int{5, 10, 15, 18, 70}

	// test oddData with valid search value
	oddIndex := iterative_binarysearch_int(oddData, 15)
	if oddIndex != 2 {
		t.Errorf("iterative_binarysearch_int oddData index of value 15 should be 2 but was %d", oddIndex)
	}

	// test oddData with invalid value
	oddIndex = iterative_binarysearch_int(oddData, 200)
	if oddIndex != -1 {
		t.Errorf("iterative_binarysearch_int oddData index of value 200 should not be found (-1) but was %d", oddIndex)
	}
}
