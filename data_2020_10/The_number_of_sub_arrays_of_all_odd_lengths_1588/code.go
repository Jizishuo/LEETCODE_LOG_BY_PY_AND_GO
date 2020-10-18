package The_number_of_sub_arrays_of_all_odd_lengths_1588

func sumOddLengthSubarrays(arr []int) int {

	sum := 0
	for i:=1;i<=len(arr);i+=2 {
		for j:=0;j+i<= len(arr);j++ {
			sum += sumStoE(arr, j, i)
		}
	}

	return sum

}

func sumStoE(arr []int, S,num int) int {
	sum := 0

	if arr != nil {
		for i:= S; i< S+num; i++ {
			sum += arr[i]
		}
	}

	return sum
}
