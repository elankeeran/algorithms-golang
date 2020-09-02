package createnumbersequencefrom2dslice


import (
	"sort"
)

func findSequenceNumber(input [][]int) []int {
	input = sort2d(input)
	output := []int{}

	for len(input) > 0 {
		firstRow := input[0]
		output = append(output, firstRow[0])
		firstRow = firstRow[1:]
		input = input[1:]
		if len(firstRow) != 0 {
			s := binarySearch(input, firstRow[0])
			idx := s[0]
			input = append(input, []int{0})
			copy(input[idx+1:], input[idx:])
			input[idx] = firstRow
		}
	}
	return output
}

func sort2d(input [][]int) [][]int {
	sort.Slice(input, func(i, j int) bool {
		// edge cases
		if len(input[i]) == 0 && len(input[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(input[i]) == 0 || len(input[j]) == 0 {
			return len(input[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}
		// both slices len() > 0, so can test this now:
		return input[i][0] < input[j][0]
	})

	return input
}

func binarySearch(a [][]int, v int) []int {
	low := 0
	high := len(a) - 1
	for low <= high {
		median := (low + high) / 2
		if a[median][0] < v {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(a) || a[low][0] != v {
		return []int{low, high}
	}
	return []int{low, high}
}

func main() {}

//below line of code also gives same output
// o := []int{}
// for i := 0; i < len(input); i++ {
// 	for j := 0; j < len(input[i]); j++ {
// 		o = append(o, input[i][j])
// 	}
// }
// sort.Ints(o)
// fmt.Println(o)
