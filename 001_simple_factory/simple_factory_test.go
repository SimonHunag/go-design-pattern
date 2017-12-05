package simple_factory

import "testing"

func TestBenzCar(t *testing.T) {

	benzCar := NewCar(1)
	benzCar.drive();
}


func TestBYDCar(t *testing.T){

	bYDCar :=NewCar(2)
	bYDCar.drive()
}