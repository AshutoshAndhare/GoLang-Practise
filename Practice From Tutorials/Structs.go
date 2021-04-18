package main

import "fmt"

type car struct {
	gasPedal uint16 //gas pedal cannot go negative so we used uint here, where uint16 => 0 to 65535
	brakePedal uint16 //same as above
	steeringWheel int16 //steering wheel can have negatives as well so int16 => -32k to 32k
	topSpeedKmh float64
}

func main()  {
	aCar := car{gasPedal: 22341,
				brakePedal: 0,
				steeringWheel: 12561,
				topSpeedKmh: 255.0,
				}

	fmt.Println(aCar.gasPedal)
}