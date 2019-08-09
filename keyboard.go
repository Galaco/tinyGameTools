package tinygametools

import "github.com/go-gl/glfw/v3.2/glfw"

// Key maps to an integer value of a keyboard key (glfw values)
type Key int

const (
	KeyUnknown      = Key(glfw.KeyUnknown)
	KeyEscape       = Key(glfw.KeyEscape)
	KeyF1           = Key(glfw.KeyF1)
	KeyF2           = Key(glfw.KeyF2)
	KeyF3           = Key(glfw.KeyF3)
	KeyF4           = Key(glfw.KeyF4)
	KeyF5           = Key(glfw.KeyF5)
	KeyF6           = Key(glfw.KeyF6)
	KeyF7           = Key(glfw.KeyF7)
	KeyF8           = Key(glfw.KeyF8)
	KeyF9           = Key(glfw.KeyF9)
	KeyF10          = Key(glfw.KeyF10)
	KeyF11          = Key(glfw.KeyF11)
	KeyF12			= Key(glfw.KeyF12)
	KeyF13          = Key(glfw.KeyF13)
	KeyF14          = Key(glfw.KeyF14)
	KeyF15          = Key(glfw.KeyF15)
	KeyF16          = Key(glfw.KeyF16)
	KeyF17          = Key(glfw.KeyF17)
	KeyF18          = Key(glfw.KeyF18)
	KeyF19          = Key(glfw.KeyF19)
	KeyF20          = Key(glfw.KeyF20)
	KeyF21          = Key(glfw.KeyF21)
	KeyF22          = Key(glfw.KeyF22)
	KeyF23          = Key(glfw.KeyF23)
	KeyF24          = Key(glfw.KeyF24)
	KeyF25          = Key(glfw.KeyF25)
	KeyKP0          = Key(glfw.KeyKP0)
	KeyKP1          = Key(glfw.KeyKP1)
	KeyKP2          = Key(glfw.KeyKP2)
	KeyKP3          = Key(glfw.KeyKP3)
	KeyKP4          = Key(glfw.KeyKP4)
	KeyKP5          = Key(glfw.KeyKP5)
	KeyKP6          = Key(glfw.KeyKP6)
	KeyKP7          = Key(glfw.KeyKP7)
	KeyKP8          = Key(glfw.KeyKP8)
	KeyKP9          = Key(glfw.KeyKP9)
	KeyKPDecimal    = Key(glfw.KeyKPDecimal)
	KeyKPDivide     = Key(glfw.KeyKPDivide)
	KeyKPMultiply   = Key(glfw.KeyKPMultiply)
	KeyKPSubtract   = Key(glfw.KeyKPSubtract)
	KeyKPAdd        = Key(glfw.KeyKPAdd)
	KeyKPEnter      = Key(glfw.KeyKPEnter)
	KeyKPEqual      = Key(glfw.KeyKPEqual)
	Key1            = Key(glfw.Key1)
	Key2            = Key(glfw.Key2)
	Key3            = Key(glfw.Key3)
	Key4            = Key(glfw.Key4)
	Key5            = Key(glfw.Key5)
	Key6            = Key(glfw.Key6)
	Key7            = Key(glfw.Key7)
	Key8            = Key(glfw.Key8)
	Key9            = Key(glfw.Key9)
	Key0            = Key(glfw.Key0)
	KeySemicolon    = Key(glfw.KeySemicolon)
	KeyMinus        = Key(glfw.KeyMinus)
	KeyEqual        = Key(glfw.KeyEqual)
	KeyBackspace    = Key(glfw.KeyBackspace)
	KeyTab          = Key(glfw.KeyTab)
	KeyQ            = Key(glfw.KeyQ)
	KeyW            = Key(glfw.KeyW)
	KeyE            = Key(glfw.KeyE)
	KeyR            = Key(glfw.KeyR)
	KeyT            = Key(glfw.KeyT)
	KeyY            = Key(glfw.KeyY)
	KeyU            = Key(glfw.KeyU)
	KeyI            = Key(glfw.KeyI)
	KeyO            = Key(glfw.KeyO)
	KeyP            = Key(glfw.KeyP)
	KeyLeftBracket  = Key(glfw.KeyLeftBracket)
	KeyRightBracket = Key(glfw.KeyRightBracket)
	KeyGraveAccent  = Key(glfw.KeyGraveAccent)
	KeyWorld1       = Key(glfw.KeyWorld1)
	KeyWorld2       = Key(glfw.KeyWorld2)
	KeyBackslash    = Key(glfw.KeyBackslash)
	KeyCapsLock     = Key(glfw.KeyCapsLock)
	KeyA            = Key(glfw.KeyA)
	KeyS            = Key(glfw.KeyS)
	KeyD            = Key(glfw.KeyD)
	KeyF            = Key(glfw.KeyF)
	KeyG            = Key(glfw.KeyG)
	KeyH            = Key(glfw.KeyH)
	KeyJ            = Key(glfw.KeyJ)
	KeyK            = Key(glfw.KeyK)
	KeyL            = Key(glfw.KeyL)
	KeySemiColon    = Key(glfw.KeySemicolon)
	KeyApostrophe   = Key(glfw.KeyApostrophe)
	KeyEnter        = Key(glfw.KeyEnter)
	KeyLeftShift    = Key(glfw.KeyLeftShift)
	KeyZ            = Key(glfw.KeyZ)
	KeyX            = Key(glfw.KeyX)
	KeyC            = Key(glfw.KeyC)
	KeyV            = Key(glfw.KeyV)
	KeyB            = Key(glfw.KeyB)
	KeyN            = Key(glfw.KeyN)
	KeyM            = Key(glfw.KeyM)
	KeyComma        = Key(glfw.KeyComma)
	KeyPeriod       = Key(glfw.KeyPeriod)
	KeySlash        = Key(glfw.KeySlash)
	KeyRightShift   = Key(glfw.KeyRightShift)
	KeyLeftCtrl     = Key(glfw.KeyLeftControl)
	KeyLeftAlt      = Key(glfw.KeyLeftAlt)
	KeySpace        = Key(glfw.KeySpace)
	KeyRightAlt     = Key(glfw.KeyRightAlt)
	KeyRightCtrl    = Key(glfw.KeyRightControl)
	KeyDelete       = Key(glfw.KeyDelete)
	KeyInsert       = Key(glfw.KeyInsert)
	KeyUp           = Key(glfw.KeyUp)
	KeyLeft         = Key(glfw.KeyLeft)
	KeyDown         = Key(glfw.KeyDown)
	KeyRight        = Key(glfw.KeyRight)
	KeyPageUp       = Key(glfw.KeyPageUp)
	KeyPageDown     = Key(glfw.KeyPageDown)
	KeyHome         = Key(glfw.KeyHome)
	KeyEnd          = Key(glfw.KeyEnd)
	KeyScrollLock   = Key(glfw.KeyScrollLock)
	KeyNumLock      = Key(glfw.KeyNumLock)
	KeyPrintScreen  = Key(glfw.KeyPrintScreen)
	KeyPause        = Key(glfw.KeyPause)
	KeyLeftSuper    = Key(glfw.KeyLeftSuper)
	KeyRightSuper   = Key(glfw.KeyRightSuper)
	KeyMenu         = Key(glfw.KeyMenu)
	KeyLast         = Key(glfw.KeyLast)
)

