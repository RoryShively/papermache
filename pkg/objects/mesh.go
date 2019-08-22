package objects

type Coordinates struct {
	X float32
	Y float32
}

type Mesh struct {
	// proj     mgl32.Mat4
	// renderer render.Renderer
	// Circles []Circle
	Coordinates []Coordinates
}

// func NewMesh(proj mgl32.Mat4, renderer render.Renderer) Mesh {
func NewMesh() Mesh {

	mesh := Mesh{
		// proj:     proj,
		// renderer: renderer,
	}
	mesh.init()
	return mesh
}

func (m *Mesh) init() {
	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 60; j++ {
			// circle := NewCircle(float64(i)*10.0, float64(j)*10.0, 10.0, 18)
			// m.Circles = append(m.Circles, circle)
			c := Coordinates{
				X: float32(i) * 10.0,
				Y: float32(j) * 10.0,
			}
			m.Coordinates = append(m.Coordinates, c)
		}
	}
}

// func (m *Mesh) Draw() {
// 	m.renderer.Clear()
// 	for _, c := range m.Circles {
// 		m.renderer.Draw(&c.va, &c.ib, &c.shader)
// 	}
// }
