package glfw

const (
	False = 0
	True  = 1
)

// Action types.
const (
	Release Action = 0 // The key or button was released.
	Press   Action = 1 // The key or button was pressed.
	Repeat  Action = 2 // The key was held down until it repeated.
)

// Mouse buttons.
const (
	MouseButton1      MouseButton = 0
	MouseButton2      MouseButton = 1
	MouseButton3      MouseButton = 2
	MouseButton4      MouseButton = 3
	MouseButton5      MouseButton = 4
	MouseButton6      MouseButton = 5
	MouseButton7      MouseButton = 6
	MouseButton8      MouseButton = 7
	MouseButtonLast   MouseButton = 8
	MouseButtonLeft   MouseButton = MouseButton1
	MouseButtonRight  MouseButton = MouseButton2
	MouseButtonMiddle MouseButton = MouseButton3
)

const (
	KeyUnknown      Key = -1
	KeySpace        Key = 32
	KeyApostrophe   Key = 39
	KeyComma        Key = 44
	KeyMinus        Key = 45
	KeyPeriod       Key = 46
	KeySlash        Key = 47
	Key0            Key = 48
	Key1            Key = 49
	Key2            Key = 50
	Key3            Key = 51
	Key4            Key = 52
	Key5            Key = 53
	Key6            Key = 54
	Key7            Key = 55
	Key8            Key = 56
	Key9            Key = 57
	KeySemicolon    Key = 59
	KeyEqual        Key = 61
	KeyA            Key = 65
	KeyB            Key = 66
	KeyC            Key = 67
	KeyD            Key = 68
	KeyE            Key = 69
	KeyF            Key = 70
	KeyG            Key = 71
	KeyH            Key = 72
	KeyI            Key = 73
	KeyJ            Key = 74
	KeyK            Key = 75
	KeyL            Key = 76
	KeyM            Key = 77
	KeyN            Key = 78
	KeyO            Key = 79
	KeyP            Key = 80
	KeyQ            Key = 81
	KeyR            Key = 82
	KeyS            Key = 83
	KeyT            Key = 84
	KeyU            Key = 85
	KeyV            Key = 86
	KeyW            Key = 87
	KeyX            Key = 88
	KeyY            Key = 89
	KeyZ            Key = 90
	KeyLeftBracket  Key = 91
	KeyBackslash    Key = 92
	KeyRightBracket Key = 93
	KeyGraveAccent  Key = 96
	KeyWorld1       Key = 161
	KeyWorld2       Key = 162
	KeyEscape       Key = 256
	KeyEnter        Key = 257
	KeyTab          Key = 258
	KeyBackspace    Key = 259
	KeyInsert       Key = 260
	KeyDelete       Key = 261
	KeyRight        Key = 262
	KeyLeft         Key = 263
	KeyDown         Key = 264
	KeyUp           Key = 265
	KeyPageUp       Key = 266
	KeyPageDown     Key = 267
	KeyHome         Key = 268
	KeyEnd          Key = 269
	KeyCapsLock     Key = 280
	KeyScrollLock   Key = 281
	KeyNumLock      Key = 282
	KeyPrintScreen  Key = 283
	KeyPause        Key = 284
	KeyF1           Key = 290
	KeyF2           Key = 291
	KeyF3           Key = 292
	KeyF4           Key = 293
	KeyF5           Key = 294
	KeyF6           Key = 295
	KeyF7           Key = 296
	KeyF8           Key = 297
	KeyF9           Key = 298
	KeyF10          Key = 299
	KeyF11          Key = 300
	KeyF12          Key = 301
	KeyF13          Key = 302
	KeyF14          Key = 303
	KeyF15          Key = 304
	KeyF16          Key = 305
	KeyF17          Key = 306
	KeyF18          Key = 307
	KeyF19          Key = 308
	KeyF20          Key = 309
	KeyF21          Key = 310
	KeyF22          Key = 311
	KeyF23          Key = 312
	KeyF24          Key = 313
	KeyF25          Key = 314
	KeyKP0          Key = 320
	KeyKP1          Key = 321
	KeyKP2          Key = 322
	KeyKP3          Key = 323
	KeyKP4          Key = 324
	KeyKP5          Key = 325
	KeyKP6          Key = 326
	KeyKP7          Key = 327
	KeyKP8          Key = 328
	KeyKP9          Key = 329
	KeyKPDecimal    Key = 330
	KeyKPDivide     Key = 331
	KeyKPMultiply   Key = 332
	KeyKPSubtract   Key = 333
	KeyKPAdd        Key = 334
	KeyKPEnter      Key = 335
	KeyKPEqual      Key = 337
	KeyLeftShift    Key = 340
	KeyLeftControl  Key = 341
	KeyLeftAlt      Key = 342
	KeyLeftSuper    Key = 343
	KeyRightShift   Key = 344
	KeyRightControl Key = 345
	KeyRightAlt     Key = 346
	KeyRightSuper   Key = 347
	KeyMenu         Key = 348
	KeyLast         Key = 349
)

