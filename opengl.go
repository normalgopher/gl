//go:build !js

package gl

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
)

func CopyBufferSubData(readTarget, writeTarget Enum, readOffset, writeOffset int, size int) {
	gl.CopyBufferSubData(uint32(readTarget), uint32(writeTarget), readOffset, writeOffset, size)
}

func GetBufferSubData(target Enum, offset int, size int, data unsafe.Pointer) {
	gl.GetBufferSubData(uint32(target), offset, size, data)
}

func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter Enum) {
	gl.BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, uint32(filter))
}

func FramebufferTextureLayer(target Enum, attachment Enum, texture uint32, level int32, layer int32) {
	gl.FramebufferTextureLayer(uint32(target), uint32(attachment), texture, level, layer)
}

func InvalidateFramebuffer(target Enum, attachments []Enum) {
	gl.InvalidateFramebuffer(uint32(target), int32(len(attachments)), (*uint32)(unsafe.SliceData(attachments)))
}

func ReadBuffer(src Enum) {
	gl.ReadBuffer(uint32(src))
}

func RenderbufferStorageMultisample(target Enum, samples int32, internalformat Enum, width int32, height int32) {
	gl.RenderbufferStorageMultisample(uint32(target), samples, uint32(internalformat), width, height)
}

func TexStorage2D(target Enum, levels int32, internalformat Enum, width int32, height int32) {
	gl.TexStorage2D(uint32(target), levels, uint32(internalformat), width, height)
}

func TexStorage3D(target Enum, levels int32, internalformat Enum, width int32, height int32, depth int32) {
	gl.TexStorage3D(uint32(target), levels, uint32(internalformat), width, height, depth)
}

func TexImage3D[T SliceConstraints](target Enum, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format Enum, xtype Enum, pixels []T) {
	gl.TexImage3D(uint32(target), level, internalformat, width, height, depth, border, uint32(format), uint32(xtype), gl.Ptr(pixels))
}

func TexSubImage3D[T SliceConstraints](target Enum, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format Enum, xtype Enum, pixels []T) {
	gl.TexSubImage3D(uint32(target), level, xoffset, yoffset, zoffset, width, height, depth, uint32(format), uint32(xtype), gl.Ptr(pixels))
}

func CopyTexSubImage3D(target Enum, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	gl.CopyTexSubImage3D(uint32(target), level, xoffset, yoffset, zoffset, x, y, width, height)
}

func CompressedTexImage3D(target Enum, level int32, internalformat Enum, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	gl.CompressedTexImage3D(uint32(target), level, uint32(internalformat), width, height, depth, border, imageSize, data)
}

func CompressedTexSubImage3D(target Enum, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format Enum, imageSize int32, data unsafe.Pointer) {
	gl.CompressedTexSubImage3D(uint32(target), level, xoffset, yoffset, zoffset, width, height, depth, uint32(format), imageSize, data)
}

func GetFragDataLocation(program Program, name string) int32 {
	return gl.GetFragDataLocation(uint32(program), unsafe.StringData(name))
}

func Uniform1ui(location Uniform, v0 uint32) {
	gl.Uniform1ui(int32(location), v0)
}

func Uniform2ui(location Uniform, v0, v1 uint32) {
	gl.Uniform2ui(int32(location), v0, v1)
}

func Uniform3ui(location Uniform, v0, v1, v2 uint32) {
	gl.Uniform3ui(int32(location), v0, v1, v2)
}

func Uniform4ui(location Uniform, v0, v1, v2, v3 uint32) {
	gl.Uniform4ui(int32(location), v0, v1, v2, v3)
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
	gl.Uniform1uiv(int32(location), int32(len(value)), unsafe.SliceData(value[:]))
}

func Uniform2uiv(location Uniform, value [2]uint32) {
	gl.Uniform2uiv(int32(location), int32(len(value)), unsafe.SliceData(value[:]))
}

func Uniform3uiv(location Uniform, value [3]uint32) {
	gl.Uniform3uiv(int32(location), int32(len(value)), unsafe.SliceData(value[:]))
}

func Uniform4uiv(location Uniform, value [4]uint32) {
	gl.Uniform4uiv(int32(location), int32(len(value)), unsafe.SliceData(value[:]))
}

func UniformMatrix3x2fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix3x2fv(int32(location), int32(len(value)), transpose, unsafe.SliceData(value))
}

func UniformMatrix4x2fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix4x2fv(int32(location), int32(len(value)), transpose, unsafe.SliceData(value))
}

