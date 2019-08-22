package window

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 1000
	height = 600
)

// InitGlfw initializes glfw and returns a Window to use.
func InitGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "OpenGL Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetCursorPosCallback(cursorPosCallback)

	return window
}
