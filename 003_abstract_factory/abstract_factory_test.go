package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {

	// 为了更加好的描述，没有使用简略的定义方式
	//iCarFactory :=new(Factory)
	//businessCar :=iCarFactory.CreateBusinessCar()
	var iCarFactory ICarFactory
	var businessCar ICar

	iCarFactory = new(Factory)
	businessCar = iCarFactory.CreateBusinessCar()
	businessCar.SayAI()
	businessCar.DriverCar()

	var sportCar ICar
	sportCar = iCarFactory.CreateSportCar()
	sportCar.SayAI()
	sportCar.DriverCar()

}
