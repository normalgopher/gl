//go:build js

package gl

import (
	"fmt"
	"unsafe"

	"github.com/normalgopher/gl/internal"
)

var currentContext internal.Context

func Init() error {
	currentContext = internal.Context{}
	return nil
}

func CopyBufferSubData(readTarget, writeTarget Enum, readOffset, writeOffset int, size int) {
	internal.Current().Funcs.CopyBufferSubData.Invoke(readTarget.Value(), writeTarget.Value(), readOffset, writeOffset, size)
}

func GetBufferSubData(target Enum, offset int, size int, data unsafe.Pointer) {
	jsData := goPtrToJsArray(data, size)
	internal.Current().Funcs.GetBufferSubData.Invoke(target.Value(), offset, jsData)
}

func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter Enum) {
	internal.Current().Funcs.BlitFramebuffer.Invoke(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter.Value())
}

func FramebufferTextureLayer(target Enum, attachment Enum, texture uint32, level int32, layer int32) {
	internal.Current().Funcs.FramebufferTextureLayer.Invoke(target.Value(), attachment.Value(), texture, level, layer)
}

func InvalidateFramebuffer(target Enum, attachments []Enum) {
	jsData := goSliceToJsArray(attachments)
	internal.Current().Funcs.InvalidateFramebuffer.Invoke(target.Value(), jsData)
}

func ReadBuffer(src Enum) {
	internal.Current().Funcs.ReadBuffer.Invoke(src.Value())
}

func RenderbufferStorageMultisample(target Enum, samples int32, internalformat Enum, width int32, height int32) {
	internal.Current().Funcs.RenderbufferStorageMultisample.Invoke(target.Value(), samples, internalformat.Value(), width, height)
}

func TexStorage2D(target Enum, levels int32, internalformat Enum, width int32, height int32) {
	internal.Current().Funcs.TexStorage2D.Invoke(target.Value(), levels, internalformat.Value(), width, height)
}

func TexStorage3D(target Enum, levels int32, internalformat Enum, width int32, height int32, depth int32) {
	internal.Current().Funcs.TexStorage3D.Invoke(target.Value(), levels, internalformat.Value(), width, height, depth)
}

func TexImage3D[T SliceConstraints](target Enum, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format Enum, xtype Enum, pixels []T) {
	jsData := goSliceToJsArray(pixels)
	internal.Current().Funcs.TexImage3D.Invoke(target.Value(), level, internalformat, width, height, depth, border, format.Value(), xtype.Value(), jsData)
}

func TexSubImage3D[T SliceConstraints](target Enum, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format Enum, xtype Enum, pixels []T) {
	jsData := goSliceToJsArray(pixels)
	internal.Current().Funcs.TexSubImage3D.Invoke(target.Value(), level, xoffset, yoffset, zoffset, width, height, depth, format.Value(), xtype.Value(), jsData)
}

func CopyTexSubImage3D(target Enum, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	internal.Current().Funcs.CopyTexSubImage3D.Invoke(target.Value(), level, xoffset, yoffset, zoffset, x, y, width, height)
}

func CompressedTexImage3D(target Enum, level int32, internalformat Enum, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	jsData := goPtrToJsArray(data, int(imageSize))
	internal.Current().Funcs.CompressedTexImage3D.Invoke(target.Value(), level, internalformat.Value(), width, height, depth, border, jsData)
}

func CompressedTexSubImage3D(target Enum, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format Enum, imageSize int32, data unsafe.Pointer) {
	jsData := goPtrToJsArray(data, int(imageSize))
	internal.Current().Funcs.CompressedTexSubImage3D.Invoke(target.Value(), level, xoffset, yoffset, zoffset, width, height, depth, format.Value(), jsData)
}

func GetFragDataLocation(program Program, name string) int32 {
	return int32(internal.Current().Funcs.GetFragDataLocation.Invoke(program.Value, name).Int())
}

func Uniform1ui(location Uniform, v0 uint32) {
	internal.Current().Funcs.Uniform1ui.Invoke(location.Value, v0)
}

func Uniform2ui(location Uniform, v0, v1 uint32) {
	internal.Current().Funcs.Uniform2ui.Invoke(location.Value, v0, v1)
}

