[![GoDoc](https://godoc.org/github.com/Galaco/tinyGameTools?status.svg)](https://godoc.org/github.com/Galaco/tinyGameTools)
[![Go report card](https://goreportcard.com/badge/github.com/galaco/tinyGameTools)](https://goreportcard.com/badge/github.com/galaco/tinyGameTools)
[![GolangCI](https://golangci.com/badges/github.com/galaco/tinyGameTools.svg)](https://golangci.com)
[![Build Status](https://travis-ci.com/Galaco/tinyGameTools.svg?branch=master)](https://travis-ci.com/Galaco/tinyGameTools)
[![CircleCI](https://circleci.com/gh/Galaco/tinyGameTools/tree/master.svg?style=svg)](https://circleci.com/gh/Galaco/tinyGameTools/tree/master)
[![codecov](https://codecov.io/gh/Galaco/tinyGameTools/branch/master/graph/badge.svg)](https://codecov.io/gh/Galaco/tinyGameTools)

# TinyGameTools

### What is it?
TinyGameTools is a lightweight package for getting started with Go game development.

It does not provide any sort of engine, only a set of lightweight utilities for creating and 
interacting with an OpenGL 4.1 window.

### Features
* Window. A simple to use glfw window creation wrapper. Will create an OpenGL 4.1 ready window
* Keyboard. Hooks into Window, and allows for querying button state, and custom input callbacks
* Mouse. Hooks into Window, and allows for query Mouse state, and custom input callbacks
* Event bus. A simple event bus for subscribing and publishing messages. Custom messages should fulfil the `Event` interface,
and away you go.

### Examples
```go
package main

import (
	"github.com/galaco/tinyGameTools"
)

func main() {
	// Create a window
	win,_ := tinygametools.NewWindow(640, 480, "My Window!")
	
	// Create a keyboard, then register callbacks
	kb := tinygametools.NewKeyboard()
	kb.RegisterCallbacks(win)
	
	// Create the mouse handler, then register callbacks
	mouse := tinyGameTools.NewMouse()
	mouse.RegisterCallbacks(win)
}
```

@TODO more documentation as features get added