package render

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// InitOpenGL initializes OpenGL and returns an intialized program.
func InitOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}