func UniformMatrix2x3fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix2x3fv(int32(location), int32(len(value)), transpose, unsafe.SliceData(value))
}

func UniformMatrix4x3fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix4x3fv(int32(location), int32(len(value)), transpose, unsafe.SliceData(value))
}

func UniformMatrix2x4fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix2x4fv(int32(location), int32(len(value)), transpose, unsafe.SliceData(value))
}

func UniformMatrix3x4fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix3x4fv(int32(location), int32(len(value)), transpose, unsafe.SliceData(value))
}

func VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32) {
	gl.VertexAttribI4i(index, x, y, z, w)
}

func VertexAttribI4iv(index uint32, v [4]int32) {
	gl.VertexAttribI4iv(index, unsafe.SliceData(v[:]))
}

func VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32) {
	gl.VertexAttribI4ui(index, x, y, z, w)
}

func VertexAttribI4uiv(index uint32, v [4]uint32) {
	gl.VertexAttribI4uiv(index, unsafe.SliceData(v[:]))
}

func VertexAttribIPointer(index uint32, size int32, xtype Enum, stride int32, pointer unsafe.Pointer) {
	gl.VertexAttribIPointer(index, size, uint32(xtype), stride, pointer)
}

func VertexAttribDivisor(index uint32, divisor uint32) {
	gl.VertexAttribDivisor(index, divisor)
}

func DrawArraysInstanced(mode Enum, first int32, count int32, instancecount int32) {
	gl.DrawArraysInstanced(uint32(mode), first, count, instancecount)
}

func DrawElementsInstanced(mode Enum, count int32, xtype Enum, indices unsafe.Pointer, instancecount int32) {
	gl.DrawElementsInstanced(uint32(mode), count, uint32(xtype), indices, instancecount)
}

func DrawRangeElements(mode Enum, start uint32, end uint32, count int32, xtype Enum, indices unsafe.Pointer) {
	gl.DrawRangeElements(uint32(mode), start, end, count, uint32(xtype), indices)
}

func DrawBuffers(bufs []Enum) {
	gl.DrawBuffers(int32(len(bufs)), (*uint32)(unsafe.SliceData(bufs)))
}

func ClearBufferfv(buffer Enum, drawbuffer int32, value []float32) {
	gl.ClearBufferfv(uint32(buffer), drawbuffer, unsafe.SliceData(value))
}

func ClearBufferiv(buffer Enum, drawbuffer int32, value []int32) {
	gl.ClearBufferiv(uint32(buffer), drawbuffer, unsafe.SliceData(value))
}

func ClearBufferuiv(buffer Enum, drawbuffer int32, value []uint32) {
	gl.ClearBufferuiv(uint32(buffer), drawbuffer, unsafe.SliceData(value))
}

func ClearBufferfi(buffer Enum, drawbuffer int32, depth float32, stencil int32) {
	gl.ClearBufferfi(uint32(buffer), drawbuffer, depth, stencil)
}

func CreateVertexArray() VertexArrayObject {
	var v uint32
	gl.CreateVertexArrays(1, &v)
	return VertexArrayObject(v)
}

func DeleteVertexArray(array VertexArrayObject) {
	gl.DeleteVertexArrays(1, (*uint32)(&array))
}

func IsVertexArray(array VertexArrayObject) bool {
	return gl.IsVertexArray(uint32(array))
}

func BindVertexArray(array VertexArrayObject) {
	gl.BindVertexArray(uint32(array))
}

func ActiveTexture(texture Enum) {
	gl.ActiveTexture(uint32(texture))
}

func AttachShader(program Program, shader Shader) {
	gl.AttachShader(uint32(program), uint32(shader))
}

func BindAttribLocation(program Program, index uint32, name string) {
	gl.BindAttribLocation(uint32(program), index, unsafe.StringData(name))
}

func BindBuffer(target Enum, buffer Buffer) {
	gl.BindBuffer(uint32(target), uint32(buffer))
}

func BindFramebuffer(target Enum, framebuffer Framebuffer) {
	gl.BindFramebuffer(uint32(target), uint32(framebuffer))
}

func BindRenderbuffer(target Enum, renderbuffer Renderbuffer) {
	gl.BindRenderbuffer(uint32(target), uint32(renderbuffer))
}

func BindTexture(target Enum, texture Texture) {
	gl.BindTexture(uint32(target), uint32(texture))
}

