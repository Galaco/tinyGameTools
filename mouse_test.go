package tinygametools

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"testing"
)

func TestMouse_AddMouseButtonCallback(t *testing.T) {
	sut := NewMouse()
	cb := func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {}

	if len(sut.externalMouseButtonCallbacks) > 0 {
		t.Error("mouse has unexpected Key callback(s)")
	}
	sut.AddMouseButtonCallback(cb)

	if len(sut.externalMouseButtonCallbacks) != 1 {
		t.Error("mouse has unexpected Key callback(s)")
	}
}

func TestMouse_AddMousePosCallback(t *testing.T) {
	sut := NewMouse()
	cb := func(w *glfw.Window, xpos float64, ypos float64) {}

	if len(sut.externalCursorPosCallbacks) > 0 {
		t.Error("mouse has unexpected CursorPos callback(s)")
	}
	sut.AddMousePosCallback(cb)

	if len(sut.externalCursorPosCallbacks) != 1 {
		t.Error("mouse has unexpected CursorPos callback(s)")
	}
}

func TestMouse_AddMouseScrollCallback(t *testing.T) {
	sut := NewMouse()
	cb := func(w *glfw.Window, xoff float64, yoff float64) {}

	if len(sut.externalMouseScrollCallbacks) > 0 {
		t.Error("mouse has unexpected Scroll callback(s)")
	}
	sut.AddMouseScrollCallback(cb)

	if len(sut.externalMouseScrollCallbacks) != 1 {
		t.Error("mouse has unexpected Scroll callback(s)")
	}
}

func TestMouse_IsButtonPressed(t *testing.T) {
	sut := NewMouse()
	sut.keyMap[MouseButton3] = true

	if sut.IsButtonPressed(MouseButton3) != true {
		t.Error("mouse returned incorrect state for button")
	}

	sut.keyMap[MouseButton3] = false
	if sut.IsButtonPressed(MouseButton3) != false {
		t.Error("mouse returned incorrect state for button")
	}
}

func TestMouse_RegisterCallbacks(t *testing.T) {
	t.Skip("requires an active window")
}

func TestMouse_X(t *testing.T) {
	sut := NewMouse()
	sut.x = 32
	if sut.X() != 32 {
		t.Errorf("mouse x pos returned incorrectly. Expected: 32, but received: %f", sut.X())
	}
}

func TestMouse_Y(t *testing.T) {
	sut := NewMouse()
	sut.y = 32
	if sut.Y() != 32 {
		t.Errorf("mouse y pos returned incorrectly. Expected: 32, but received: %f", sut.Y())
	}
}

func TestNewMouse(t *testing.T) {
	sut := NewMouse()

	if sut == nil {
		t.Errorf("nil returned, but expected object")
	}
}
