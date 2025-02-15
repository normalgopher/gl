//go:build js

package glfw

import "syscall/js"

var keycodeMap = map[string]Key{
	"Space":        KeySpace,
	"Quote":        KeyApostrophe, //???
	"Comma":        KeyComma,
	"Minus":        KeyMinus,
	"Period":       KeyPeriod,
	"Slash":        KeySlash,
	"Digit0":       Key0,
	"Digit1":       Key1,
	"Digit2":       Key2,
	"Digit3":       Key3,
	"Digit4":       Key4,
	"Digit5":       Key5,
	"Digit6":       Key6,
	"Digit7":       Key7,
	"Digit8":       Key8,
	"Digit9":       Key9,
	"Semicolon":    KeySemicolon,
	"Equal":        KeyEqual,
	"KeyA":         KeyA,
	"KeyB":         KeyB,
	"KeyC":         KeyC,
	"KeyD":         KeyD,
	"KeyE":         KeyE,
	"KeyF":         KeyF,
	"KeyG":         KeyG,
	"KeyH":         KeyH,
	"KeyI":         KeyI,
	"KeyJ":         KeyJ,
	"KeyK":         KeyK,
	"KeyL":         KeyL,
	"KeyM":         KeyM,
	"KeyN":         KeyN,
	"KeyO":         KeyO,
	"KeyP":         KeyP,
	"KeyQ":         KeyQ,
	"KeyR":         KeyR,
	"KeyS":         KeyS,
	"KeyT":         KeyT,
	"KeyU":         KeyU,
	"KeyV":         KeyV,
	"KeyW":         KeyW,
	"KeyX":         KeyX,
	"KeyY":         KeyY,
	"KeyZ":         KeyZ,
	"BracketLeft":  KeyLeftBracket,
	"Backslash":    KeyBackslash,
	"BracketRight": KeyRightBracket,
	//	"KeyGraveAccent": KeyGraveAccent,
	// "KeyWorld1": KeyWorld1,
	// "KeyWorld2": KeyWorld2,
	"Escape":      KeyEscape,
	"Enter":       KeyEnter,
	"Tab":         KeyTab,
	"Backspace":   KeyBackspace,
	"Insert":      KeyInsert,
	"Delete":      KeyDelete,
	"ArrowRight":  KeyRight,
	"ArrowLeft":   KeyLeft,
	"ArrowDown":   KeyDown,
	"ArrowUp":     KeyUp,
	"PageUp":      KeyPageUp,
	"PageDown":    KeyPageDown,
	"Home":        KeyHome,
	"End":         KeyEnd,
	"CapsLock":    KeyCapsLock,
	"ScrollLock":  KeyScrollLock,
	"NumLock":     KeyNumLock,
	"PrintScreen": KeyPrintScreen,
	"Pause":       KeyPause,
	"F1":          KeyF1,
	"F2":          KeyF2,
	"F3":          KeyF3,
	"F4":          KeyF4,
	"F5":          KeyF5,
	"F6":          KeyF6,
	"F7":          KeyF7,
	"F8":          KeyF8,
	"F9":          KeyF9,
	"F10":         KeyF10,
	"F11":         KeyF11,
	"F12":         KeyF12,
	"F13":         KeyF13,
	"F14":         KeyF14,
	"F15":         KeyF15,
	"F16":         KeyF16,
	"F17":         KeyF17,
	"F18":         KeyF18,
	"F19":         KeyF19,
	"F20":         KeyF20,
	"F21":         KeyF21,
	"F22":         KeyF22,
	"F23":         KeyF23,
	"F24":         KeyF24,
	// "F25": KeyF25,
	"Numpad0":        KeyKP0,
	"Numpad1":        KeyKP1,
	"Numpad2":        KeyKP2,
	"Numpad3":        KeyKP3,
	"Numpad4":        KeyKP4,
	"Numpad5":        KeyKP5,
	"Numpad6":        KeyKP6,
	"Numpad7":        KeyKP7,
	"Numpad8":        KeyKP8,
	"Numpad9":        KeyKP9,
	"NumpadDecimal":  KeyKPDecimal,
	"NumpadDivide":   KeyKPDivide,
	"NumpadMultiply": KeyKPMultiply,
	"NumpadSubtract": KeyKPSubtract,
	"NumpadAdd":      KeyKPAdd,
	"NumpadEnter":    KeyKPEnter,
	"NumpadEqual":    KeyKPEqual,
	"ShiftLeft":      KeyLeftShift,
	"ControlLeft":    KeyLeftControl,
	"AltLeft":        KeyLeftAlt,
	"OSLeft":         KeyLeftSuper,
	"MetaLeft":       KeyLeftSuper,
	"ShiftRight":     KeyRightShift,
	"ControlRight":   KeyRightControl,
	"AltRight":       KeyRightAlt,
	"OSRight":        KeyRightSuper,
	"MetaRight":      KeyRightSuper,
	"ContextMenu":    KeyMenu, // ????
}

func toKey(ke js.Value) Key {
	keyStr := ke.Get("code").String()
	key, ok := keycodeMap[keyStr]
	if !ok {
		return KeyUnknown
	}
	return key
}

func modKeys(ke js.Value) ModifierKey {
	mods := ModifierKey(0)
	if ke.Get("shiftKey").Bool() {
		mods += ModShift
	}
	if ke.Get("ctrlKey").Bool() {
		mods += ModControl
	}
	if ke.Get("altKey").Bool() {
		mods += ModAlt
	}
	if ke.Get("metaKey").Bool() {
		mods += ModSuper
	}
	return mods
}
