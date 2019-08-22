package render

import (

	// "github.com/go-gl/gl/v4.5-core/gl"

	// "github.com/RoryShively/gltest/pkg/objects"
	"fmt"

	"github.com/RoryShively/papermache/pkg/objects"
	"github.com/RoryShively/papermache/pkg/window"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Render draw to a window using an OpenGL program
func Render(win *glfw.Window) {

	circle := objects.NewCircle(0.0, 0.0, 4.0, 64)
	positions := circle.Positions
	indices := circle.Indices

	sizeFloat32 := 4

	va := NewVertexArray()
	defer va.Close()

	vb := NewVertexBuffer(positions, sizeFloat32*len(positions))
	defer vb.Close()
	vb.Bind()

	layout := VertexBufferLayout{}
	layout.Push("float", 2)
	va.AddBuffer(&vb, &layout)

	ib := NewIndexBuffer(indices, len(indices))
	defer ib.Close()
	ib.Bind()

	shader := NewShader("res/shaders/Basic.shader")
	defer shader.Close()
	shader.Bind()
	shader.SetUniform4f("u_Color", 0.8, 0.8, 0.8, 0.8)

	// Clear
	va.Unbind()
	shader.Unbind()
	vb.Unbind()
	// Clear

	proj := mgl32.Ortho(0.0, 1000.0, 0.0, 600.0, -1.0, 1.0)
	mesh := objects.NewMesh()

	renderer := Renderer{}

	shader.Bind()

	r := float32(0.0)
	i := float32(0.005)
	for !win.ShouldClose() {

		renderer.Clear()

		// shader.Bind()
		shader.SetUniform4f("u_Color", r, 0.8, 0.8, 0.8)

		// view := mgl32.Translate3D(100.0, 100.0, 0.0)
		// mvp := proj.Mul4(view)
		// shader.SetUniformMat4f("u_MVP", mvp)
		// renderer.Draw(&va, &ib, &shader)
		fmt.Printf("x: %v    y: %v\n", window.X, window.Y)

		for _, c := range mesh.Coordinates {
			// view := mgl32.Translate3D(c.X, c.Y, 0.0)
			view := mgl32.Translate3D(c.X+(1/(c.X-float32(window.X))), c.Y+(1/(c.Y-float32(window.Y))), 0.0)
			mvp := proj.Mul4(view)
			shader.SetUniformMat4f("u_MVP", mvp)
			renderer.Draw(&va, &ib, &shader)
		}

		fmt.Println(".")

		if r > 1.0 || r < 0.0 {
			i *= -1.0
		}
		r += i

		glfw.PollEvents()
		win.SwapBuffers()
	}
}
