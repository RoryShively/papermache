package render

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexBufferElement interface {
	GetCount() int32
	GetType() uint32
	GetNormalized() bool
	GetElemSize() int32
}

type UIntBufferElement struct {
	count int32
}

func (e *UIntBufferElement) GetType() uint32     { return gl.UNSIGNED_INT }
func (e *UIntBufferElement) GetCount() int32     { return e.count }
func (e *UIntBufferElement) GetNormalized() bool { return false }
func (e *UIntBufferElement) GetElemSize() int32  { return 4 }

type FloatBufferElement struct {
	count int32
}

func (e *FloatBufferElement) GetType() uint32     { return gl.FLOAT }
func (e *FloatBufferElement) GetCount() int32     { return e.count }
func (e *FloatBufferElement) GetNormalized() bool { return false }
func (e *FloatBufferElement) GetElemSize() int32  { return 4 }

type UCharBufferElement struct {
	count int32
}

func (e *UCharBufferElement) GetType() uint32     { return gl.UNSIGNED_BYTE }
func (e *UCharBufferElement) GetCount() int32     { return e.count }
func (e *UCharBufferElement) GetNormalized() bool { return true }
func (e *UCharBufferElement) GetElemSize() int32  { return 1 }

// VertexBufferLayout -
type VertexBufferLayout struct {
	elements []VertexBufferElement
	stride   int32
}

// func NewVertexBufferLayout(stride int32) VertexBufferLayout {
// 	return VertexBufferLayout{[]VertexBufferElement{}, stride}
// }

func (l *VertexBufferLayout) GetElements() []VertexBufferElement {
	return l.elements
}

func (l *VertexBufferLayout) GetStride() int32 {
	return l.stride
}

// func (l *VertexBufferLayout) Push(elem VertexBufferElement) {
// 	l.elements = append(l.elements, elem)
// 	l.stride += elem.GetElemSize()
// }

func (l *VertexBufferLayout) Push(elemType string, count int32) {
	var elem VertexBufferElement
	switch elemType {
	case "float":
		elem = &FloatBufferElement{count}
	case "uint":
		elem = &UIntBufferElement{count}
	case "uchar":
		elem = &UCharBufferElement{count}
	default:
		log.Fatal("Invalid elemType passed to VertexBufferLayout.Push()")
	}

	l.elements = append(l.elements, elem)
	l.stride += count * elem.GetElemSize()
}