func Uniform3ui(location Uniform, v0, v1, v2 uint32) {
	internal.Current().Funcs.Uniform3ui.Invoke(location.Value, v0, v1, v2)
}

func Uniform4ui(location Uniform, v0, v1, v2, v3 uint32) {
	internal.Current().Funcs.Uniform4ui.Invoke(location.Value, v0, v1, v2, v3)
}

func UniformNui(location Uniform, v ...uint32) {
	switch len(v) {
	case 1:
		Uniform1ui(location, v[0])
	case 2:
		Uniform2ui(location, v[0], v[1])
	case 3:
		Uniform3ui(location, v[0], v[1], v[2])
	case 4:
		Uniform4ui(location, v[0], v[1], v[2], v[3])
	default:
		panic(fmt.Sprintf("invalid uniform value count: %v", len(v)))
	}
}

func Uniform1uiv(location Uniform, value [1]uint32) {
	jsData := goSliceToJsArray(value[:])
	internal.Current().Funcs.Uniform1uiv.Invoke(location.Value, jsData)
}

func Uniform2uiv(location Uniform, value [2]uint32) {
	jsData := goSliceToJsArray(value[:])
	internal.Current().Funcs.Uniform2uiv.Invoke(location.Value, jsData)
}

func Uniform3uiv(location Uniform, value [3]uint32) {
	jsData := goSliceToJsArray(value[:])
	internal.Current().Funcs.Uniform3uiv.Invoke(location.Value, jsData)
}

func Uniform4uiv(location Uniform, value [4]uint32) {
	jsData := goSliceToJsArray(value[:])
	internal.Current().Funcs.Uniform4uiv.Invoke(location.Value, jsData)
}

func UniformMatrix3x2fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix3x2fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix4x2fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix4x2fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix2x3fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix2x3fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix4x3fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix4x3fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix2x4fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix2x4fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix3x4fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix3x4fv.Invoke(location.Value, transpose, jsData)
}

func VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32) {
	internal.Current().Funcs.VertexAttribI4i.Invoke(index, x, y, z, w)
}

func VertexAttribI4iv(index uint32, v [4]int32) {
	jsData := goSliceToJsArray(v[:])
	internal.Current().Funcs.VertexAttribI4iv.Invoke(index, jsData)
}

func VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32) {
	internal.Current().Funcs.VertexAttribI4ui.Invoke(index, x, y, z, w)
}

func VertexAttribI4uiv(index uint32, v [4]uint32) {
	jsData := goSliceToJsArray(v[:])
	internal.Current().Funcs.VertexAttribI4uiv.Invoke(index, jsData)
}

func VertexAttribIPointer(index uint32, size int32, xtype Enum, stride int32, pointer unsafe.Pointer) {
	internal.Current().Funcs.VertexAttribIPointer.Invoke(index, size, xtype.Value(), stride, pointer)
}

func VertexAttribDivisor(index uint32, divisor uint32) {
	internal.Current().Funcs.VertexAttribDivisor.Invoke(index, divisor)
}

func DrawArraysInstanced(mode Enum, first int32, count int32, instancecount int32) {
	internal.Current().Funcs.DrawArraysInstanced.Invoke(mode.Value(), first, count, instancecount)
}

func DrawElementsInstanced(mode Enum, count int32, xtype Enum, indices unsafe.Pointer, instancecount int32) {
	internal.Current().Funcs.DrawElementsInstanced.Invoke(mode.Value(), count, xtype.Value(), indices, instancecount)
}

func DrawRangeElements(mode Enum, start uint32, end uint32, count int32, xtype Enum, indices unsafe.Pointer) {
	internal.Current().Funcs.DrawRangeElements.Invoke(mode.Value(), start, end, count, xtype.Value(), indices)
}

func DrawBuffers(bufs []Enum) {
	jsData := goSliceToJsArray(bufs)
	internal.Current().Funcs.DrawBuffers.Invoke(jsData)
}

func ClearBufferfv(buffer Enum, drawbuffer int32, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.ClearBufferfv.Invoke(buffer.Value(), drawbuffer, jsData)
}

