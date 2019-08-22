package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	sizeInt32 = 4
)

// IndexBuffer -
type IndexBuffer struct {
	rendererID uint32
	count      int
}

// NewIndexBuffer -
// TODO: Should test using uint32 for data
func NewIndexBuffer(data []int32, count int) IndexBuffer {
	var rendererID uint32
	gl.GenBuffers(1, &rendererID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, rendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, sizeInt32*count, gl.Ptr(data), gl.STATIC_DRAW)
	return IndexBuffer{rendererID, count}
}

// Close -
func (b *IndexBuffer) Close() {
	gl.DeleteBuffers(1, &b.rendererID)
}

// Count -
func (b *IndexBuffer) Count() int {
	return b.count
}

// Bind -
func (b *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.rendererID)
}

// Unbind -
func (b *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
