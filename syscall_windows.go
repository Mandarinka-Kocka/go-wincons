package cons

import (
	"syscall"
	"unsafe"
)

// Retrieves a handle for standard input, output, or error streams in Windows.
//
// Parameters:
//
//	flag: A flag indicating whether to get the standard input, output, or error handle.
//
// Returns:
//
//	Handle: A handle to the requested standard stream.
//	error: If the function successfully retrieves the handle, it returns nil. Otherwise, it returns an error.
func GetStdHandle(flag uint32) (Handle, error) {
	hStdHandle, _, err := procGetStdHandle.Call(uintptr(flag))
	if err != errorSuccess {
		return invalidHandle, err
	}

	return Handle(hStdHandle), nil
}

// Sets the console mode for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the console mode should be set.
//
// Returns:
//
//	uint32: The new console mode after setting it.
//	error: If the function successfully sets the console mode, it returns nil. Otherwise, it returns an error.
func GetMode(hStdout Handle) (DWord, error) {
	var dwMode DWord
	if _, _, err := procGetConsoleMode.Call(uintptr(hStdout), touintptr(&dwMode)); err != errorSuccess {
		return 0, err
	}

	return dwMode, nil
}

// retrieves the cursor information for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which cursor information is to be obtained.
//
// Returns:
//
//	CursorInfo: The cursor information, including size and visibility.
//	error: If the function successfully retrieves the cursor information, it returns nil. Otherwise, it returns an error.
func GetCursorInfo(hStdout Handle) (CursorInfo, error) {
	var curinfo CursorInfo
	if _, _, err := procGetConsoleCursorInfo.Call(uintptr(hStdout), touintptr(&curinfo)); err != errorSuccess {
		return CursorInfo{}, err
	}

	return curinfo, nil
}

// retrieves the current console code page (input and output code page) in Windows.
//
// Returns:
//
//	uint32: The current code page.
//	error: If the function successfully retrieves the code page, it returns nil. Otherwise, it returns an error.
func GetInputCodePage() (uint32, error) {
	cp, _, err := procGetConsoleCP.Call()
	if err != errorSuccess {
		return 0, err
	}

	return uint32(cp), nil
}

// retrieves information about the screen buffer of the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which screen buffer information is to be obtained.
//	p: A pointer to a ScreenBufferInfo structure where the information will be stored.
//
// Returns:
//
//	error: If the function successfully retrieves the screen buffer information, it returns nil. Otherwise, it returns an error.
func GetScreenBufferInfo(hStdout Handle, p *ScreenBufferInfo) error {
	if _, _, err := procGetConsoleScrrenBufferInfo.Call(uintptr(hStdout), touintptr(p)); err != errorSuccess {
		return err
	}

	return nil
}

// Retrieves the current output code page for the console in Windows.
//
// Returns:
//
//	uint32: The current output code page.
//	error: If the function successfully retrieves the output code page, it returns nil. Otherwise, it returns an error.
func GetOutputCodePage() (uint32, error) {
	cp, _, err := procGetConsoleOutputCP.Call()
	if err != errorSuccess {
		return 0, err
	}

	return uint32(cp), nil
}

// reads console input records from the specified standard input handle in Windows.
//
// Parameters:
//
//	hStdin: The handle to the standard input stream from which input records will be read.
//	buffer: A pointer to a buffer where input records will be stored.
//	length: The maximum number of input records to read.
//	counter: A pointer to a counter that will receive the actual number of input records read.
//
// Returns:
//
//	error: If the function successfully reads the input records, it returns nil. Otherwise, it returns an error.
func ReadInput(hStdin Handle, buffer unsafe.Pointer, length uint32, counter *uint32) error {
	_, _, err := procReadConsoleInput.Call(uintptr(hStdin), uintptr(buffer), uintptr(length), touintptr(counter))
	if err != errorSuccess {
		return err
	}

	return nil
}

// Sets the console mode for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the console mode should be set.
//	dwMode: The new console mode to be applied to the standard output.
//
// Returns:
//
//	error: If the function successfully sets the console mode, it returns nil. Otherwise, it returns an error.
func SetMode(hStdout Handle, mode DWord) error {
	if _, _, err := procSetConsoleMode.Call(uintptr(hStdout), uintptr(mode)); err != errorSuccess {
		return err
	}

	return nil
}

// Sets the output code page for the console in Windows.
//
// Parameters:
//
//	cp: The code page to set as the output code page for the console.
//
// Returns:
//
//	error: If the function successfully sets the output code page, it returns nil. Otherwise, it returns an error.
func SetOutputCodePage(cp uint32) error {
	if _, _, err := procSetConsoleOutputCP.Call(uintptr(cp)); err != errorSuccess {
		return err
	}

	return nil
}

// Sets the cursor position for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the cursor position is to be set.
//	pos: The new cursor position specified as a Coord structure.
//
// Returns:
//
//	error: If the function successfully sets the cursor position, it returns nil. Otherwise, it returns an error.
func SetCursorPosition(hStdout Handle, newpos Coord) error {
	if _, _, err := procSetCursorPosition.Call(uintptr(hStdout), strutouintptr(&newpos)); err != errorSuccess {
		return err
	}

	return nil
}

// Sets the input code page for the console in Windows.
//
// Parameters:
//
//	cp: The code page to set as the input code page for the console.
//
// Returns:
//
//	error: If the function successfully sets the input code page, it returns nil. Otherwise, it returns an error.
func SetInputCodePage(cp uint32) error {
	if _, _, err := procSetConsoleCP.Call(uintptr(cp)); err != errorSuccess {
		return err
	}

	return nil
}

