package render

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	// "github.com/go-gl/gl/v4.5-core/gl"
)

// ShaderProgramSource -
type ShaderProgramSource struct {
	VertexShader   string
	FragmentShader string
}

// Shader -
type Shader struct {
	rendererID uint32
	filepath   string
	uCache     map[string]int32
}

// NewShader -
func NewShader(filepath string) Shader {
	shader := Shader{
		filepath: filepath,
		uCache:   make(map[string]int32, 32),
	}
	source := shader.parseShader()
	shader.createShader(source)

	return shader
}

// Close -
func (s *Shader) Close() {
	gl.DeleteProgram(s.rendererID)
}

// Bind -
func (s *Shader) Bind() {
	gl.UseProgram(s.rendererID)
}

// Unbind -
func (s *Shader) Unbind() {
	gl.UseProgram(0)
}

// SetUniform4f -
func (s *Shader) SetUniform4f(name string, v0 float32, v1 float32, v2 float32, v3 float32) {
	gl.Uniform4f(s.getUniformLocation(name), v0, v1, v2, v3)
}

// SetUniformMat4f -
func (s *Shader) SetUniformMat4f(name string, matrix mgl32.Mat4) {
	gl.UniformMatrix4fv(s.getUniformLocation(name), 1, false, &matrix[0])
}

// GetUniformLocation -
func (s *Shader) getUniformLocation(name string) int32 {
	if loc, ok := s.uCache[name]; ok {
		return loc
	}
	loc := gl.GetUniformLocation(s.rendererID, gl.Str(name+"\x00"))
	s.uCache[name] = loc
	if loc == -1 {
		log.Printf("Warning: uniform %v doesn't exist\n", name)
	}
	return loc
}

func (s *Shader) parseShader() ShaderProgramSource {
	file, err := os.Open(s.filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var source ShaderProgramSource
	var mode string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#shader") {
			if strings.Contains(line, "vertex") {
				mode = "vertex"
			} else if strings.Contains(line, "fragment") {
				mode = "fragment"
			}
		} else {
			if mode == "vertex" {
				source.VertexShader = fmt.Sprintf("%v%v\n", source.VertexShader, line)
			} else if mode == "fragment" {
				source.FragmentShader = fmt.Sprintf("%v%v\n", source.FragmentShader, line)
			}
		}
	}

	source.VertexShader += "\x00"
	source.FragmentShader += "\x00"

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return source
}

func (s *Shader) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (s *Shader) createShader(source ShaderProgramSource) {
	vertexShader, err := s.compileShader(source.VertexShader, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := s.compileShader(source.FragmentShader, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	gl.ValidateProgram(prog)

	gl.DetachShader(prog, vertexShader)
	gl.DetachShader(prog, fragmentShader)

	s.rendererID = prog
}
