//go:build js

package glfw

import (
	"syscall/js"

	"github.com/normalgopher/gl/internal"
)

type Window struct {
	canvas  js.Value
	context *internal.Context

	shouldClose bool
	hidden      bool

	// Unavailable browser APIs
	missing struct {
		pointerLock bool // Pointer Lock API.
		fullscreen  bool // Fullscreen API.
	}

	keys      map[Key]Action
	mouse     map[MouseButton]Action
	cursorPos [2]float64

	framebufferSizeCB FramebufferSizeCallback
	sizeCB            SizeCallback
	keyCB             KeyCallback
	charCB            CharCallback
	mouseButtonCB     MouseButtonCallback
	cursorPosCB       CursorPosCallback
	scrollCB          ScrollCallback
	focusCB           FocusCallback
}

func (w *Window) setupEventListener() {
	history := htmlWindow.Get("history")
	// ???
	history.Call("pushState", nil, nil)
	history.Call("pushState", nil, nil)
	history.Call("pushState", nil, nil)

	htmlWindow.Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) any {
		history.Call("pushState", nil, nil)
		return nil
	}))

	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) any {
		width := js.Global().Get("innerWidth").Int()
		height := js.Global().Get("innerHeight").Int()

		w.canvas.Set("width", width)
		w.canvas.Set("height", height)
		w.canvas.Get("style").Call("setProperty", "width", width)
		w.canvas.Get("style").Call("setProperty", "height", height)

		if w.framebufferSizeCB != nil {
			w.framebufferSizeCB(w, width, height)
		}
		if w.sizeCB != nil {
			w.sizeCB(w, width, height)
		}

		return nil
	}))

	document.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) any {
		ke := args[0]
		ke.Call("preventDefault")

		action := Press
		if ke.Get("repeat").Bool() {
			action = Repeat
		}

		key := toKey(ke)

		if key != KeyUnknown {
			w.keys[key] = action
		}

		if w.keyCB != nil {
			mods := modKeys(ke)

			w.keyCB(w, key, -1, action, mods)
		}

		if w.charCB != nil {
			keyStr := ke.Get("key").String()
			if len(keyStr) == 1 {
				keyRunes := []rune(keyStr)
				w.charCB(w, keyRunes[0])
			}
		}

		return nil
	}))

	document.Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) any {
		ke := args[0]
		ke.Call("preventDefault")

		key := toKey(ke)

		if key != KeyUnknown {
			w.keys[key] = Release
		}

		if w.keyCB != nil {
			mods := modKeys(ke)

			w.keyCB(w, key, -1, Release, mods)
		}

		return nil
	}))

	document.Call("addEventListener", "mousedown", js.FuncOf(func(this js.Value, args []js.Value) any {
		me := args[0]
		me.Call("preventDefault")

		button := MouseButton(me.Get("button").Int())
		if !(button >= MouseButton1 && button < MouseButtonLast) {
			return nil
		}

		w.mouse[button] = Press

		if w.mouseButtonCB != nil {
			w.mouseButtonCB(w, button, Press, 0)
		}

		return nil
	}))

	document.Call("addEventListener", "mouseup", js.FuncOf(func(this js.Value, args []js.Value) any {
		me := args[0]
		me.Call("preventDefault")

		button := MouseButton(me.Get("button").Int())
		if !(button >= MouseButton1 && button < MouseButtonLast) {
			return nil
		}

		w.mouse[button] = Release

		if w.mouseButtonCB != nil {
			w.mouseButtonCB(w, button, Release, 0)
		}

		return nil
	}))

	document.Call("addEventListener", "contextmenu", js.FuncOf(func(this js.Value, args []js.Value) any {
		me := args[0]
		me.Call("preventDefault")

		return nil
	}))

	document.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) any {
		me := args[0]
		me.Call("preventDefault")

		if w.cursorPosCB != nil {
			w.cursorPosCB(w, me.Get("clientX").Float(), me.Get("clientY").Float())
		}

		return nil
	}))

	document.Call("addEventListener", "wheel", js.FuncOf(func(this js.Value, args []js.Value) any {
		we := args[0]
		we.Call("preventDefault")
		deltaX, deltaY := we.Get("deltaX").Float(), we.Get("deltaY").Float()

		if deltaX > 0 {
			deltaX = 1
		} else if deltaX < 0 {
			deltaX = -1
		}

		if deltaY > 0 {
			deltaY = 1
		} else if deltaY < 0 {
			deltaY = -1
		}

		if w.scrollCB != nil {
			w.scrollCB(w, deltaX, deltaY)
		}

		return nil
	}), map[string]any{"passive": false})

	htmlWindow.Call("addEventListener", "focus", js.FuncOf(func(this js.Value, args []js.Value) any {
		if w.focusCB != nil {
			w.focusCB(w, true)
		}

		return nil
	}))

	htmlWindow.Call("addEventListener", "blur", js.FuncOf(func(this js.Value, args []js.Value) any {
		for m := range w.mouse {
			w.mouse[m] = Release
		}
		for k := range w.keys {
			w.keys[k] = Release
		}

		if w.focusCB != nil {
			w.focusCB(w, false)
		}

		return nil
	}))

	document.Call("addEventListener", "visibilitychange", js.FuncOf(func(this js.Value, args []js.Value) any {
		go func() {
			state := document.Get("visibilityState").String()

			switch state {
			case "hidden":
				w.hidden = true
				for m := range w.mouse {
					w.mouse[m] = Release
				}
				for k := range w.keys {
					w.keys[k] = Release
				}
			case "visible":
				w.hidden = false
			}
		}()

		return nil
	}))
}

