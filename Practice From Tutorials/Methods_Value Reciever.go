package main

import "fmt"

const uSixteenBitMax = 65535 //This is just the upper limit of the uint16 type uused in the gasPedal
const kmhMultiple = 1.60934 //This is just the conversion of kmh to mph 

type car struct {
	gasPedal uint16 //gas pedal cannot go negative so we used uint here, where uint16 => 0 to 65535
	brakePedal uint16 //same as above
	steeringWheel int16 //steering wheel can have negatives as well so int16 => -32k to 32k
	topSpeedKmh float64
}
/*Value reciever methods are used to refer the values of various elements of a struct. They just take the value
  from the struct to use directly in various calculations.*/
func (c car) kmh() float64 { //Now this is a Value Reciever Method, where it references the struct car
							 //and creates a method of name kmh() which does look like function but isn't.
							 //And returns some value in float64.

	 return float64(c.gasPedal) * (c.topSpeedKmh/uSixteenBitMax) //here we just return the ans of a formula 
																//that calculates the current speed of the car
																//i.e. (22341*225)/65535
}

func (c car) mph() float64 {
	return (float64(c.gasPedal) * (c.topSpeedKmh/uSixteenBitMax)) / kmhMultiple
}

func main()  {
	aCar := car{gasPedal: 22341,
				brakePedal: 0,
				steeringWheel: 12561,
				topSpeedKmh: 225.0,
				}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}