func BlendColor(red, green, blue, alpha Clampf) {
	gl.BlendColor(float32(red), float32(green), float32(blue), float32(alpha))
}

func BlendEquation(mode Enum) {
	gl.BlendEquation(uint32(mode))
}

func BlendEquationSeparate(modeRGB, modeAlpha Enum) {
	gl.BlendEquationSeparate(uint32(modeRGB), uint32(modeAlpha))
}

func BlendFunc(sfactor, dfactor Enum) {
	gl.BlendFunc(uint32(sfactor), uint32(dfactor))
}

func BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha Enum) {
	gl.BlendFuncSeparate(uint32(sfactorRGB), uint32(dfactorRGB), uint32(sfactorAlpha), uint32(dfactorAlpha))
}

func CheckFramebufferStatus(target Enum) {
	gl.CheckFramebufferStatus(uint32(target))
}

func Clear(mask uint32) {
	gl.Clear(mask)
}

func ClearColor(red, green, blue, alpha Clampf) {
	gl.ClearColor(float32(red), float32(green), float32(blue), float32(alpha))
}

func ClearDepth(depth Clampf) {
	gl.ClearDepth(float64(depth))
}

func ClearStencil(s int32) {
	gl.ClearStencil(s)
}

func ColorMask(red, green, blue, alpha bool) {
	gl.ColorMask(red, green, blue, alpha)
}

func CompileShader(shader Shader) {
	gl.CompileShader(uint32(shader))
}

func CopyTexImage2D(target Enum, level int32, internalformat Enum, x int32, y int32, width int32, height int32, border int32) {
	gl.CopyTexImage2D(uint32(target), level, uint32(internalformat), x, y, width, height, border)
}

func CopyTexSubImage2D(target Enum, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	gl.CopyTexSubImage2D(uint32(target), level, xoffset, yoffset, x, y, width, height)
}

func CreateBuffer() Buffer {
	var b uint32
	gl.CreateBuffers(1, &b)
	return Buffer(b)
}

func CreateFramebuffer() Framebuffer {
	var v uint32
	gl.CreateFramebuffers(1, &v)
	return Framebuffer(v)
}

func CreateProgram() Program {
	return Program(gl.CreateProgram())
}

func CreateRenderbuffer() Renderbuffer {
	var v uint32
	gl.CreateRenderbuffers(1, &v)
	return Renderbuffer(v)
}

func CreateShader(xtype Enum) Shader {
	return Shader(gl.CreateShader(uint32(xtype)))
}

func CreateTexture() Texture {
	var v uint32
	gl.GenTextures(1, &v)
	return Texture(v)
}

func CullFace(mode Enum) {
	gl.CullFace(uint32(mode))
}

func DeleteBuffer(buffer Buffer) {
	gl.DeleteBuffers(1, (*uint32)(&buffer))
}

func DeleteFramebuffer(buffer Framebuffer) {
	gl.DeleteFramebuffers(1, (*uint32)(&buffer))
}

func DeleteProgram(program Program) {
	gl.DeleteProgram(uint32(program))
}

func DeleteRenderbuffer(buffer Renderbuffer) {
	gl.DeleteRenderbuffers(1, (*uint32)(&buffer))
}

func DeleteShader(shader Shader) {
	gl.DeleteShader(uint32(shader))
}

func DeleteTexture(texture Texture) {
	gl.DeleteTextures(1, (*uint32)(&texture))
}

func DepthFunc(xfunc Enum) {
	gl.DepthFunc(uint32(xfunc))
}

func DepthMask(flag bool) {
	gl.DepthMask(flag)
}

func DepthRange(near, far Clampf) {
	gl.DepthRange(float64(near), float64(far))
}

func DetachShader(program Program, shader Shader) {
	gl.DetachShader(uint32(program), uint32(shader))
}

func Disable(cap Enum) {
	gl.Disable(uint32(cap))
}

func DisableVertexAttribArray(index uint32) {
	gl.DisableVertexAttribArray(index)
}

func DrawArrays(mode Enum, first int32, count int32) {
	gl.DrawArrays(uint32(mode), first, count)
}

func DrawElements[T IndexType](mode Enum, indices []T) {
	xtype := gl.UNSIGNED_BYTE
	switch reflect.ValueOf(indices[0]).Interface().(type) {
	case uint8:
		xtype = gl.UNSIGNED_BYTE
	case uint16:
		xtype = gl.UNSIGNED_SHORT
	case uint32:
		xtype = gl.UNSIGNED_INT
	}
	gl.DrawElements(uint32(mode), int32(len(indices)), uint32(xtype), gl.Ptr(indices))
}

