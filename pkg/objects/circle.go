package objects

import (
	"fmt"
	"math"
	// "github.com/RoryShively/papermache/pkg/render"
)

type Circle struct {
	Positions []float32
	Indices   []int32

	radius       float64
	segmentCount int
	posX         float64
	posY         float64
	theta        float64

	// va     render.VertexArray
	// vb     render.VertexBuffer
	// ib     render.IndexBuffer
	// layout render.VertexBufferLayout
	// shader render.Shader
}

func NewCircle(x float64, y float64, radius float64, segmentCount int) Circle {
	theta := 2 * math.Pi / float64(segmentCount)

	circle := Circle{
		radius:       radius,
		segmentCount: segmentCount,
		posX:         x,
		posY:         y,
		theta:        theta,
		Positions:    []float32{float32(x), float32(y)},
	}
	circle.initiate()

	return circle
}

func (c *Circle) initiate() {
	fmt.Printf(`
		radius: %v
		segments: %v
	`, c.radius, c.segmentCount)
	for i := 0; i <= c.segmentCount; i++ {
		x := float32(c.posX + c.radius*math.Cos(c.theta*float64(i)))
		y := float32(c.posY + c.radius*math.Sin(c.theta*float64(i)))

		c.Positions = append(c.Positions, x, y)
		c.Indices = append(c.Indices, 0, int32((i)+1), int32((i)+2))
	}

	// TODO: Fix this workaround that deletes the last 3 indices
	c.Indices = c.Indices[:(len(c.Indices) - 3)]
	// fmt.Println(c.Positions)
	// fmt.Println(c.Indices)

	// sizeFloat32 := 4

	// c.va = render.NewVertexArray()
	// defer c.va.Close()

	// c.vb = render.NewVertexBuffer(c.Positions, sizeFloat32*len(c.Positions))
	// defer c.vb.Close()
	// c.vb.Bind()

	// c.layout = render.VertexBufferLayout{}
	// c.layout.Push("float", 2)
	// c.va.AddBuffer(&c.vb, &c.layout)

	// c.ib = render.NewIndexBuffer(c.Indices, len(c.Indices))
	// defer c.ib.Close()
	// c.ib.Bind()

	// c.shader = render.NewShader("res/shaders/Basic.shader")
	// defer c.shader.Close()
	// c.shader.Bind()
	// c.shader.SetUniform4f("u_Color", 0.8, 0.8, 0.8, 0.8)
}
