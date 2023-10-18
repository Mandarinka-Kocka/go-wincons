# Go Windows Console Package

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Overview

WinCons is a binding library for the Win32 console API, primarily aimed at addressing the relative weaknesses of Go in interacting with the system console.

## Advantages

- **Low Learning Curve**
The naming of functions and types in the WinCons library closely aligns with the Win32 API. This means that if you have experience with the Win32 API, you can quickly start using this library without incurring additional learning costs.

- **Simplified Usage**
"WinCons" not only offers standard console operation functionality, but also provides a range of advanced features designed to streamline common console interactions. These extended capabilities encompass cursor control, buffer management, console input handling, and more.

## Installation

`go get https://github.com/Mandarinkocka/go-wincons`	

## sample

```go
package main

import (
	"log"

	"github.com/mandarinkocka/go-wincons"
)

func main() {
	hConsOutput, err := cons.GetStdHandle(cons.StdOutputHandle)
	if err != nil {
		log.Fatalln(err)
	}

	hConsInput, err := cons.GetStdHandle(cons.StdInputHandle)
	if err != nil {
		log.Fatalln(err)
	}

	if !cons.IsEnableMode(hConsOutput, cons.EnableVirtualTerminalProcessing) {
		if err := cons.EnableMode(hConsOutput, cons.EnableVirtualTerminalProcessing); err != nil {
			log.Fatalln(err)
		}
	}

	cons.ClearScreenBuffer(hConsOutput)

	cons.SetInputCodePage(cons.Utf8)
	cons.SetOutputCodePage(cons.Utf8)
	cons.SetWindowTitle("sample")
	cons.SetCursorVisible(hConsOutput, false)
	defer cons.SetCursorVisible(hConsOutput, true)

	cons.Pause(hConsInput, "\033[4mPress any key to exit\033[0m")
}
```

## License
[MIT License](./LICENSE)