func Enable(cap Enum) {
	gl.Enable(uint32(cap))
}

func EnableVertexAttribArray(index uint32) {
	gl.EnableVertexAttribArray(index)
}

func Finish() {
	gl.Finish()
}

func Flush() {
	gl.Flush()
}

func FramebufferRenderbuffer(target, attachment, renderbuffertarget Enum, renderbuffer Renderbuffer) {
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachment), uint32(renderbuffertarget), uint32(renderbuffer))
}

func FramebufferTexture2D(target, attachment, textarget Enum, texture Texture, level int32) {
	gl.FramebufferTexture2D(uint32(target), uint32(attachment), uint32(textarget), uint32(texture), level)
}

func FrontFace(mode Enum) {
	gl.FrontFace(uint32(mode))
}

func GenerateMipmap(target Enum) {
	gl.GenerateMipmap(uint32(target))
}

func GetError() Enum {
	return Enum(gl.GetError())
}

func BufferData(target Enum, size int, data unsafe.Pointer, usage Enum) {
	gl.BufferData(uint32(target), size, data, uint32(usage))
}

func BufferSubData(target Enum, offset int, size int, data unsafe.Pointer) {
	gl.BufferSubData(uint32(target), offset, size, data)
}

func VertexAttribPointer(index uint32, size int32, xtype Enum, normalized bool, stride int32, offset uintptr) {
	gl.VertexAttribPointerWithOffset(index, size, uint32(xtype), normalized, stride, offset)
}

func Init() error {
	return gl.Init()
}

func Viewport(x int32, y int32, width int32, height int32) {
	gl.Viewport(x, y, width, height)
}

func Scissor(x int32, y int32, width int32, height int32) {
	gl.Scissor(x, y, width, height)
}

func Uniform1i(location Uniform, v0 int32) {
	gl.Uniform1i(int32(location), v0)
}

func Uniform1f(location Uniform, v0 float32) {
	gl.Uniform1f(int32(location), v0)
}

func Uniform2fv(location Uniform, value []float32) {
	gl.Uniform2fv(int32(location), int32(len(value)), unsafe.SliceData(value))
}

func Uniform3fv(location Uniform, value []float32) {
	gl.Uniform3fv(int32(location), int32(len(value)), unsafe.SliceData(value))
}

func Uniform4fv(location Uniform, value []float32) {
	gl.Uniform4fv(int32(location), int32(len(value)), unsafe.SliceData(value))
}

func UniformMatrix2fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix2fv(int32(location), int32(len(value)/4), transpose, unsafe.SliceData(value))
}

func UniformMatrix3fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix3fv(int32(location), int32(len(value)/9), transpose, unsafe.SliceData(value))
}

func UniformMatrix4fv(location Uniform, transpose bool, value []float32) {
	gl.UniformMatrix4fv(int32(location), int32(len(value)/16), transpose, unsafe.SliceData(value))
}

func ShaderSource(shader Shader, source string) {
	csource, free := gl.Strs(source + "\x00")
	defer free()
	gl.ShaderSource(uint32(shader), 1, csource, nil)
}

func GetShaderParameter(shader Shader, pname Enum) any {
	var v int32
	gl.GetShaderiv(uint32(shader), uint32(pname), &v)
	switch pname {
	case SHADER_TYPE:
		return Enum(v)
	case DELETE_STATUS, COMPILE_STATUS:
		return v == TRUE
	default:
		panic(fmt.Sprintf("invalid shader parameter: %v", pname))
	}
}

func GetShaderInfoLog(shader Shader) string {
	var logLen int32
	gl.GetShaderiv(uint32(shader), gl.INFO_LOG_LENGTH, &logLen)
	b := make([]byte, logLen)
	gl.GetShaderInfoLog(uint32(shader), logLen, nil, &b[0])
	return string(b)
}

func LinkProgram(program Program) {
	gl.LinkProgram(uint32(program))
}

func GetProgramParameter(program Program, pname Enum) any {
	var v int32
	gl.GetProgramiv(uint32(program), uint32(pname), &v)
	switch pname {
	case DELETE_STATUS, LINK_STATUS, VALIDATE_STATUS:
		return v == TRUE
	case ATTACHED_SHADERS, ACTIVE_ATTRIBUTES, ACTIVE_UNIFORMS, TRANSFORM_FEEDBACK_VARYINGS, ACTIVE_UNIFORM_BLOCKS:
		return v
	case TRANSFORM_FEEDBACK_BUFFER_MODE:
		return Enum(v)
	default:
		panic(fmt.Sprintf("invalid program parameter: %v", pname))
	}
}

