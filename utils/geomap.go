package utils

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type Coord struct {
	x int
	y int
}

func Geomap() []Coord {
	// Top left, top right, bottom right, bottom left
	// For 4 times, wait for mouse click.
	// If clicked, get current loc, and append to array
	fmt.Println("Mapper started!")

	eventChan := make(chan Coord)
	go listen(eventChan)

	coords := make([]Coord, 4)

	for i := 0; i < 4; i++ {
		v, ok := <-eventChan
		if !ok {
			panic("BRUH")
		}
		coords[i] = v
		fmt.Println("Added", v, "to coords")
	}

	return coords
}

func listen(eventChan chan Coord) {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		x, y := robotgo.Location()
		fmt.Println("Clicked ", x, y)
		eventChan <- Coord{x, y}
	})

	s := hook.Start()
	<-hook.Process(s)
}
