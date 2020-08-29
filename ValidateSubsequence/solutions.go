package validatesubsequence

func main() {}

//Solution1 valid the sequence of the array and return true or false
// O(n) time | O(1) space
func Solution1(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return seqIdx == len(sequence)
}

//Solution2 valid the sequence of the array and return true or false
// O(n) time | O(1) space
func Solution2(array []int, sequence []int) bool {
	seqIdx := 0
	for _, value := range array {
		if seqIdx == len(sequence) {
			break
		}
		if value == sequence[seqIdx] {
			seqIdx++
		}
	}
	return seqIdx == len(sequence)
}
