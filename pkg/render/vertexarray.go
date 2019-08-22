package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// VertexArray -
type VertexArray struct {
	rendererID uint32
}

// NewVertexArray is the constructor for VertexArray
func NewVertexArray() VertexArray {
	var va uint32
	gl.GenVertexArrays(1, &va)
	gl.BindVertexArray(va)
	return VertexArray{va}
}

// Close is the destructor for VertexArray
func (a *VertexArray) Close() {
	gl.DeleteVertexArrays(1, &a.rendererID)
}

// Bind -
func (a *VertexArray) Bind() {
	gl.BindVertexArray(a.rendererID)
}

// Unbind -
func (a *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

// AddBuffer -
func (a *VertexArray) AddBuffer(vb *VertexBuffer, layout *VertexBufferLayout) {
	a.Bind()
	vb.Bind()
	elements := layout.GetElements()
	offset := 0
	for i, elem := range elements {
		gl.EnableVertexAttribArray(uint32(i))
		gl.VertexAttribPointer(
			uint32(i),
			elem.GetCount(),
			elem.GetType(),
			elem.GetNormalized(),
			layout.GetStride(),
			// nil,
			gl.PtrOffset(offset),
		)
		offset += int(elem.GetCount() * elem.GetElemSize())
	}
}
