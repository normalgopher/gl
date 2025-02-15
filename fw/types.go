package glfw

type Hint int
type MouseButton int
type Action int
type ModifierKey int
type Key int
type InputMode int

type SizeCallback func(w *Window, width int, height int)
type FramebufferSizeCallback func(w *Window, width int, height int)
type CursorPosCallback func(w *Window, xpos float64, ypos float64)
type MouseButtonCallback func(w *Window, button MouseButton, action Action, mods ModifierKey)
type ScrollCallback func(w *Window, xoff float64, yoff float64)
type KeyCallback func(w *Window, key Key, scancode int, action Action, mods ModifierKey)
type CharCallback func(w *Window, char rune)
type FocusCallback func(w *Window, focused bool)
