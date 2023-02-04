package ParkingSystem

type ParkingSystem struct {
	Arr []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Arr: []int{0, big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.Arr[carType] <= 0 {
		return false
	}
	this.Arr[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
