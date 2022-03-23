package stateLogger

type Position struct {
	X, Y int
}

type User struct {
	ID       int
	state    bool
	position Position
}

func stateLogger() {
	print("Hello, World!")
}
