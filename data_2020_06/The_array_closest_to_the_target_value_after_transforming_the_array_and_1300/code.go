package The_array_closest_to_the_target_value_after_transforming_the_array_and_1300

func findBestValue(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if target>arr[0] {
			return arr[0]
		} else {
			return target
		}
	}
	dest := target/n
	if target%n>n/2 {
		dest +=1
	}
	l := make([]int, 0)
	r := make([]int, 0)
	for i:=0;i<n;i++ {
		if arr[i]<=dest {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}
	if len(l) == 0 {
		return dest
	}
	return findBestValue(r, target - sum(l))
}

func sum(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}