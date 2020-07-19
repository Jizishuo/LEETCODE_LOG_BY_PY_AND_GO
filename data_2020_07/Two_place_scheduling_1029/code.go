package Two_place_scheduling_1029

func twoCitySchedCost(costs [][]int) int {
	for i:=0;i<len(costs);i++{
		for j:=i+1;j<len(costs);j++ {
			if (costs[i][0]-costs[i][1])>(costs[j][0]-costs[j][1]) {
				costs[i], costs[j]=costs[j], costs[i]
			}
		}
	}
	sum := 0
	for k:=0;k<len(costs)/2;k++ {
		sum +=costs[k][0]
	}
	for p:=len(costs)/2;p<len(costs);p++ {
		sum += costs[p][1]
	}
	return sum
}