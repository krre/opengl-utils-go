package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.1/glfw"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	videoMode := glfw.GetPrimaryMonitor().GetVideoMode()
	screenWidth := videoMode.Width
	screenHeight := videoMode.Height
	screenRatio := 0.8

	width := int(screenRatio * float64(screenWidth))
	height := int(screenRatio * float64(screenHeight))
	x := int((screenWidth - width) / 2)
	y := int((screenHeight - height) / 2)

	window, err := glfw.CreateWindow(width, height, "Paint", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetPos(x, y)
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	window.SwapBuffers()

	for !window.ShouldClose() {

		glfw.PollEvents()
	}
}