func ClearBufferiv(buffer Enum, drawbuffer int32, value []int32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.ClearBufferiv.Invoke(buffer.Value(), drawbuffer, jsData)
}

func ClearBufferuiv(buffer Enum, drawbuffer int32, value []uint32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.ClearBufferuiv.Invoke(buffer.Value(), drawbuffer, jsData)
}

func ClearBufferfi(buffer Enum, drawbuffer int32, depth float32, stencil int32) {
	internal.Current().Funcs.ClearBufferfi.Invoke(buffer.Value(), drawbuffer, depth, stencil)
}

func CreateVertexArray() VertexArrayObject {
	return VertexArrayObject{internal.Current().Funcs.CreateVertexArray.Invoke()}
}

func DeleteVertexArray(array VertexArrayObject) {
	internal.Current().Funcs.DeleteVertexArray.Invoke(array.Value)
}

func IsVertexArray(array VertexArrayObject) bool {
	return internal.Current().Funcs.IsVertexArray.Invoke(array.Value).Bool()
}

func BindVertexArray(array VertexArrayObject) {
	internal.Current().Funcs.BindVertexArray.Invoke(array.Value)
}

func ActiveTexture(texture Enum) {
	internal.Current().Funcs.ActiveTexture.Invoke(texture.Value())
}

func AttachShader(program Program, shader Shader) {
	internal.Current().Funcs.AttachShader.Invoke(program.Value, shader.Value)
}

func BindAttribLocation(program Program, index uint32, name string) {
	internal.Current().Funcs.BindAttribLocation.Invoke(program.Value, index, name)
}

func BindBuffer(target Enum, buffer Buffer) {
	internal.Current().Funcs.BindBuffer.Invoke(target.Value(), buffer.Value)
}

func BindFramebuffer(target Enum, framebuffer Framebuffer) {
	internal.Current().Funcs.BindFramebuffer.Invoke(target.Value(), framebuffer.Value)
}

func BindRenderbuffer(target Enum, renderbuffer Renderbuffer) {
	internal.Current().Funcs.BindRenderbuffer.Invoke(target.Value(), renderbuffer.Value)
}

func BindTexture(target Enum, texture Texture) {
	internal.Current().Funcs.BindTexture.Invoke(target.Value(), texture.Value)
}

func BlendColor(red, green, blue, alpha Clampf) {
	internal.Current().Funcs.BlendColor.Invoke(red.Value(), green.Value(), blue.Value(), alpha.Value())
}

func BlendEquation(mode Enum) {
	internal.Current().Funcs.BlendEquation.Invoke(mode.Value())
}

func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	internal.Current().Funcs.BlendEquationSeparate.Invoke(modeRGB.Value(), modeAlpha.Value())
}

func BlendFunc(sfactor, dfactor Enum) {
	internal.Current().Funcs.BlendFunc.Invoke(sfactor.Value(), dfactor.Value())
}

func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	internal.Current().Funcs.BlendFuncSeparate.Invoke(sfactorRGB.Value(), dfactorRGB.Value(), sfactorAlpha.Value(), dfactorAlpha.Value())
}

func CheckFramebufferStatus(target Enum) {
	internal.Current().Funcs.CheckFramebufferStatus.Invoke(target.Value())
}

func Clear(mask uint32) {
	internal.Current().Funcs.Clear.Invoke(mask)
}

func ClearColor(red, green, blue, alpha Clampf) {
	internal.Current().Funcs.ClearColor.Invoke(red.Value(), green.Value(), blue.Value(), alpha.Value())
}

func ClearDepth(depth Clampf) {
	internal.Current().Funcs.ClearDepth.Invoke(depth.Value())
}

func ClearStencil(s int32) {
	internal.Current().Funcs.ClearStencil.Invoke(s)
}

func ColorMask(red, green, blue, alpha bool) {
	internal.Current().Funcs.ColorMask.Invoke(red, green, blue, alpha)
}

func CompileShader(shader Shader) {
	internal.Current().Funcs.CompileShader.Invoke(shader.Value)
}

func CopyTexImage2D(target Enum, level int32, internalformat Enum, x int32, y int32, width int32, height int32, border int32) {
	internal.Current().Funcs.CopyTexImage2D.Invoke(target.Value(), level, internalformat.Value(), x, y, width, height, border)
}