const (
	ModShift   ModifierKey = 1
	ModControl ModifierKey = 2
	ModAlt     ModifierKey = 4
	ModSuper   ModifierKey = 8
)

// Init related hints. (Use with glfw.InitHint)
const (
	JoystickHatButtons  Hint = 327681 // Specifies whether to also expose joystick hats as buttons, for compatibility with earlier versions of GLFW that did not have glfwGetJoystickHats.
	CocoaChdirResources Hint = 331777 // Specifies whether to set the current directory to the application to the Contents/Resources subdirectory of the application's bundle, if present.
	CocoaMenubar        Hint = 331778 // Specifies whether to create a basic menu bar, either from a nib or manually, when the first window is created, which is when AppKit is initialized.
)

// Window related hints/attributes.
const (
	Focused                Hint = 131073 // Specifies whether the window will be given input focus when created. This hint is ignored for full screen and initially hidden windows.
	Iconified              Hint = 131074 // Specifies whether the window will be minimized.
	Maximized              Hint = 131080 // Specifies whether the window is maximized.
	Visible                Hint = 131076 // Specifies whether the window will be initially visible.
	Hovered                Hint = 131083 // Specifies whether the cursor is currently directly over the content area of the window, with no other windows between. See Cursor enter/leave events for details.
	Resizable              Hint = 131075 // Specifies whether the window will be resizable by the user.
	Decorated              Hint = 131077 // Specifies whether the window will have window decorations such as a border, a close widget, etc.
	Floating               Hint = 131079 // Specifies whether the window will be always-on-top.
	AutoIconify            Hint = 131078 // Specifies whether fullscreen windows automatically iconify (and restore the previous video mode) on focus loss.
	CenterCursor           Hint = 131081 // Specifies whether the cursor should be centered over newly created full screen windows. This hint is ignored for windowed mode windows.
	TransparentFramebuffer Hint = 131082 // Specifies whether the framebuffer should be transparent.
	FocusOnShow            Hint = 131084 // Specifies whether the window will be given input focus when glfwShowWindow is called.
	ScaleToMonitor         Hint = 139276 // Specified whether the window content area should be resized based on the monitor content scale of any monitor it is placed on. This includes the initial placement when the window is created.
)

// Context related hints.
const (
	ClientAPI               Hint = 139265 // Specifies which client API to create the context for. Hard constraint.
	ContextVersionMajor     Hint = 139266 // Specifies the client API version that the created context must be compatible with.
	ContextVersionMinor     Hint = 139267 // Specifies the client API version that the created context must be compatible with.
	ContextRobustness       Hint = 139269 // Specifies the robustness strategy to be used by the context.
	ContextReleaseBehavior  Hint = 139273 // Specifies the release behavior to be used by the context.
	OpenGLForwardCompatible Hint = 139270 // Specifies whether the OpenGL context should be forward-compatible. Hard constraint.
	OpenGLDebugContext      Hint = 139271 // Specifies whether to create a debug OpenGL context, which may have additional error and performance issue reporting functionality. If OpenGL ES is requested, this hint is ignored.
	OpenGLProfile           Hint = 139272 // Specifies which OpenGL profile to create the context for. Hard constraint.
	ContextCreationAPI      Hint = 139275 // Specifies which context creation API to use to create the context.
)

