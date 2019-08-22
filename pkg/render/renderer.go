package render

import "github.com/go-gl/gl/v4.1-core/gl"

type Renderer struct {
}

func (r *Renderer) Draw(va *VertexArray, ib *IndexBuffer, shader *Shader) {
	shader.Bind()
	va.Bind()
	ib.Bind()
	gl.DrawElements(gl.TRIANGLES, int32(ib.Count()), gl.UNSIGNED_INT, nil)
}

func (r *Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
