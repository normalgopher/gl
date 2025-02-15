//go:build !js

package glfw

import "github.com/go-gl/glfw/v3.3/glfw"

type Window struct {
	internal *glfw.Window

	framebufferSizeCB FramebufferSizeCallback
	sizeCB            SizeCallback
	keyCB             KeyCallback
	charCB            CharCallback
	mouseButtonCB     MouseButtonCallback
	cursorPosCB       CursorPosCallback
	scrollCB          ScrollCallback
	focusCB           FocusCallback
}

func (w *Window) MakeContextCurrent() {
	w.internal.MakeContextCurrent()
}

func (w *Window) Destroy() {
	w.internal.Destroy()
}

func (w *Window) GetSize() (width, height int) {
	return w.internal.GetSize()
}

func (w *Window) GetFramebufferSize() (width, height int) {
	return w.internal.GetFramebufferSize()
}

func (w *Window) SetFocusCallback(cb FocusCallback) FocusCallback {
	prev := w.focusCB
	w.focusCB = cb
	w.internal.SetFocusCallback(func(_ *glfw.Window, focused bool) {
		cb(w, focused)
	})
	return prev
}

func (w *Window) SetSizeCallback(cb SizeCallback) SizeCallback {
	prev := w.sizeCB
	w.sizeCB = cb
	w.internal.SetSizeCallback(func(_ *glfw.Window, width, height int) {
		cb(w, width, height)
	})
	return prev
}

func (w *Window) SetFramebufferSizeCallback(cb FramebufferSizeCallback) FramebufferSizeCallback {
	prev := w.framebufferSizeCB
	w.framebufferSizeCB = cb
	w.internal.SetFramebufferSizeCallback(func(_ *glfw.Window, width, height int) {
		cb(w, width, height)
	})
	return prev
}

func (w *Window) SetCursorPosCallback(cb CursorPosCallback) CursorPosCallback {
	prev := w.cursorPosCB
	w.cursorPosCB = cb
	w.internal.SetCursorPosCallback(func(_ *glfw.Window, xpos, ypos float64) {
		cb(w, xpos, ypos)
	})
	return prev
}

func (w *Window) SetMouseButtonCallback(cb MouseButtonCallback) MouseButtonCallback {
	prev := w.mouseButtonCB
	w.mouseButtonCB = cb
	w.internal.SetMouseButtonCallback(func(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		cb(w, MouseButton(button), Action(action), ModifierKey(mods))
	})
	return prev
}

func (w *Window) SetScrollCallback(cb ScrollCallback) ScrollCallback {
	prev := w.scrollCB
	w.scrollCB = cb
	w.internal.SetScrollCallback(func(_ *glfw.Window, xoff, yoff float64) {
		cb(w, xoff, yoff)
	})
	return prev
}

func (w *Window) SetKeyCallback(cb KeyCallback) KeyCallback {
	prev := w.keyCB
	w.keyCB = cb
	w.internal.SetKeyCallback(func(_ *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		cb(w, Key(key), scancode, Action(action), ModifierKey(mods))
	})
	return prev
}

func (w *Window) SetCharCallback(cb CharCallback) CharCallback {
	prev := w.charCB
	w.charCB = cb
	w.internal.SetCharCallback(func(_ *glfw.Window, char rune) {
		cb(w, char)
	})
	return prev
}

func (w *Window) SwapBuffers() {
	w.internal.SwapBuffers()
}

func (w *Window) ShouldClose() bool {
	return w.internal.ShouldClose()
}

func (w *Window) SetShouldClose(closed bool) {
	w.internal.SetShouldClose(closed)
}

func (w *Window) GetKey(key Key) Action {
	return Action(w.internal.GetKey(glfw.Key(key)))
}

func (w *Window) GetMouseButton(button MouseButton) Action {
	return Action(w.internal.GetMouseButton(glfw.MouseButton(button)))
}

func (w *Window) SetInputMode(mode InputMode, value int) {
	w.internal.SetInputMode(glfw.InputMode(mode), value)
}
