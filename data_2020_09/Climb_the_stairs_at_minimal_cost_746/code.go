package Climb_the_stairs_at_minimal_cost_746

func minCostClimbingStairs(cost []int) int {
	f0 := 0
	f1 := 0
	f2 := 0
	for i:= len(cost)-1;i>=0;i-- {
		if f1<f2 {
			f0 = cost[i] +f1
		} else {
			f0 = cost[i] + f2
		}
		f2 = f1
		f1= f0
	}
	if f1 < f2 {
		return f1
	} else {
		return f2
	}
}