func CopyTexSubImage2D(target Enum, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	internal.Current().Funcs.CopyTexSubImage2D.Invoke(target.Value(), level, xoffset, yoffset, x, y, width, height)
}

func CreateBuffer() Buffer {
	return Buffer{internal.Current().Funcs.CreateBuffer.Invoke()}
}

func CreateFramebuffer() Framebuffer {
	return Framebuffer{internal.Current().Funcs.CreateFramebuffer.Invoke()}
}

func CreateProgram() Program {
	return Program{internal.Current().Funcs.CreateProgram.Invoke()}
}

func CreateRenderbuffer() Renderbuffer {
	return Renderbuffer{internal.Current().Funcs.CreateRenderbuffer.Invoke()}
}

func CreateShader(xtype Enum) Shader {
	return Shader{internal.Current().Funcs.CreateShader.Invoke(xtype.Value())}
}

func CreateTexture() Texture {
	return Texture{internal.Current().Funcs.CreateTexture.Invoke()}
}

func CullFace(mode Enum) {
	internal.Current().Funcs.CullFace.Invoke(mode.Value())
}

func DeleteBuffer(buffer Buffer) {
	internal.Current().Funcs.DeleteBuffer.Invoke(buffer.Value)
}

func DeleteFramebuffer(buffer Framebuffer) {
	internal.Current().Funcs.DeleteFramebuffer.Invoke(buffer.Value)
}

func DeleteProgram(program Program) {
	internal.Current().Funcs.DeleteProgram.Invoke(program.Value)
}

func DeleteRenderbuffer(buffer Renderbuffer) {
	internal.Current().Funcs.DeleteRenderbuffer.Invoke(buffer.Value)
}

func DeleteShader(shader Shader) {
	internal.Current().Funcs.DeleteShader.Invoke(shader.Value)
}

func DeleteTexture(texture Texture) {
	internal.Current().Funcs.DeleteTexture.Invoke(texture.Value)
}

func DepthFunc(xfunc Enum) {
	internal.Current().Funcs.DepthFunc.Invoke(xfunc.Value())
}

func DepthMask(flag bool) {
	internal.Current().Funcs.DepthMask.Invoke(flag)
}

func DepthRange(near, far Clampf) {
	internal.Current().Funcs.DepthRange.Invoke(near.Value(), far.Value())
}

func DetachShader(program Program, shader Shader) {
	internal.Current().Funcs.DetachShader.Invoke(program.Value, shader.Value)
}

func Disable(cap Enum) {
	internal.Current().Funcs.Disable.Invoke(cap.Value())
}

func DisableVertexAttribArray(index uint32) {
	internal.Current().Funcs.DisableVertexAttribArray.Invoke(index)
}

func DrawArrays(mode Enum, first int32, count int32) {
	internal.Current().Funcs.DrawArrays.Invoke(mode.Value(), first, count)
}

func DrawElements[T IndexType](mode Enum, xtype Enum, indices []T) {
	jsData := goSliceToJsArray(indices)
	internal.Current().Funcs.DrawElements.Invoke(mode.Value(), len(indices), xtype.Value(), jsData)
}

func Enable(cap Enum) {
	internal.Current().Funcs.Enable.Invoke(cap.Value())
}

func EnableVertexAttribArray(index uint32) {
	internal.Current().Funcs.EnableVertexAttribArray.Invoke(index)
}

func Finish() {
	internal.Current().Funcs.Finish.Invoke()
}

func Flush() {
	internal.Current().Funcs.Flush.Invoke()
}

func FramebufferRenderbuffer(target, attachment, renderbuffertarget Enum, renderbuffer Renderbuffer) {
	internal.Current().Funcs.FramebufferRenderbuffer.Invoke(target.Value(), attachment.Value(), renderbuffertarget.Value(), renderbuffer.Value)
}

func FramebufferTexture2D(target, attachment, textarget Enum, texture Texture, level int32) {
	internal.Current().Funcs.FramebufferTexture2D.Invoke(target.Value(), attachment.Value(), textarget.Value(), texture.Value, level)
}

func FrontFace(mode Enum) {
	internal.Current().Funcs.FrontFace.Invoke(mode.Value())
}

