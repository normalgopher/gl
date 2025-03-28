//go:build js

package glfw

import (
	"errors"
	"fmt"
	"syscall/js"

	"github.com/normalgopher/gl/internal"
)

func Init() error {
	return nil
}

func Terminate() {
}

var (
	hints = map[Hint]int{}
)

func WindowHint(target Hint, hint int) {
	hints[target] = hint
}

var (
	htmlWindow = js.Global().Get("window")
	document   = js.Global().Get("document")
	navigator  = js.Global().Get("navigator")
)

func getCanvas() js.Value {
	canvas := document.Call("getElementById", "glfw")
	if canvas.Equal(js.Null()) {
		canvas = document.Call("querySelector", "canvas")
	}
	return canvas
}

type contextAttributes struct {
	Alpha                           bool
	Depth                           bool
	Stencil                         bool
	Antialias                       bool
	PremultipliedAlpha              bool
	PreserveDrawingBuffer           bool
	PreferLowPowerToHighPerformance bool
	FailIfMajorPerformanceCaveat    bool
}

func (ca contextAttributes) Map() map[string]any {
	return map[string]any{
		"alpha":                           ca.Alpha,
		"depth":                           ca.Depth,
		"stencil":                         ca.Stencil,
		"antialias":                       ca.Antialias,
		"premultipliedAlpha":              ca.PremultipliedAlpha,
		"preserveDrawingBuffer":           ca.PreserveDrawingBuffer,
		"preferLowPowerToHighPerformance": ca.PreferLowPowerToHighPerformance,
		"failIfMajorPerformanceCaveat":    ca.FailIfMajorPerformanceCaveat,
	}
}

// https://www.glfw.org/docs/3.3/window_guide.html
func defaultAttributes() *contextAttributes {
	return &contextAttributes{
		Alpha:                 false,
		Depth:                 true,
		Stencil:               false,
		Antialias:             false,
		PremultipliedAlpha:    false,
		PreserveDrawingBuffer: false,
	}
}

func createContext(canvas js.Value, attrs contextAttributes) (js.Value, error) {
	gl := canvas.Call("getContext", "webgl2", attrs.Map())
	if gl.Equal(js.Null()) {
		return js.Null(), errors.New("can't create webgl2 context")
	}

	return gl, nil
}

func CreateWindow(width int, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	canvas := getCanvas()

	canvas.Set("width", width)
	canvas.Set("height", height)
	canvas.Get("style").Call("setProperty", "width", fmt.Sprintf("%vpx", width))
	canvas.Get("style").Call("setProperty", "height", fmt.Sprintf("%vpx", height))

	document.Set("title", title)

	// TODO: store and check hints
	attrs := defaultAttributes()
	context, err := internal.NewContext(canvas, attrs.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	w := &Window{
		canvas:  canvas,
		context: context,
		keys:    make(map[Key]Action),
		mouse:   make(map[MouseButton]Action),
	}

	if w.canvas.Get("requestPointerLock").Equal(js.Undefined()) ||
		document.Get("exitPointerLock").Equal(js.Undefined()) {

		w.missing.pointerLock = true
	}

	w.setupEventListener()

	return w, nil
}

func SwapInterval(interval int) {
	// TODO: implement using a requestAnimationFrame loop
	// glfw.SwapInterval(interval)
}

func GetPrimaryMonitor() *Monitor {
	return &Monitor{}
}

func PollEvents() {
	// glfw.PollEvents()
}

type Monitor struct{}

type VidMode struct {
	Width       int // The width, in pixels, of the video mode.
	Height      int // The height, in pixels, of the video mode.
	RedBits     int // The bit depth of the red channel of the video mode.
	GreenBits   int // The bit depth of the green channel of the video mode.
	BlueBits    int // The bit depth of the blue channel of the video mode.
	RefreshRate int // The refresh rate, in Hz, of the video mode.
}

func (m *Monitor) GetVideoMode() *VidMode {
	return &VidMode{
		Width:       js.Global().Get("innerWidth").Int(),
		Height:      js.Global().Get("innerHeight").Int(),
		RedBits:     8,
		GreenBits:   8,
		BlueBits:    8,
		RefreshRate: 60, // TODO: this has to be queryable somewhere: it's not
	}
}
