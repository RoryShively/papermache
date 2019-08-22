package main

import (
	"log"
	"runtime"

	"github.com/RoryShively/papermache/pkg/render"
	"github.com/RoryShively/papermache/pkg/window"
	"github.com/go-gl/glfw/v3.2/glfw"
	//"github.com/acj/gonect"
)

func main() {
	log.Println("Starting OpenGL")

	runtime.LockOSThread()

	win := window.InitGlfw()
	defer glfw.Terminate()

	render.InitOpenGL()

	// for !win.ShouldClose() {
	// 	// render.Render(win, program)
	// 	render.Render(win)
	// }
	render.Render(win)

	// proj := mgl32.Ortho(0.0, 1000.0, 0.0, 600.0, -1.0, 1.0)
	// renderer := render.Renderer{}

	// mesh := objects.NewMesh(proj, renderer)

	// for !win.ShouldClose() {
	// 	renderer.Clear()
	// 	mesh.Draw()
	// 	glfw.PollEvents()
	// 	win.SwapBuffers()
	// }
}