func GenerateMipmap(target Enum) {
	internal.Current().Funcs.GenerateMipmap.Invoke(target.Value())
}

func GetError() Enum {
	return Enum(internal.Current().Funcs.GetError.Invoke().Int())
}

func BufferData(target Enum, size int, data unsafe.Pointer, usage Enum) {
	jsData := goPtrToJsArray(data, size)
	if jsData.IsNull() {
		internal.Current().Funcs.BufferData.Invoke(target.Value(), size, usage.Value())
	} else {
		internal.Current().Funcs.BufferData.Invoke(target.Value(), jsData, usage.Value())
	}
}

func BufferSubData(target Enum, offset int, size int, data unsafe.Pointer) {
	jsData := goPtrToJsArray(data, size)
	internal.Current().Funcs.BufferSubData.Invoke(target.Value(), offset, jsData)
}

func VertexAttribPointer(index uint32, size int32, xtype Enum, normalized bool, stride int32, offset uintptr) {
	internal.Current().Funcs.VertexAttribPointer.Invoke(index, size, xtype.Value(), normalized, stride, offset)
}

func Viewport(x int32, y int32, width int32, height int32) {
	internal.Current().Funcs.Viewport.Invoke(x, y, width, height)
}

func Scissor(x int32, y int32, width int32, height int32) {
	internal.Current().Funcs.Scissor.Invoke(x, y, width, height)
}

func Uniform1i(location Uniform, v0 int32) {
	internal.Current().Funcs.Uniform1i.Invoke(location.Value, v0)
}

func Uniform1f(location Uniform, v0 float32) {
	internal.Current().Funcs.Uniform1f.Invoke(location.Value, v0)
}

func Uniform2fv(location Uniform, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.Uniform2fv.Invoke(location.Value, jsData)
}

func Uniform3fv(location Uniform, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.Uniform3fv.Invoke(location.Value, jsData)
}

func Uniform4fv(location Uniform, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.Uniform4fv.Invoke(location.Value, jsData)
}

func UniformMatrix2fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix2fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix3fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix3fv.Invoke(location.Value, transpose, jsData)
}

func UniformMatrix4fv(location Uniform, transpose bool, value []float32) {
	jsData := goSliceToJsArray(value)
	internal.Current().Funcs.UniformMatrix4fv.Invoke(location.Value, transpose, jsData)
}

func ShaderSource(shader Shader, source string) {
	internal.Current().Funcs.ShaderSource.Invoke(shader.Value, source)
}

func GetShaderParameter(shader Shader, pname Enum) any {
	v := internal.Current().Funcs.GetShaderParameter.Invoke(shader.Value, pname.Value())
	switch pname {
	case SHADER_TYPE:
		return Enum(v.Int())
	case DELETE_STATUS, COMPILE_STATUS:
		return v.Bool()
	default:
		panic(fmt.Sprintf("invalid shader parameter: %v", pname))
	}
}

func GetShaderInfoLog(shader Shader) string {
	return internal.Current().Funcs.GetShaderInfoLog.Invoke(shader.Value).String()
}

func LinkProgram(program Program) {
	internal.Current().Funcs.LinkProgram.Invoke(program.Value)
}

func GetProgramParameter(program Program, pname Enum) any {
	v := internal.Current().Funcs.GetProgramParameter.Invoke(program.Value, pname.Value())
	switch pname {
	case DELETE_STATUS, LINK_STATUS, VALIDATE_STATUS:
		return v.Bool()
	case ATTACHED_SHADERS, ACTIVE_ATTRIBUTES, ACTIVE_UNIFORMS, TRANSFORM_FEEDBACK_VARYINGS, ACTIVE_UNIFORM_BLOCKS:
		return int32(v.Int())
	case TRANSFORM_FEEDBACK_BUFFER_MODE:
		return Enum(v.Int())
	default:
		panic(fmt.Sprintf("invalid program parameter: %v", pname))
	}
}

func GetProgramInfoLog(program Program) string {
	return internal.Current().Funcs.GetProgramInfoLog.Invoke(program.Value).String()
}

