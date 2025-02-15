//go:build js

package gl

import (
	"fmt"
	"reflect"
	"syscall/js"
	"unsafe"
)

var (
	jsUint8Array   = js.Global().Get("Uint8Array")
	jsUint16Array  = js.Global().Get("Uint16Array")
	jsUint32Array  = js.Global().Get("Uint32Array")
	jsInt32Array   = js.Global().Get("Int32Array")
	jsFloat32Array = js.Global().Get("Float32Array")
	jsMemory       = jsUint8Array.New(8)
	jsMemBuf       = jsMemory.Get("buffer")
)

var (
	uint8Type   = reflect.TypeFor[uint8]()
	uint16Type  = reflect.TypeFor[uint16]()
	uint32Type  = reflect.TypeFor[uint32]()
	int32Type   = reflect.TypeFor[int32]()
	float32Type = reflect.TypeFor[float32]()
)

func jsArrayType[T SliceConstraints]() (arrayType js.Value, size int) {
	t := reflect.TypeFor[T]()

	switch t {
	case uint8Type:
		arrayType = jsUint8Array
	case uint16Type:
		arrayType = jsUint16Array
	case uint32Type:
		arrayType = jsUint32Array
	case int32Type:
		arrayType = jsInt32Array
	case float32Type:
		arrayType = jsFloat32Array
	default:
		panic(fmt.Sprintf("invalid array type %s", t.Name()))
	}

	return arrayType, int(t.Size())
}

func goSliceToJsArray[T SliceConstraints](v []T) js.Value {
	jsData := js.Null()
	if v != nil {
		arrayType, size := jsArrayType[T]()
		if jsMemory.Length() < size*len(v) {
			jsMemory = jsUint8Array.New(size * len(v))
			jsMemBuf = jsMemory.Get("buffer")
		}
		byteSlice := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(v))), size*len(v))
		js.CopyBytesToJS(jsMemory, byteSlice)
		jsData = arrayType.New(jsMemBuf, 0, len(v))
	}
	return jsData
}

func goPtrToJsArray(v unsafe.Pointer, size int) js.Value {
	jsData := js.Null()
	if v != nil {
		jsData = jsUint8Array.New(size)
		if js.CopyBytesToJS(jsData, unsafe.Slice((*byte)(v), size)) != size {
			panic("failed to copy data to js")
		}
		jsData = jsData.Call("subarray", 0, size)
	}
	return jsData
}
