package Find_the_highest_altitude_1732

func largestAltitude(gain []int) int {
	mt, tmp := 0, 0
	for i:=0;i<len(gain);i++ {
		tmp += gain[i]
		if tmp > mt {
			mt = tmp
		}
	}
	return mt
}