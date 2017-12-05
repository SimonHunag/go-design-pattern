package factory_method

import (
	"testing"
	"fmt"
)

func TestBenzCar(t *testing.T) {

	var carFactory CarFactory
	carFactory = new(BenzCarFactory)
	benzCar :=carFactory.Create()
	fmt.Println(benzCar.SayAI())
	benzCar.DriverCar()


	carFactory = new(TeslaCarFactory)
	teslaCar :=carFactory.Create()
	fmt.Println(teslaCar.SayAI())
	teslaCar.DriverCar()

}