func (w *Window) MakeContextCurrent() {
	internal.SetCurrent(w.context)
}

func (w *Window) Destroy() {
	w.canvas.Call("remove")
}

func (w *Window) GetSize() (width, height int) {
	width = w.canvas.Call("getBoundingClientRect").Get("width").Int()
	height = w.canvas.Call("getBoundingClientRect").Get("height").Int()
	return width, height
}

func (w *Window) GetFramebufferSize() (width, height int) {
	return w.canvas.Get("width").Int(), w.canvas.Get("height").Int()
}

func (w *Window) SetSizeCallback(cb SizeCallback) SizeCallback {
	prev := w.sizeCB
	w.sizeCB = cb
	return prev
}

func (w *Window) SetFramebufferSizeCallback(cb FramebufferSizeCallback) FramebufferSizeCallback {
	prev := w.framebufferSizeCB
	w.framebufferSizeCB = cb
	return prev
}

func (w *Window) SetCursorPosCallback(cb CursorPosCallback) CursorPosCallback {
	prev := w.cursorPosCB
	w.cursorPosCB = cb
	return prev
}

func (w *Window) SetMouseButtonCallback(cb MouseButtonCallback) MouseButtonCallback {
	prev := w.mouseButtonCB
	w.mouseButtonCB = cb
	return prev
}

func (w *Window) SetScrollCallback(cb ScrollCallback) ScrollCallback {
	prev := w.scrollCB
	w.scrollCB = cb
	return prev
}

func (w *Window) SetKeyCallback(cb KeyCallback) KeyCallback {
	prev := w.keyCB
	w.keyCB = cb
	return prev
}

func (w *Window) SetCharCallback(cb CharCallback) CharCallback {
	prev := w.charCB
	w.charCB = cb
	return prev
}

func (w *Window) SetFocusCallback(cb FocusCallback) FocusCallback {
	prev := w.focusCB
	w.focusCB = cb
	return prev
}

func (w *Window) SwapBuffers() {
	// TODO: maybe create 2 framebuffers and swap between them?
}

func (w *Window) ShouldClose() bool {
	return w.shouldClose
}

func (w *Window) SetShouldClose(closed bool) {
	w.shouldClose = closed
}

func (w *Window) GetKey(key Key) Action {
	return w.keys[key]
}

func (w *Window) GetMouseButton(button MouseButton) Action {
	return w.mouse[button]
}

func (w *Window) SetInputMode(mode InputMode, value int) {
	switch mode {
	case CursorMode:
		if w.missing.pointerLock {
			return
		}
		switch value {
		case CursorNormal:
			document.Call("exitPointerLock")
			w.canvas.Get("style").Call("setProperty", "cursor", "initial")
		case CursorHidden:
			w.canvas.Call("exitPointerLock")
			w.canvas.Get("style").Call("setProperty", "cursor", "none")
		case CursorDisabled:
			w.canvas.Call("requestPointerLock")
		}
	}
}
