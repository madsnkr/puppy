package puppy

import "github.com/madsnkr/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func Bark1() string {
	return "Hello from v1"
}

func Bark2() string {
	return "Hello from v1.2"
}

func Bark3() string {
	return "Hello from v1.3"
}