func GetUniformLocation(program Program, name string) Uniform {
	return Uniform{internal.Current().Funcs.GetUniformLocation.Invoke(program.Value, name)}
}

func UseProgram(program Program) {
	internal.Current().Funcs.UseProgram.Invoke(program.Value)
}

func TexImage2D[T SliceConstraints](target Enum, level int32, internalformat Enum, width int32, height int32, border int32, format Enum, xtype Enum, pixels []T) {
	pixData := goSliceToJsArray(pixels)
	internal.Current().Funcs.TexImage2D.Invoke(target.Value(), level, internalformat.Value(), width, height, border, format.Value(), xtype.Value(), pixData)
}

func TexParameteri(target Enum, pname Enum, param int32) {
	internal.Current().Funcs.TexParameteri.Invoke(target.Value(), pname.Value(), param)
}

func GetParameter(pname Enum) any {
	val := internal.Current().Funcs.GetParameter.Invoke(pname.Value())

	switch pname {
	case ACTIVE_TEXTURE, BLEND_DST_ALPHA, BLEND_DST_RGB, BLEND_EQUATION_ALPHA, BLEND_EQUATION_RGB, BLEND_SRC_ALPHA, BLEND_SRC_RGB, CULL_FACE_MODE, DEPTH_FUNC, FRONT_FACE, GENERATE_MIPMAP_HINT, IMPLEMENTATION_COLOR_READ_FORMAT, IMPLEMENTATION_COLOR_READ_TYPE, STENCIL_BACK_FAIL, STENCIL_BACK_FUNC, STENCIL_BACK_PASS_DEPTH_FAIL, STENCIL_BACK_PASS_DEPTH_PASS, STENCIL_FAIL, STENCIL_FUNC, STENCIL_PASS_DEPTH_FAIL, STENCIL_PASS_DEPTH_PASS, DRAW_BUFFER0, DRAW_BUFFER1, DRAW_BUFFER2, DRAW_BUFFER3, DRAW_BUFFER4, DRAW_BUFFER5, DRAW_BUFFER6, DRAW_BUFFER7, DRAW_BUFFER8, DRAW_BUFFER9, DRAW_BUFFER10, DRAW_BUFFER11, DRAW_BUFFER12, DRAW_BUFFER13, DRAW_BUFFER14, DRAW_BUFFER15, FRAGMENT_SHADER_DERIVATIVE_HINT, READ_BUFFER:
		return Enum(val.Int())
	case ALIASED_LINE_WIDTH_RANGE, ALIASED_POINT_SIZE_RANGE, DEPTH_RANGE:
		v := make([]float32, 2)
		for i := range val.Length() {
			v[i] = float32(val.Index(i).Float())
		}
		return v
	case ALPHA_BITS, GREEN_BITS, MAX_COMBINED_TEXTURE_IMAGE_UNITS, MAX_CUBE_MAP_TEXTURE_SIZE, MAX_FRAGMENT_UNIFORM_VECTORS, MAX_RENDERBUFFER_SIZE, MAX_TEXTURE_IMAGE_UNITS, MAX_TEXTURE_SIZE, MAX_VARYING_VECTORS, MAX_VERTEX_ATTRIBS, MAX_VERTEX_TEXTURE_IMAGE_UNITS, MAX_VERTEX_UNIFORM_VECTORS, PACK_ALIGNMENT, RED_BITS, SAMPLE_BUFFERS, SAMPLES, STENCIL_BACK_REF, STENCIL_BITS, STENCIL_CLEAR_VALUE, STENCIL_REF, SUBPIXEL_BITS, UNPACK_ALIGNMENT, MAX_3D_TEXTURE_SIZE, MAX_ARRAY_TEXTURE_LAYERS, MAX_COLOR_ATTACHMENTS, MAX_COMBINED_UNIFORM_BLOCKS, MAX_DRAW_BUFFERS, MAX_ELEMENTS_INDICES, MAX_ELEMENTS_VERTICES, MAX_FRAGMENT_INPUT_COMPONENTS, MAX_FRAGMENT_UNIFORM_BLOCKS, MAX_FRAGMENT_UNIFORM_COMPONENTS, MAX_PROGRAM_TEXEL_OFFSET, MAX_SAMPLES, MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, MAX_UNIFORM_BUFFER_BINDINGS, MAX_VARYING_COMPONENTS, MAX_VERTEX_OUTPUT_COMPONENTS, MAX_VERTEX_UNIFORM_BLOCKS, MAX_VERTEX_UNIFORM_COMPONENTS, MIN_PROGRAM_TEXEL_OFFSET, PACK_ROW_LENGTH, PACK_SKIP_PIXELS, PACK_SKIP_ROWS, UNIFORM_BUFFER_OFFSET_ALIGNMENT, UNPACK_IMAGE_HEIGHT, UNPACK_ROW_LENGTH, UNPACK_SKIP_IMAGES, UNPACK_SKIP_PIXELS, UNPACK_SKIP_ROWS:
		return int32(val.Int())
	case ARRAY_BUFFER_BINDING:
		return Buffer{val}
	case BLEND, CULL_FACE, DEPTH_TEST, DEPTH_WRITEMASK, DITHER, POLYGON_OFFSET_FILL, SAMPLE_ALPHA_TO_COVERAGE, SAMPLE_COVERAGE, SAMPLE_COVERAGE_INVERT, SCISSOR_TEST, STENCIL_TEST, RASTERIZER_DISCARD, TRANSFORM_FEEDBACK_ACTIVE, TRANSFORM_FEEDBACK_PAUSED:
		return val.Bool()
	case BLEND_COLOR, COLOR_CLEAR_VALUE:
		v := make([]float32, 4)
		for i := range val.Length() {
			v[i] = float32(val.Index(i).Float())
		}
		return v
	case BLUE_BITS, DEPTH_BITS:
		return int32(val.Int())
	case COLOR_WRITEMASK:
		v := make([]bool, 4)
		for i := range val.Length() {
			v[i] = val.Index(i).Bool()
		}
		return v
	case COMPRESSED_TEXTURE_FORMATS:
		v := make([]int32, val.Length())
		for i := range val.Length() {
			v[i] = int32(val.Index(i).Int())
		}
		return v
	case CURRENT_PROGRAM:
		return Program{val}
	case DEPTH_CLEAR_VALUE, LINE_WIDTH, POLYGON_OFFSET_FACTOR, POLYGON_OFFSET_UNITS, SAMPLE_COVERAGE_VALUE, MAX_TEXTURE_LOD_BIAS:
		return float32(val.Float())
	case ELEMENT_ARRAY_BUFFER_BINDING, COPY_READ_BUFFER_BINDING, COPY_WRITE_BUFFER_BINDING, PIXEL_PACK_BUFFER_BINDING, PIXEL_UNPACK_BUFFER_BINDING, TRANSFORM_FEEDBACK_BUFFER_BINDING, UNIFORM_BUFFER_BINDING:
		return Buffer{val}
	case FRAMEBUFFER_BINDING, READ_FRAMEBUFFER_BINDING:
		return Framebuffer{val}
	case STENCIL_BACK_VALUE_MASK, STENCIL_BACK_WRITEMASK, STENCIL_VALUE_MASK, STENCIL_WRITEMASK:
		return uint32(val.Int())
	case MAX_VIEWPORT_DIMS:
		v := make([]int32, 2)
		for i := range val.Length() {
			v[i] = int32(val.Index(i).Int())
		}
		return v
	case RENDERBUFFER_BINDING:
		return Renderbuffer{val}
	case SCISSOR_BOX, VIEWPORT:
		v := make([]int32, 4)
		for i := range val.Length() {
			v[i] = int32(val.Index(i).Int())
		}
		return v
	case TEXTURE_BINDING_2D, TEXTURE_BINDING_CUBE_MAP, TEXTURE_BINDING_2D_ARRAY, TEXTURE_BINDING_3D:
		return Texture{val}
	case VERTEX_ARRAY_BINDING:
		return VertexArrayObject{val}
	case MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, MAX_ELEMENT_INDEX, MAX_SERVER_WAIT_TIMEOUT, MAX_UNIFORM_BLOCK_SIZE:
		return int64(val.Int())
	case SAMPLER_BINDING:
		return Sampler{val}
	default:
		panic(fmt.Sprintf("invalid parameter name: %v", pname))
	}
}
