package tinygametools

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"testing"
)

func TestNewWindow(t *testing.T) {
	t.Skip("requires an supported gl environment")
}

func TestWindow_Handle(t *testing.T) {
	handle := &glfw.Window{}
	sut := Window{
		glfwHandle: handle,
	}

	if sut.Handle() != handle {
		t.Error("window returning unexpect glfw.Window handle")
	}
}