// Framebuffer related hints.
const (
	ContextRevision        Hint = 139268
	RedBits                Hint = 135169 // Specifies the desired bit depth of the default framebuffer.
	GreenBits              Hint = 135170 // Specifies the desired bit depth of the default framebuffer.
	BlueBits               Hint = 135171 // Specifies the desired bit depth of the default framebuffer.
	AlphaBits              Hint = 135172 // Specifies the desired bit depth of the default framebuffer.
	DepthBits              Hint = 135173 // Specifies the desired bit depth of the default framebuffer.
	StencilBits            Hint = 135174 // Specifies the desired bit depth of the default framebuffer.
	AccumRedBits           Hint = 135175 // Specifies the desired bit depth of the accumulation buffer.
	AccumGreenBits         Hint = 135176 // Specifies the desired bit depth of the accumulation buffer.
	AccumBlueBits          Hint = 135177 // Specifies the desired bit depth of the accumulation buffer.
	AccumAlphaBits         Hint = 135178 // Specifies the desired bit depth of the accumulation buffer.
	AuxBuffers             Hint = 135179 // Specifies the desired number of auxiliary buffers.
	Stereo                 Hint = 135180 // Specifies whether to use stereoscopic rendering. Hard constraint.
	Samples                Hint = 135181 // Specifies the desired number of samples to use for multisampling. Zero disables multisampling.
	SRGBCapable            Hint = 135182 // Specifies whether the framebuffer should be sRGB capable.
	RefreshRate            Hint = 135183 // Specifies the desired refresh rate for full screen windows. If set to zero, the highest available refresh rate will be used. This hint is ignored for windowed mode windows.
	DoubleBuffer           Hint = 135184 // Specifies whether the framebuffer should be double buffered. You nearly always want to use double buffering. This is a hard constraint.
	CocoaGraphicsSwitching Hint = 143363 // Specifies whether to in Automatic Graphics Switching, i.e. to allow the system to choose the integrated GPU for the OpenGL context and move it between GPUs if necessary or whether to force it to always run on the discrete GPU.
	CocoaRetinaFramebuffer Hint = 143361 // Specifies whether to use full resolution framebuffers on Retina displays.
)

// Naming related hints. (Use with glfw.WindowHintString)
const (
	CocoaFrameNAME  Hint = 143362 // Specifies the UTF-8 encoded name to use for autosaving the window frame, or if empty disables frame autosaving for the window.
	X11ClassName    Hint = 147457 // Specifies the desired ASCII encoded class parts of the ICCCM WM_CLASS window property.nd instance parts of the ICCCM WM_CLASS window property.
	X11InstanceName Hint = 147458 // Specifies the desired ASCII encoded instance parts of the ICCCM WM_CLASS window property.nd instance parts of the ICCCM WM_CLASS window property.
)

// Values for the OpenGLProfile hint.
const (
	OpenGLAnyProfile    int = 0
	OpenGLCoreProfile   int = 204801
	OpenGLCompatProfile int = 204802
)

// Input modes.
const (
	CursorMode             InputMode = 208897 // See Cursor mode values
	StickyKeysMode         InputMode = 208898 // Value can be either 1 or 0
	StickyMouseButtonsMode InputMode = 208899 // Value can be either 1 or 0
	LockKeyMods            InputMode = 208900 // Value can be either 1 or 0
	RawMouseMotion         InputMode = 208901 // Value can be either 1 or 0
)

// Cursor mode values.
const (
	CursorNormal   int = 212993
	CursorHidden   int = 212994
	CursorDisabled int = 212995
)
