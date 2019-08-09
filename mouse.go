package tinygametools

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	MouseButton1      = Key(glfw.MouseButton1)
	MouseButton2      = Key(glfw.MouseButton2)
	MouseButton3      = Key(glfw.MouseButton3)
	MouseButton4      = Key(glfw.MouseButton4)
	MouseButton5      = Key(glfw.MouseButton5)
	MouseButton6      = Key(glfw.MouseButton6)
	MouseButton7      = Key(glfw.MouseButton7)
	MouseButton8      = Key(glfw.MouseButton8)
	MouseButtonLast   = Key(glfw.MouseButtonLast)
	MouseButtonLeft   = Key(glfw.MouseButtonLeft)
	MouseButtonRight  = Key(glfw.MouseButtonRight)
	MouseButtonMiddle = Key(glfw.MouseButtonMiddle)
)

// Mouse is a small wrapper for handling mouse information flow
// between an application and glfw
type Mouse struct {
	x, y                         float64
	keyMap                       map[Key]bool
	externalCursorPosCallbacks   []glfw.CursorPosCallback
	externalMouseButtonCallbacks []glfw.MouseButtonCallback
	externalMouseScrollCallbacks []glfw.ScrollCallback
}

// X returns mouse X coordinate
func (mouse *Mouse) X() float64 {
	return mouse.x
}

// Y returns mouse Y coordinate
func (mouse *Mouse) Y() float64 {
	return mouse.y
}

// IsButtonPressed returns if a particular mouse button has been pressed.
func (mouse *Mouse) IsButtonPressed(key Key) bool {
	return mouse.keyMap[key]
}

// AddMousePosCallback
func (mouse *Mouse) AddMousePosCallback(callback glfw.CursorPosCallback) {
	mouse.externalCursorPosCallbacks = append(mouse.externalCursorPosCallbacks, callback)
}

// AddMouseButtonCallback
func (mouse *Mouse) AddMouseButtonCallback(callback glfw.MouseButtonCallback) {
	mouse.externalMouseButtonCallbacks = append(mouse.externalMouseButtonCallbacks, callback)
}

// AddMouseScrollCallback
func (mouse *Mouse) AddMouseScrollCallback(callback glfw.ScrollCallback) {
	mouse.externalMouseScrollCallbacks = append(mouse.externalMouseScrollCallbacks, callback)
}

// RegisterCallbacks will set the passed windows event callbacks to this instance
func (mouse *Mouse) RegisterCallbacks(win *Window) {
	win.Handle().SetCursorPosCallback(mouse.glfwCursorPosCallback)
	win.Handle().SetMouseButtonCallback(mouse.glfwMouseButtonCallback)
	win.Handle().SetScrollCallback(mouse.glfwScrollCallback)
}

func (mouse *Mouse) glfwCursorPosCallback(window *glfw.Window, xpos float64, ypos float64) {
	mouse.x = xpos
	mouse.y = ypos

	for _, callback := range mouse.externalCursorPosCallbacks {
		callback(window, xpos, ypos)
	}
}

func (mouse *Mouse) glfwMouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		mouse.keyMap[Key(button)] = true
	case glfw.Release:
		mouse.keyMap[Key(button)] = false
	}

	for _, callback := range mouse.externalMouseButtonCallbacks {
		callback(w, button, action, mod)
	}
}

func (mouse *Mouse) glfwScrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	for _, callback := range mouse.externalMouseScrollCallbacks {
		callback(w, xoff, yoff)
	}
}

// NewMouse returns a new Mouse
func NewMouse() *Mouse {
	return &Mouse{
		keyMap: map[Key]bool{},
	}
}
