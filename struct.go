package main

import "fmt"

//defining or declearing a struct
type car struct {
	gasPedal      uint16 //min: 0,      max: 65535
	brakePedal    uint16 //min: 0,      max: 65535
	steeringWheel int16  //min: -32768  max: 32768
	topSpeedKmh   float64
}

func main() {
	aCar := car{gasPedal: 16535, brakePedal: 0, steeringWheel: 12562, topSpeedKmh: 225.0}
	//it can also declear as -- aCar := car{22314,0,12562,225.0}
	fmt.Println("gas_pedal:", aCar.gasPedal, "brake_pedal:", aCar.brakePedal, "steering_wheel:", aCar.steeringWheel)
}
