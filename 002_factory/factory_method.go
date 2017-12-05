package factory_method

import "fmt"

// 一辆车的基本操作
type Car interface {
	// 车上有一个会说好的人工智能
	SayAI() string

	// 车当然也是可以开的
	DriverCar();
}


// 工厂创建一辆汽车 Car的工厂接口
type CarFactory interface {
	Create() Car
}

type BenzCar struct {

}

func (car *BenzCar) SayAI() string  {
	return "I'm Benz."
}

func (car *BenzCar) DriverCar()  {
	fmt.Println("Benz Go!")
}

type BenzCarFactory struct {

}

func (factory *BenzCarFactory) Create() Car  {
	// 新的go提供了new的构造体
	car :=new(BenzCar)
	return car
}

// 还有没有结束，我还有一辆车
type TeslaCar struct {

}

func (car *TeslaCar) SayAI() string{
	return "I'm Tesla."
}

func (car *TeslaCar) DriverCar() {
	fmt.Println("Tesla Go!")
}

type TeslaCarFactory struct {
	
}

func (factory *TeslaCarFactory) Create() Car  {
	car :=new(TeslaCar)
	return car
}


