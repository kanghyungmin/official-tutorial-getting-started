// package main

// import "fmt"

// type ColorType int

// const (
// 	Red ColorType = iota
// 	Blue
// 	Green
// 	Yellow
// )

// func colorToString(color ColorType) string {
// 	switch color {
// 	case Red:
// 		return "Red"
// 	case Blue:
// 		return "Blue"
// 	case Green:
// 		return "Green"
// 	case Yellow:
// 		return "Yellow"
// 	default:
// 		return "Undefined"
// 	}
// }

// func getMyFavoriteColor() ColorType {
// 	return Blue
// }

// func main() {
// 	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
// }

package main

import (
	"fmt"
)

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(dir Direction) string {
	switch dir {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}
func getDirection(angle float64) Direction {
	switch {
	case angle > 315, (angle >= 0 && angle <= 45):
		return North
	case angle > 45 && angle <= 135:
		return East
	case angle > 135 && angle <= 225:
		return South
	case angle > 225 && angle <= 315:
		return West
	default:
		return None
	}

}

func main() {
	fmt.Println("Direction is", DirectionToString(getDirection(38.3)))
	fmt.Println("Direction is", DirectionToString(getDirection(235.8)))
	fmt.Println("Direction is", DirectionToString(getDirection(94.2)))
	fmt.Println("Direction is", DirectionToString(getDirection(-30)))

}