func GetProgramInfoLog(program Program) string {
	var logLen int32
	gl.GetProgramiv(uint32(program), gl.INFO_LOG_LENGTH, &logLen)
	b := make([]byte, logLen)
	gl.GetProgramInfoLog(uint32(program), logLen, nil, &b[0])
	return string(b)
}

func GetUniformLocation(program Program, name string) Uniform {
	return Uniform(gl.GetUniformLocation(uint32(program), gl.Str(name+"\x00")))
}

func UseProgram(program Program) {
	gl.UseProgram(uint32(program))
}

func TexImage2D[T SliceConstraints](target Enum, level int32, internalformat Enum, width int32, height int32, border int32, format Enum, xtype Enum, pixels []T) {
	gl.TexImage2D(uint32(target), level, int32(internalformat), width, height, border, uint32(format), uint32(xtype), gl.Ptr(pixels))
}

func TexParameteri(target Enum, pname Enum, param int32) {
	gl.TexParameteri(uint32(target), uint32(pname), param)
}

func GetParameter(pname Enum) any {
	switch pname {
	case ACTIVE_TEXTURE, BLEND_DST_ALPHA, BLEND_DST_RGB, BLEND_EQUATION_ALPHA, BLEND_EQUATION_RGB, BLEND_SRC_ALPHA, BLEND_SRC_RGB, CULL_FACE_MODE, DEPTH_FUNC, FRONT_FACE, GENERATE_MIPMAP_HINT, IMPLEMENTATION_COLOR_READ_FORMAT, IMPLEMENTATION_COLOR_READ_TYPE, STENCIL_BACK_FAIL, STENCIL_BACK_FUNC, STENCIL_BACK_PASS_DEPTH_FAIL, STENCIL_BACK_PASS_DEPTH_PASS, STENCIL_FAIL, STENCIL_FUNC, STENCIL_PASS_DEPTH_FAIL, STENCIL_PASS_DEPTH_PASS, DRAW_BUFFER0, DRAW_BUFFER1, DRAW_BUFFER2, DRAW_BUFFER3, DRAW_BUFFER4, DRAW_BUFFER5, DRAW_BUFFER6, DRAW_BUFFER7, DRAW_BUFFER8, DRAW_BUFFER9, DRAW_BUFFER10, DRAW_BUFFER11, DRAW_BUFFER12, DRAW_BUFFER13, DRAW_BUFFER14, DRAW_BUFFER15, FRAGMENT_SHADER_DERIVATIVE_HINT, READ_BUFFER:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Enum(v)
	case ALIASED_LINE_WIDTH_RANGE, ALIASED_POINT_SIZE_RANGE, DEPTH_RANGE:
		v := make([]float32, 2)
		gl.GetFloatv(uint32(pname), &v[0])
		return v
	case ALPHA_BITS, GREEN_BITS, MAX_COMBINED_TEXTURE_IMAGE_UNITS, MAX_CUBE_MAP_TEXTURE_SIZE, MAX_FRAGMENT_UNIFORM_VECTORS, MAX_RENDERBUFFER_SIZE, MAX_TEXTURE_IMAGE_UNITS, MAX_TEXTURE_SIZE, MAX_VARYING_VECTORS, MAX_VERTEX_ATTRIBS, MAX_VERTEX_TEXTURE_IMAGE_UNITS, MAX_VERTEX_UNIFORM_VECTORS, PACK_ALIGNMENT, RED_BITS, SAMPLE_BUFFERS, SAMPLES, STENCIL_BACK_REF, STENCIL_BITS, STENCIL_CLEAR_VALUE, STENCIL_REF, SUBPIXEL_BITS, UNPACK_ALIGNMENT, MAX_3D_TEXTURE_SIZE, MAX_ARRAY_TEXTURE_LAYERS, MAX_COLOR_ATTACHMENTS, MAX_COMBINED_UNIFORM_BLOCKS, MAX_DRAW_BUFFERS, MAX_ELEMENTS_INDICES, MAX_ELEMENTS_VERTICES, MAX_FRAGMENT_INPUT_COMPONENTS, MAX_FRAGMENT_UNIFORM_BLOCKS, MAX_FRAGMENT_UNIFORM_COMPONENTS, MAX_PROGRAM_TEXEL_OFFSET, MAX_SAMPLES, MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, MAX_UNIFORM_BUFFER_BINDINGS, MAX_VARYING_COMPONENTS, MAX_VERTEX_OUTPUT_COMPONENTS, MAX_VERTEX_UNIFORM_BLOCKS, MAX_VERTEX_UNIFORM_COMPONENTS, MIN_PROGRAM_TEXEL_OFFSET, PACK_ROW_LENGTH, PACK_SKIP_PIXELS, PACK_SKIP_ROWS, UNIFORM_BUFFER_OFFSET_ALIGNMENT, UNPACK_IMAGE_HEIGHT, UNPACK_ROW_LENGTH, UNPACK_SKIP_IMAGES, UNPACK_SKIP_PIXELS, UNPACK_SKIP_ROWS:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return v
	case ARRAY_BUFFER_BINDING:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Buffer(v)
	case BLEND, CULL_FACE, DEPTH_TEST, DEPTH_WRITEMASK, DITHER, POLYGON_OFFSET_FILL, SAMPLE_ALPHA_TO_COVERAGE, SAMPLE_COVERAGE, SAMPLE_COVERAGE_INVERT, SCISSOR_TEST, STENCIL_TEST, RASTERIZER_DISCARD, TRANSFORM_FEEDBACK_ACTIVE, TRANSFORM_FEEDBACK_PAUSED:
		var v bool
		gl.GetBooleanv(uint32(pname), &v)
		return v
	case BLEND_COLOR, COLOR_CLEAR_VALUE:
		v := make([]float32, 4)
		gl.GetFloatv(uint32(pname), &v[0])
		return v
	case BLUE_BITS, DEPTH_BITS:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return v
	case COLOR_WRITEMASK:
		v := make([]bool, 4)
		gl.GetBooleanv(uint32(pname), &v[0])
		return v
	case COMPRESSED_TEXTURE_FORMATS:
		var n int32
		gl.GetIntegerv(uint32(pname), &n)
		v := make([]int32, n)
		gl.GetIntegerv(uint32(pname), &v[0])
		return v
	case CURRENT_PROGRAM:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Program(v)
	case DEPTH_CLEAR_VALUE, LINE_WIDTH, POLYGON_OFFSET_FACTOR, POLYGON_OFFSET_UNITS, SAMPLE_COVERAGE_VALUE, MAX_TEXTURE_LOD_BIAS:
		var v float32
		gl.GetFloatv(uint32(pname), &v)
		return v
	case ELEMENT_ARRAY_BUFFER_BINDING, COPY_READ_BUFFER_BINDING, COPY_WRITE_BUFFER_BINDING, PIXEL_PACK_BUFFER_BINDING, PIXEL_UNPACK_BUFFER_BINDING, TRANSFORM_FEEDBACK_BUFFER_BINDING, UNIFORM_BUFFER_BINDING:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Buffer(v)
	case FRAMEBUFFER_BINDING, READ_FRAMEBUFFER_BINDING:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Framebuffer(v)
	case STENCIL_BACK_VALUE_MASK, STENCIL_BACK_WRITEMASK, STENCIL_VALUE_MASK, STENCIL_WRITEMASK:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return uint32(v)
	case MAX_VIEWPORT_DIMS:
		v := make([]int32, 2)
		gl.GetIntegerv(uint32(pname), &v[0])
		return v
	case RENDERBUFFER_BINDING:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Renderbuffer(v)
	case SCISSOR_BOX, VIEWPORT:
		v := make([]int32, 4)
		gl.GetIntegerv(uint32(pname), &v[0])
		return v
	case TEXTURE_BINDING_2D, TEXTURE_BINDING_CUBE_MAP, TEXTURE_BINDING_2D_ARRAY, TEXTURE_BINDING_3D:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Texture(v)
	case VERTEX_ARRAY_BINDING:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return VertexArrayObject(v)
	case MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, MAX_ELEMENT_INDEX, MAX_SERVER_WAIT_TIMEOUT, MAX_UNIFORM_BLOCK_SIZE:
		var v int64
		gl.GetInteger64v(uint32(pname), &v)
		return v
	case SAMPLER_BINDING:
		var v int32
		gl.GetIntegerv(uint32(pname), &v)
		return Sampler(v)
	default:
		panic(fmt.Sprintf("invalid parameter name: %v", pname))
	}
}
