package glhelp

const (
	BaseShaderHeader = "#version 300 es\n"
	floatPrecision   = "#ifdef GL_ES\n\tprecision highp float;\n#endif\n"
)

func VertexShaderHeader() string {
	return BaseShaderHeader
}

func FragmentShaderHeader() string {
	return BaseShaderHeader + "\n" + floatPrecision
}
