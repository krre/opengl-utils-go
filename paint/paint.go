package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.1/glfw"
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

	window, err := glfw.CreateWindow(width, height, "Window", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetPos(x, y)
	window.MakeContextCurrent()

	for !window.ShouldClose() {

		glfw.PollEvents()
	}
}
