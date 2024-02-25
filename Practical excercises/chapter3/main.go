package chapter3

import "fmt"

func Main() {

	//assignment
	/*
		var x int
		x = 3

		var x = 3

		x := 3
	*/
	fahrenheit_to_celcius := func(fahrenheit float64) float64 {
		return ((fahrenheit - 32) * 5 / 9)
	}
	var fahrenheit float64
	fmt.Println("Print a Fahrenheit number to convert into Celcius:")
	fmt.Scanf("%f\n", &fahrenheit)
	fmt.Println(fahrenheit, "Fahrenheit is equal to", fmt.Sprintf("%.2f", fahrenheit_to_celcius(fahrenheit)), "Celcius")

	feet_into_meters := func(feet int64) float64 {
		return 0.348 * float64(feet)
	}
	var feet int64
	fmt.Println("Type feet number to convert into meters:")
	fmt.Scanf("%d", &feet)
	fmt.Println(feet, "Feet is equal to", fmt.Sprintf("%.2f", feet_into_meters(feet)), "Meters")
}
