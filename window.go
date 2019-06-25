package tinygametools

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

// Window
type Window struct {
	glfwHandle *glfw.Window
}

// Handle returns internal glfwHandle handle
func (w *Window) Handle() *glfw.Window {
	return w.glfwHandle
}

// NewWindow constructs a new GLFW Window
func NewWindow(width int, height int, name string) (*Window, error) {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return nil, err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		return nil, err
	}

	return &Window{
		glfwHandle: window,
	}, nil
}
