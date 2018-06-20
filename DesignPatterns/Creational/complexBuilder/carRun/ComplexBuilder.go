package main

import (
	"Go-Train-ME/DesignPatterns/Creational/complexBuilder"
)

func main() {
	assembly := car.NewBuilder().Paint(car.RedColor)

	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Drive()
}
