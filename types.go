package gl

// Floating-point value, clamped to the range [0,1]
type Clampf float32

func (c Clampf) Value() float32 {
	return float32(c)
}

type Enum uint32

func (e Enum) Value() uint32 {
	return uint32(e)
}

type SliceConstraints interface {
	~float32 | ~uint8 | ~uint16 | ~uint32 | ~int32
}

type IndexType interface {
	~uint8 | ~uint16 | ~uint32
}
