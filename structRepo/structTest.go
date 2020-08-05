package structRepo

type Car struct {
	gas_pedal     int
	brake_pedal   int
	steering_well int
}

func CallStruct(a, b, c int) (Car_a Car) {
	Car_a = Car{a, b, c}

	return
}
