package Design_of_the_parking__system_1603

type ParkingSystem struct {
	stops [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.stops[carType-1] > 0 {
		this.stops[carType-1]--
		return true
	} else {
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
