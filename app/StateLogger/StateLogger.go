package StateLogger

type Position struct {
	X, Y int
}

type User struct {
	ID       int
	state    bool
	position Position
}

func StateLogger() {
	print("Hello, World!")
}
