package The_issue_of_changing_wine_1518

func numWaterBottles(numBottles int, numExchange int) int {
	// var sum int
	sum := numBottles

	for numBottles >= numExchange {
		sum =+ numBottles/numExchange
		numBottles = numBottles%numExchange + numBottles/numExchange
	}
	return sum
}