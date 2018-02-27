package main

import "fmt"

const sixteenbitmax float64 = 65535
const miles_to_km float64 = 1.609

type car struct { //type, var, const shows they are data type, variable and constants while declaring
	pedal          uint16 //min 0 max 65535 //no negatice
	steering_wheel int16
	top_speed      float64
}

func (c car) kmh() float64 { //not a normal function so its declaration different "(c car)" part
	//associating it with structure car
	return c.top_speed
}

func main() {
	my_car := car{pedal: 341, steering_wheel: 843, top_speed: 230}
	his_car := car{437, 586, 150.56}
	fmt.Println(his_car.top_speed, my_car.pedal)
	fmt.Println(his_car.kmh())
}
