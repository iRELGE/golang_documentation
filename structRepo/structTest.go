package structRepo

type Car struct {
	Gas_pedal     int
	Brake_pedal   int
	Steering_well int
}

func CallStruct() (Car_a Car) {
	Car_a = Car{}

	return
}