// Sets the title of the console window in Windows.
//
// Parameters:
//
//	title: The string to set as the new title of the console window.
//
// Returns:
//
//	error: If the function successfully sets the title, it returns nil. Otherwise, it returns an error.
func SetWindowTitle(title string) error {
	cstr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}

	if _, _, err := procSetConsoleTitle.Call(touintptr(cstr)); err != nil {
		return err
	}

	return nil
}

// Sets the cursor information for the specified standard output handle in Windows.
//
// Parameters:
//   hStdout: The handle to the standard output stream for which cursor information is to be set.
//   p: A pointer to a CursorInfo structure containing cursor size and visibility settings.
//
// Returns:
//   error: If the function successfully sets the cursor information, it returns nil. Otherwise, it returns an error.

// Fills the specified number of character cells with a given character in the console window.
func SetCursorInfo(hStdout Handle, p *CursorInfo) error {
	if _, _, err := procSetConsoleCursorInfo.Call(uintptr(hStdout), touintptr(p)); err != errorSuccess {
		return err
	}

	return nil
}

// Fills the specified number of character cells with a given character in the console window.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream where the characters will be filled.
//	char: The character to be used for filling.
//	length: The number of character cells to be filled.
//	wcoord: The starting coordinates for filling.
//
// Returns:
//
//	uint16: The number of character cells successfully filled.
//	error: If the function successfully fills the character cells, it returns nil. Otherwise, it returns an error.
func FillCharacter(hStdout Handle, char rune, length int, wcoord Coord) (uint16, error) {
	var counter uint16

	if _, _, err := procFillConsoleOutputCharacter.Call(
		uintptr(hStdout), uintptr(char),
		uintptr(length), strutouintptr(&wcoord),
		touintptr(&counter)); err != errorSuccess {
		return counter, err
	}

	return counter, nil
}

// Fills the specified number of character cells with the given attribute in the console window.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream where the attributes will be filled.
//	attribute: The attribute to be used for filling.
//	length: The number of character cells to be filled.
//	wcoord: The starting coordinates for filling.
//
// Returns:
//
//	uint16: The number of character cells with attributes successfully filled.
//	error: If the function successfully fills the character cells with attributes, it returns nil. Otherwise, it returns an error.
func FillAttribute(hStdout Handle, attribute uint16, length int, wcoord Coord) (uint16, error) {
	var counter uint16

	if _, _, err := procFillConsoleOutputAttribute.Call(
		uintptr(hStdout), uintptr(attribute),
		uintptr(length), strutouintptr(&wcoord),
		touintptr(&counter)); err != errorSuccess {
		return 0, err
	}

	return counter, nil
}

// Writes the specified attribute to the character cells in the console window.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream where the attributes will be written.
//	attribute: The attribute to be written to the character cells.
//	length: The number of character cells to write the attribute to.
//	wcoord: The starting coordinates for writing the attribute.
//
// Returns:
//
//	uint16: The number of character cells with attributes successfully written.
//	error: If the function successfully writes the attribute to the character cells, it returns nil. Otherwise, it returns an error.
func WriteAttribute(hStdout Handle, attribute uint16, length uint32, wcoord Coord) (uint16, error) {
	var counter uint16

	if _, _, err := procWriteConsoleOutputAttribute.Call(
		uintptr(hStdout), touintptr(&attribute),
		uintptr(length), strutouintptr(&wcoord),
		touintptr(&counter)); err != errorSuccess {
		return counter, err
	}

	return counter, nil
}

// Writes the specified character to a specified number of character cells in the console window.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream where characters will be written.
//	char: The character to be written to the character cells.
//	length: The number of character cells to write the character to.
//	wcoord: The starting coordinates for writing the character.
//
// Returns:
//
//	uint16: The number of character cells with the character successfully written.
//	error: If the function successfully writes the character to the character cells, it returns nil. Otherwise, it returns an error.
func WriteCharacter(hStdout Handle, char uint16, length uint32, wcoord Coord) (uint16, error) {
	var counter uint16

	if _, _, err := procWriteConsoleOutputCharacter.Call(
		uintptr(hStdout), uintptr(char),
		uintptr(length), strutouintptr(&wcoord),
		touintptr(&counter)); err != errorSuccess {
		return counter, err
	}

	return counter, nil
}

// Scrolls a portion of the screen buffer contents within the console window.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream where scrolling will be performed.
//	scrollrect: A pointer to a SmallRect structure specifying the portion of the screen buffer to be scrolled.
//	cliprect: A pointer to a SmallRect structure specifying the clipping area for scrolling.
//	dest: A Coord structure specifying the destination coordinates for the scroll operation.
//	fill: A CharInfo structure specifying the character attributes to be used for filling new space created by scrolling.
//
// Returns:
//
//	error: If the function successfully performs the scrolling operation, it returns nil. Otherwise, it returns an error.
func ScrollScreenBuffer(hStdout Handle, scrollrect, cliprect *SmallRect, dest Coord, fill CharInfo) error {
	if _, _, err := procScrollConsoleScreenBuffer.Call(
		uintptr(hStdout), touintptr(scrollrect),
		touintptr(cliprect), touintptr(&dest),
		touintptr(&fill)); err != errorSuccess {
		return err
	}

	return nil
}
