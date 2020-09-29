package Minimum_time_to_access_all_points_1266

func minTimeToVisitAllPoints(points [][]int) int {
	t := 0
	for i:=0;i<len(points)-1;i++ {
		xT := abs(points[i][0]-points[i+1][0])
		yT := abs(points[i][1]-points[i+1][1])
		if xT>yT {
			t += xT
		} else {
			t += yT
		}
	}
	return t
}

func abs(x int) int {
	if x<0{
		x *= -1
	}
	return x
}