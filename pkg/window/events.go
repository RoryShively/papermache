package window

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	X float64 = 0
	Y float64 = 0
)

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action == glfw.Press {
		fmt.Println("Clicked")
	} else if action == glfw.Release {
		fmt.Println("Released")
	}
}

func cursorPosCallback(w *glfw.Window, xpos float64, ypos float64) {
	// fmt.Printf("x: %v   y: %v", xpos, ypos)a
	X = xpos
	Y = ypos
}
