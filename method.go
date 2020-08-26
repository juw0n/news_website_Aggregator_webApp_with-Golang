package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

//defining or declearing a struct
type car struct {
	gasPedal      uint16 //min: 0,      max: 65535
	brakePedal    uint16 //min: 0,      max: 65535
	steeringWheel int16  //min: -32768  max: 32768
	topSpeedKmh   float64
}

//value receivers methods
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

//value receivers methods
func (c car) mph() float64 {
	//c.topSpeedKmh = 25 //The returned number will be changed, but tye original value in the struct wii not change.
	return float64(c.gasPedal) * (c.topSpeedKmh / kmhMultiple / usixteenbitmax)
}

//pointer receivers method. This will modify the struct itself via pointer.
//Now, i will modifying the struct itself via pointer
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed //this will change the original topSpeedKmh value in the car struct
}

func main() {
	aCar := car{gasPedal: 16535, brakePedal: 0, steeringWheel: 12562, topSpeedKmh: 225.0}
	//it can also declear as -- aCar := car{22314,0,12562,225.0}
	fmt.Println("gas_pedal:", aCar.gasPedal, "brake_pedal:", aCar.brakePedal, "steering_wheel:", aCar.steeringWheel)
	fmt.Println("Car is going", aCar.mph(), "MPH,", aCar.kmh(), "KMH, and top speed is", aCar.topSpeedKmh)
	aCar.newTopSpeed(500)
	fmt.Println("Car is going", aCar.mph(), "MPH,", aCar.kmh(), "KMH, and top speed is", aCar.topSpeedKmh)
}
