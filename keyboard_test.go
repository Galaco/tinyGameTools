package tinygametools

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"testing"
)

func TestKeyboard_AddCharInputCallback(t *testing.T) {
	sut := NewKeyboard()
	cb := func(w *glfw.Window, char rune) {}

	if len(sut.externalCharCallbacks) > 0 {
		t.Error("keyboard has unexpected charInput callback(s)")
	}
	sut.AddCharInputCallback(cb)

	if len(sut.externalCharCallbacks) != 1 {
		t.Error("keyboard has unexpected charInput callback(s)")
	}
}

func TestKeyboard_AddKeyCallback(t *testing.T) {
	sut := NewKeyboard()
	cb := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {}

	if len(sut.externalKeyCallbacks) > 0 {
		t.Error("keyboard has unexpected Key callback(s)")
	}
	sut.AddKeyCallback(cb)

	if len(sut.externalKeyCallbacks) != 1 {
		t.Error("keyboard has unexpected Key callback(s)")
	}
}

func TestKeyboard_IsKeyDown(t *testing.T) {
	sut := NewKeyboard()
	sut.keyMap[Key4] = true

	if sut.IsKeyDown(Key4) != true {
		t.Error("keyboard returned incorrect state for key")
	}

	sut.keyMap[Key4] = false
	if sut.IsKeyDown(Key4) != false {
		t.Error("keyboard returned incorrect state for key")
	}
}

func TestKeyboard_RegisterCallbacks(t *testing.T) {
	t.Skip("requires an active window")
}

func TestNewKeyboard(t *testing.T) {
	sut := NewKeyboard()

	if sut == nil {
		t.Errorf("nil returned, but expected object")
	}
}
