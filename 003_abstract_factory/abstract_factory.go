package abstract_factory

import "fmt"

type ICar interface {
	// 车上有一个会说好的人工智能
	SayAI()

	// 车当然也是可以开的
	DriverCar();
}

type SportCar struct {
}

func (*SportCar) SayAI() {
	fmt.Println("I'm sport car")
}

func (*SportCar) DriverCar() {
	fmt.Println("Go!Go! Sport Car.")
}

type BusinessCar struct {
}

func (*BusinessCar) SayAI() {
	fmt.Println("I'm business car")
}

func (*BusinessCar) DriverCar() {
	fmt.Println("Go!Go! Business Car.")
}

type ICarFactory interface {
	CreateBusinessCar() ICar
	CreateSportCar() ICar
}

type Factory struct {
}

func (*Factory) CreateBusinessCar() ICar {
	return new(BusinessCar)
}

func (*Factory) CreateSportCar() ICar {
	return new(SportCar)
}
