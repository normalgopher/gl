//go:build !js

package gl

type Buffer uint32
type Framebuffer uint32
type Program uint32
type Renderbuffer uint32
type Sampler uint32
type Shader uint32
type Texture uint32
type Uniform int32
type VertexArrayObject uint32

func (b Buffer) Valid() bool            { return b != 0 }
func (f Framebuffer) Valid() bool       { return f != 0 }
func (p Program) Valid() bool           { return p != 0 }
func (r Renderbuffer) Valid() bool      { return r != 0 }
func (s Sampler) Valid() bool           { return s != 0 }
func (s Shader) Valid() bool            { return s != 0 }
func (t Texture) Valid() bool           { return t != 0 }
func (u Uniform) Valid() bool           { return u != 0 }
func (v VertexArrayObject) Valid() bool { return v != 0 }

var (
	NoBuffer            Buffer            = 0
	NoFramebuffer       Framebuffer       = 0
	NoProgram           Program           = 0
	NoRenderbuffer      Renderbuffer      = 0
	NoSampler           Sampler           = 0
	NoShader            Shader            = 0
	NoTexture           Texture           = 0
	NoUniform           Uniform           = 0
	NoVertexArrayObject VertexArrayObject = 0
)
