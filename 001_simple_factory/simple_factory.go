package simple_factory

import "fmt"

type Car interface {
	drive()
}

type BenzCar struct {

}


func (* BenzCar) drive(){
	fmt.Printf("Driving Benz")
}

type BYDCar struct {

}

func (* BYDCar) drive(){
	fmt.Printf("Driving BYD")
}

func NewCar(car int) Car{

	if car == 1 {
		return new(BenzCar)
	}else if car == 2 {
		return new(BYDCar)
	}
	return nil;

}
