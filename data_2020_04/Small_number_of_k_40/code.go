package Small_number_of_k_40



func getLeastNumbers(arr []int, k int) []int {
	//sort.Ints(arr)
	//return arr[:k]
	if len(arr) == 0 {
		return arr
	}
	left := 0
	right := len(arr) -1
	index := partation(arr, left, right)
	for index != k-1 {
		if index>k-1 {
			right = index-1
			index = partation(arr, left, right)
		} else {
			left = index +1
			index = partation(arr, left, right)
		}
	}
	result := []int{}
	for i :=0; i<k;i++ {
		result = append(result, arr[i])
	}
	return result
}


func partation(num []int, left, right int) int {
	if left == right {
		return right
	}
	if left <right {
		base := num[left]
		l := left
		r := right
		for {
			for num[r] >=base && l<r {
				r--
			}
			for num[l] <=base && l<r {
				l++
			}
			if l >= r {
				break
			}
			num[l], num[r] = num[r], num[l]
		}
		num[left] = num[l]
		num[l] = base
	}
	return -1
}