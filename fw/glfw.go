//go:build !js

package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func Init() error {
	return glfw.Init()
}

func Terminate() {
	glfw.Terminate()
}

func WindowHint(target Hint, hint int) {
	glfw.WindowHint(glfw.Hint(target), hint)
}

func CreateWindow(width int, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	var mon *glfw.Monitor
	var shar *glfw.Window
	if monitor != nil {
		mon = monitor.internal
	}
	if share != nil {
		shar = share.internal
	}
	w, err := glfw.CreateWindow(width, height, title, mon, shar)
	if err != nil {
		return nil, err
	}

	return &Window{
		internal: w,
	}, nil
}

func SwapInterval(interval int) {
	glfw.SwapInterval(interval)
}

type Monitor struct {
	internal *glfw.Monitor
}

func GetPrimaryMonitor() *Monitor {
	return &Monitor{
		internal: glfw.GetPrimaryMonitor(),
	}
}

func PollEvents() {
	glfw.PollEvents()
}

type VidMode struct {
	Width       int // The width, in pixels, of the video mode.
	Height      int // The height, in pixels, of the video mode.
	RedBits     int // The bit depth of the red channel of the video mode.
	GreenBits   int // The bit depth of the green channel of the video mode.
	BlueBits    int // The bit depth of the blue channel of the video mode.
	RefreshRate int // The refresh rate, in Hz, of the video mode.
}

func (m *Monitor) GetVideoMode() *VidMode {
	v := m.internal.GetVideoMode()
	return &VidMode{
		Width:       v.Width,
		Height:      v.Height,
		RedBits:     v.RedBits,
		GreenBits:   v.GreenBits,
		BlueBits:    v.BlueBits,
		RefreshRate: v.RefreshRate,
	}
}
