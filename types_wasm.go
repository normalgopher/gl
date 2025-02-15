//go:build js

package gl

import "syscall/js"

type Buffer struct {
	js.Value
}
type Framebuffer struct {
	js.Value
}
type Program struct {
	js.Value
}
type Renderbuffer struct {
	js.Value
}
type Sampler struct {
	js.Value
}
type Shader struct {
	js.Value
}
type Texture struct {
	js.Value
}
type Uniform struct {
	js.Value
}
type VertexArrayObject struct {
	js.Value
}

func valid(v js.Value) bool {
	return !(v.Equal(js.Undefined()) || v.Equal(js.Null()))
}

func (b *Buffer) Valid() bool            { return valid(b.Value) }
func (f *Framebuffer) Valid() bool       { return valid(f.Value) }
func (p *Program) Valid() bool           { return valid(p.Value) }
func (r *Renderbuffer) Valid() bool      { return valid(r.Value) }
func (s *Sampler) Valid() bool           { return valid(s.Value) }
func (s *Shader) Valid() bool            { return valid(s.Value) }
func (t *Texture) Valid() bool           { return valid(t.Value) }
func (u *Uniform) Valid() bool           { return valid(u.Value) }
func (v *VertexArrayObject) Valid() bool { return valid(v.Value) }

var (
	NoBuffer            Buffer            = Buffer{js.Null()}
	NoFramebuffer       Framebuffer       = Framebuffer{js.Null()}
	NoProgram           Program           = Program{js.Null()}
	NoRenderbuffer      Renderbuffer      = Renderbuffer{js.Null()}
	NoSampler           Sampler           = Sampler{js.Null()}
	NoShader            Shader            = Shader{js.Null()}
	NoTexture           Texture           = Texture{js.Null()}
	NoUniform           Uniform           = Uniform{js.Null()}
	NoVertexArrayObject VertexArrayObject = VertexArrayObject{js.Null()}
)
