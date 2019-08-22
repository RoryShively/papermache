package render

import "github.com/go-gl/gl/v4.1-core/gl"

// VertexBuffer -
type VertexBuffer struct {
	rendererID uint32
}

// NewVertexBuffer -
func NewVertexBuffer(data interface{}, size int) VertexBuffer {
	var rendererID uint32
	gl.GenBuffers(1, &rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	return VertexBuffer{rendererID}
}

// Close -
func (b *VertexBuffer) Close() {
	gl.DeleteBuffers(1, &b.rendererID)
}

// Bind -
func (b *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.rendererID)
}

// Unbind -
func (b *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
