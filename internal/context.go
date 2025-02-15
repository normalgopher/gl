//go:build js

package internal

import (
	"errors"
	"reflect"
	"strings"
	"syscall/js"
)

var (
	current *Context
)

func Current() *Context {
	if current == nil {
		panic("no current context")
	}
	return current
}

func SetCurrent(c *Context) {
	current = c
}

type Context struct {
	Context js.Value

	Funcs struct {
		CopyBufferSubData              js.Value
		GetBufferSubData               js.Value
		BlitFramebuffer                js.Value
		FramebufferTextureLayer        js.Value
		InvalidateFramebuffer          js.Value
		ReadBuffer                     js.Value
		RenderbufferStorageMultisample js.Value
		TexStorage2D                   js.Value
		TexStorage3D                   js.Value
		TexImage3D                     js.Value
		TexSubImage3D                  js.Value
		CopyTexSubImage3D              js.Value
		CompressedTexImage3D           js.Value
		CompressedTexSubImage3D        js.Value
		GetFragDataLocation            js.Value
		Uniform1ui                     js.Value
		Uniform2ui                     js.Value
		Uniform3ui                     js.Value
		Uniform4ui                     js.Value
		Uniform1uiv                    js.Value
		Uniform2uiv                    js.Value
		Uniform3uiv                    js.Value
		Uniform4uiv                    js.Value
		UniformMatrix3x2fv             js.Value
		UniformMatrix4x2fv             js.Value
		UniformMatrix2x3fv             js.Value
		UniformMatrix4x3fv             js.Value
		UniformMatrix2x4fv             js.Value
		UniformMatrix3x4fv             js.Value
		VertexAttribI4i                js.Value
		VertexAttribI4iv               js.Value
		VertexAttribI4ui               js.Value
		VertexAttribI4uiv              js.Value
		VertexAttribIPointer           js.Value
		VertexAttribDivisor            js.Value
		DrawArraysInstanced            js.Value
		DrawElementsInstanced          js.Value
		DrawRangeElements              js.Value
		DrawBuffers                    js.Value
		ClearBufferfv                  js.Value
		ClearBufferiv                  js.Value
		ClearBufferuiv                 js.Value
		ClearBufferfi                  js.Value
		CreateVertexArray              js.Value
		DeleteVertexArray              js.Value
		IsVertexArray                  js.Value
		BindVertexArray                js.Value
		ActiveTexture                  js.Value
		AttachShader                   js.Value
		BindAttribLocation             js.Value
		BindBuffer                     js.Value
		BindFramebuffer                js.Value
		BindRenderbuffer               js.Value
		BindTexture                    js.Value
		BlendColor                     js.Value
		BlendEquation                  js.Value
		BlendEquationSeparate          js.Value
		BlendFunc                      js.Value
		BlendFuncSeparate              js.Value
		CheckFramebufferStatus         js.Value
		Clear                          js.Value
		ClearColor                     js.Value
		ClearDepth                     js.Value
		ClearStencil                   js.Value
		ColorMask                      js.Value
		CompileShader                  js.Value
		CopyTexImage2D                 js.Value
		CopyTexSubImage2D              js.Value
		CreateBuffer                   js.Value
		CreateFramebuffer              js.Value
		CreateProgram                  js.Value
		CreateRenderbuffer             js.Value
		CreateShader                   js.Value
		CreateTexture                  js.Value
		CullFace                       js.Value
		DeleteBuffer                   js.Value
		DeleteFramebuffer              js.Value
		DeleteProgram                  js.Value
		DeleteRenderbuffer             js.Value
		DeleteShader                   js.Value
		DeleteTexture                  js.Value
		DepthFunc                      js.Value
		DepthMask                      js.Value
		DepthRange                     js.Value
		DetachShader                   js.Value
		Disable                        js.Value
		DisableVertexAttribArray       js.Value
		DrawArrays                     js.Value
		DrawElements                   js.Value
		Enable                         js.Value
		EnableVertexAttribArray        js.Value
		Finish                         js.Value
		Flush                          js.Value
		FramebufferRenderbuffer        js.Value
		FramebufferTexture2D           js.Value
		FrontFace                      js.Value
		GenerateMipmap                 js.Value
		GetError                       js.Value
		BufferData                     js.Value
		BufferSubData                  js.Value
		VertexAttribPointer            js.Value
		Viewport                       js.Value
		Scissor                        js.Value
		Uniform1i                      js.Value
		Uniform1f                      js.Value
		Uniform2fv                     js.Value
		Uniform3fv                     js.Value
		Uniform4fv                     js.Value
		UniformMatrix2fv               js.Value
		UniformMatrix3fv               js.Value
		UniformMatrix4fv               js.Value
		ShaderSource                   js.Value
		GetShaderParameter             js.Value
		GetShaderInfoLog               js.Value
		LinkProgram                    js.Value
		GetProgramParameter            js.Value
		GetProgramInfoLog              js.Value
		GetUniformLocation             js.Value
		UseProgram                     js.Value
		TexImage2D                     js.Value
		TexParameteri                  js.Value
		GetParameter                   js.Value
	}
}

func NewContext(canvas js.Value, attrs map[string]any) (*Context, error) {
	gl := canvas.Call("getContext", "webgl2", attrs)
	if gl.Equal(js.Null()) {
		return nil, errors.New("can't create webgl2 context")
	}

	c := &Context{
		Context: gl,
	}

	c.bindFunctions()

	return c, nil
}

func (c *Context) bindFunctions() {
	v := reflect.ValueOf(&c.Funcs).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		name := t.Field(i).Name
		name = strings.ToLower(name[:1]) + name[1:]
		f.Set(reflect.ValueOf(c.Context.Get(name).Call("bind", c.Context)))
	}
}
