package main

import "fmt"

type Engine struct {
	fuel    int8
	battery int8
}

func (e *Engine) work() {
	if e.fuel <= 0 && e.battery <= 0 {
		return
	}
	e.fuel -= 1
	e.battery -= 1
}

type Car struct {
	engine Engine
}

func (c *Car) drive(s int8) string {
	var j, i int
	var message string
	switch {
	case s == 20:
		j = 3
		message = "You're driving slowly"
	case s == 40:
		j = 5
		message = "You're driving normally"
	case s == 80:
		j = 8
		message = "You're going fast"
	case s == 120:
		j = 12
		message = "Slow down!"
	}
	for i = 0; i <= j; i++ {
		c.engine.work()
	}
	return message
}
func main() {
	var f, b, speed int8
	fmt.Println("Refuel the car.")
	fmt.Println("Enter the amount of fuel")
	fmt.Scan(&f)
	fmt.Println("Enter the amount of battery")
	fmt.Scan(&b)
	car := Car{Engine{f, b}}

	fmt.Println("Enter speed: 20, 40, 80, 120")
	fmt.Scan(&speed)
	fmt.Println(car.drive(speed))
	fmt.Println(car)

}