// Keyboard provides the current state of the keyboard
type Keyboard struct {
	keyMap map[Key]bool

	externalKeyCallbacks  []glfw.KeyCallback
	externalCharCallbacks []glfw.CharCallback
}

// AddKeyCallback add a custom callback to when a key is pressed
func (keyboard *Keyboard) AddKeyCallback(callback glfw.KeyCallback) {
	keyboard.externalKeyCallbacks = append(keyboard.externalKeyCallbacks, callback)
}

// AddCharInputCallback add a custom callback for the true character press.
// See glfw docs to distinguish between this and KeyCallback
func (keyboard *Keyboard) AddCharInputCallback(callback glfw.CharCallback) {
	keyboard.externalCharCallbacks = append(keyboard.externalCharCallbacks, callback)
}

// RegisterCallbacks will set the passed windows event callbacks to this instance
func (keyboard *Keyboard) RegisterCallbacks(win *Window) {
	win.Handle().SetKeyCallback(keyboard.glfwKeyCallback)
	win.Handle().SetCharCallback(keyboard.glfwCharCallback)
}

// glfwKeyCallback is the glfw library callback function for handling
// keyboard input. Additional callbacks will be executed afterwards.
func (keyboard *Keyboard) glfwKeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		keyboard.keyMap[Key(key)] = true
	case glfw.Release:
		keyboard.keyMap[Key(key)] = false
	}

	for _, callback := range keyboard.externalKeyCallbacks {
		callback(w, key, scancode, action, mods)
	}
}

func (keyboard *Keyboard) glfwCharCallback(w *glfw.Window, char rune) {
	for _, callback := range keyboard.externalCharCallbacks {
		callback(w, char)
	}
}

// IsKeyDown returns whether a particular key is pressed.
func (keyboard *Keyboard) IsKeyDown(key Key) bool {
	return keyboard.keyMap[key]
}

// NewKeyboard returns a new keyboard.
func NewKeyboard() *Keyboard {
	return &Keyboard{
		keyMap: map[Key]bool{},
	}
}
